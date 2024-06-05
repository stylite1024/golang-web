# 参考：https://github.com/adnanh/webhook/blob/master/Makefile
# 二进制包名
# BINARY ?=$(shell grep "AppName.*=" config/info.go | awk  -F'"' '{print $2}')
BINARY	?=$(shell grep "AppName.*=" config/info.go | grep -oP '"\K[^"]+')

# 版本号
VERSION	?=$(shell grep "Version.*=" config/info.go | grep -oP '"\K[^"]+')

# go版本
# GO_VERSION ?= $(shell go version | grep -o '[0-9]\+.[0-9]\+.[0-9]\+')
GO_VERSION	?= $(shell grep '^go' go.mod | cut -d' ' -f2)

# 设置当前的GOOS和GOARCH
# 系统平台，如linux、windows
GOOS	?= $(shell go env GOOS)
# 系统架构，如amd64
GOARCH	?= $(shell go env GOARCH)

# 查看支持操作系统和架构命令：go tool dist list
# 定义可以打包的操作系统列表
PLATFORMS	:= windows linux darwin
# 定义可以打包的架构列表
ARCHES	:= amd64 arm64 386 

# 打包docker镜像变量
DOCKERFILE_PATH         ?= ./Dockerfile
DOCKERBUILD_CONTEXT     ?= ./
DOCKER_NAMESPACE		?= go-app

# 使用内置目标名.PHONY声明这些“伪目标”名是“伪目标”，而不是与“伪目标”同名的文件
.PHONY: help update run build build-image build-run-container remove-image remove-container push-image build-push-image all clean 

default: help

update:
	@go mod tidy

run: update
	@go run .

build: update
ifeq ("$(GOOS)", "windows")
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-s -w" -o ./bin/${BINARY}.exe  .
else
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-s -w" -o ./bin/${BINARY}  .
endif

build-image: remove-container remove-image
# @docker build -t  ${DOCKER_REPO}/${BINARY}:${VERSION} -f ${DOCKERFILE_PATH} --build-arg GOVERSION="${GO_VERSION}" ${DOCKERBUILD_CONTEXT}
	@docker build -t  ${BINARY}:${VERSION} -f ${DOCKERFILE_PATH} --build-arg GOVERSION="${GO_VERSION}" ${DOCKERBUILD_CONTEXT}
	@docker tag ${BINARY}:${VERSION} ${BINARY}:latest

build-run-container: remove-container build-image
	@docker run -d --name ${BINARY} ${BINARY}:${VERSION}
	@docker ps|grep ${BINARY}

remove-image: remove-container
	@docker rmi -f ${BINARY}:latest
	@docker rmi -f ${BINARY}:${VERSION}

remove-container:
	@docker rm -f ${BINARY}

push-image:
	@docker tag	${BINARY}:${VERSION} ${DOCKER_NAMESPACE}/${BINARY}:${VERSION}
	@docker tag ${BINARY}:latest ${DOCKER_NAMESPACE}/${BINARY}:latest
	@docker push ${DOCKER_NAMESPACE}/${BINARY}:${VERSION}
	@docker push ${DOCKER_NAMESPACE}/${BINARY}:latest

build-push-image: build-image push-image

all: update
	@for platform in $(PLATFORMS); do \
		for arch in $(ARCHES); do \
			echo "building  ${BINARY}-$${platform}-$${arch}  binary file..."; \
			GOOS=$${platform} GOARCH=$${arch} go build -ldflags "-s -w" -o "./bin/${BINARY}-$${platform}-$${arch}" .; \
		done; \
	done

clean:
	@if [ -d bin ] ; then rm -rf ./bin ; fi
	@echo "clean successful"

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "	help				: Show help"
	@echo "	update				: Run 'go mod tidy'"
	@echo "	Run					: Run 'go run .'"
	@echo "	Build				: Build the binary of this project for current platform"
	@echo "	build-image			: Build docker image"
	@echo "	build-run-container	: Build docker image and run docker container"
	@echo "	remove-images		: Remove docker image"
	@echo "	remove-container	: Remove docker container"
	@echo "	all 				: Build multiple platform multiple arch binary of this project"
	@echo "	clean				: Cleaning up all the generated binary files"
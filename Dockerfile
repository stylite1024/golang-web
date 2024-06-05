ARG GOVERSION="1.22.0"

# 构建镜像
FROM docker.io/library/golang:${GOVERSION}-alpine AS Builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 工作目录：/build
WORKDIR /build

# 将代码复制到容器中的/buil下
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -ldflags "-s -w" -o app .

# 运行镜像
FROM docker.io/library/alpine:3.20

# 信息
LABEL maintainer="https://oby.ink"

# 移动到用于存放生成的二进制文件的 /dist 目录
# WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
COPY --from=Builder /build/app /usr/bin/app

# 拷贝docker-entrypoint.sh
COPY scripts/docker-entrypoint.sh /usr/bin/docker-entrypoint.sh

# 声明服务端口
EXPOSE 8080

# 数据卷
VOLUME ["/data"]

# 设置容器启动时执行的命令,用于检查环境
ENTRYPOINT ["/usr/bin/docker-entrypoint.sh"]

# 启动容器时运行的命令
CMD ["app"]

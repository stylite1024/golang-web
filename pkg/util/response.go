package util

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// https://juejin.cn/post/7258119695010824250
// https://www.jb51.net/jiaoben/293311yib.html

type Response struct {
	Code int         `json:"code" example:"200"` // 响应状态码
	Data interface{} `json:"data"`               // 响应数据
	Msg  interface{} `json:"msg"`                // 响应信息
	// RequestId string      `json:"requestId"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c iris.Context) interface{} {
	var requestId interface{}
	requestId = c.GetID()
	if requestId == nil {
		requestId = uuid.New().String()
		c.SetID(requestId)
	}
	return requestId
}

// 失败数据处理
func ErrorWithErr(c iris.Context, code int, msg string, err error) {
	var res Response
	res.Code = code
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}
	// res.RequestId = GenerateMsgIDFromContext(c)
	c.JSON(res)
}

// 失败数据处理
func Error(c iris.Context, code int, msg interface{}) {
	var res Response
	res.Code = code
	res.Msg = msg
	// res.RequestId = GenerateMsgIDFromContext(c)
	c.JSON(res)
}

// 通常成功数据处理
func OK(c iris.Context, msg string, data interface{}) {
	var res Response
	res.Code = 200
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	// res.RequestId = GenerateMsgIDFromContext(c)
	c.JSON(res)
}

// 分页数据处理
func PageOK(c iris.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res Page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	// res.RequestId = GenerateMsgIDFromContext(c)
	OK(c, msg, res)
}

package response

import (
	"framework/component/paginate"
	"framework/component/trace"
	"framework/util/strings"
	"github.com/gin-gonic/gin"
	"reflect"
)

const(
	// 请求ID前缀
	requestPrefix = "request:"
	// 成功提示信息
	successMessage = "success"
)

// 数据对象
type Data map[string]interface{}

// Response
type Response struct {
	// 提示信息
	Message string `json:"msg"`

	// 业务码
	Code int `json:"code"`

	// 数据
	Data interface{} `json:"data"`

	// 元数据，如分页,请求ID
	Meta Meta `json:"meta"`
}


// requestId
func requestId() string  {
	return requestPrefix + strings.UUID()
}

// meta
func meta() Meta {
	return Meta{
		RequestId:  requestId(),
		TraceId:  trace.Get(),
	}
}
// Meta 元数据
type Meta struct {
	// 请求ID
	RequestId string `json:"request_id"`

	// 全局追踪ID,服务化必备
	TraceId string `json:"trace_id"`
	// 分页
	Pagination *paginate.Pagination `json:"pagination"`
}


// wrapper context装饰器
type wrapper struct {
	ctx *gin.Context
}

// Wrap
func Wrap(ctx *gin.Context) *wrapper  {
	return &wrapper{
		ctx: ctx,
	}
}

// newResponse 构造一个Response
func newResponse(code int, message string) Response  {
	return Response{
		Message: message,
		Code:    code,
		Data:    nil,
		Meta:    Meta{
			RequestId:  requestId(),
			TraceId:  trace.Get(),
			Pagination: nil,
		},
	}
}

// output output response
func (w *wrapper) output(response Response) {
	w.ctx.JSON(200, response)
}

// Success 输出成功响应
func (w *wrapper) Success(data interface{})  {
	response := newResponse(0, successMessage)
	response.Data = data

	w.output(response)
}

// Error 输出错误响应
func (w *wrapper) Error(code int, message string)  {
	response := newResponse(code, message)

	w.output(response)
}

// Paginate 输出分页响应内容
func (w *wrapper) Paginate(data interface{}, pagination *paginate.Pagination) {
	response := newResponse(0, successMessage)
	// 如果data为nil,则默认设置为[]
	v := reflect.ValueOf(data)
	if v.IsNil() {
		data = []interface{}{}
	}
	response.Data = data
	response.Meta.Pagination = pagination

	w.output(response)
}
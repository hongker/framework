package response

import (
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/component/paginate"
	"reflect"
)

// wrapper context装饰器
type wrapper struct {
	ctx *gin.Context
}

// Wrap
func Wrap(ctx *gin.Context) *wrapper {
	return &wrapper{
		ctx: ctx,
	}
}

// output output response
func (w *wrapper) output(response *Response) {
	w.ctx.JSON(200, response)
}

// Success 输出成功响应
func (w *wrapper) Success(data interface{}) {
	response := defaultPool.Get()
	response.Data = data

	w.output(response)
}

// Error 输出错误响应
func (w *wrapper) Error(code int, message string) {
	response := defaultPool.Get()
	response.Code = code
	response.Message = message

	w.output(response)
}

// Paginate 输出分页响应内容
func (w *wrapper) Paginate(data interface{}, pagination *paginate.Pagination) {
	response := defaultPool.Get()
	// 如果data为nil,则默认设置为[]
	v := reflect.ValueOf(data)
	if v.IsNil() {
		data = []interface{}{}
	}
	response.Data = data
	response.Meta.Pagination = pagination

	w.output(response)
}


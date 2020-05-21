package response

import (
	"github.com/hongker/framework/component/paginate"
	"github.com/hongker/framework/component/trace"
	"github.com/hongker/framework/util/strings"
)

const (
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

func (r *Response) Reset() {
	r.Code = 0
	r.Meta = meta()
	r.Message = successMessage
	r.Data = nil
}

// requestId
func requestId() string {
	return requestPrefix + strings.UUID()
}

// meta
func meta() Meta {
	return Meta{
		RequestId: requestId(),
		TraceId:   trace.Get(),
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


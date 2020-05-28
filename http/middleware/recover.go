package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/component/log"
	"github.com/hongker/framework/component/trace"
	"github.com/hongker/framework/errors"
	"github.com/hongker/framework/http/response"
)

// Recover 处理因系统panic导致的错误，需要拦截，定制化返回响应内容，不能直接输出与系统相关的内容，可以自定义
func Recover(ctx *gin.Context) {
	defer func(traceId string) {
		trace.Set(traceId)
		defer trace.GC()
		if r := recover(); r != nil {
			if err, ok := r.(*errors.Error); ok {
				response.Wrap(ctx).Error(err.Code, err.Message)
			} else {
				log.Error("Recover", log.Content{
					"param":      r,
					"trace_info": "", // 记录触发错误的调用链路
				})
				response.Wrap(ctx).Error(500, "System Error")
			}
		}
	}(trace.Get())
	ctx.Next()
}

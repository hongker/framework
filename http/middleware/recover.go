package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/component/log"
	"github.com/hongker/framework/http/response"
)

// Recover 处理因系统panic导致的错误，需要拦截，定制化返回响应内容，不能直接输出与系统相关的内容，可以自定义
func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			response.Wrap(ctx).Error(500, "System Error")

			log.Error("Recover", log.Content{
				"param": r,
				"trace_info": "", // 记录触发错误的调用链路
			})
		}
	}()
	ctx.Next()
}

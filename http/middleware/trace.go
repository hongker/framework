package middleware

import (
	"github.com/hongker/framework/component/trace"
	"github.com/hongker/framework/config"
	"github.com/gin-gonic/gin"
	"strings"
)

// Trace 全局链路中间件
func Trace(c *gin.Context)  {
	// 如果是其他服务的http请求，则从头部信息获取，接续链路
	traceId := strings.TrimSpace(c.GetHeader(config.Server().TraceHeader))
	if traceId == "" {
		traceId = trace.NewUUID()
	}
	trace.Set(traceId)

	c.Next()

	defer trace.GC()
}

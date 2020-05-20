package middleware

import "github.com/gin-gonic/gin"

// FaviconFilter 图标过滤器
func FaviconFilter(ctx *gin.Context) {
	// 如果是请求的/favicon.ico,默认跳过
	if ctx.Request.RequestURI == "/favicon.ico" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

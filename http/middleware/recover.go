package middleware

import "github.com/gin-gonic/gin"

// Recover 处理因系统panic导致的错误，需要拦截，定制化返回响应内容，不能直接输出与系统相关的内容
func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// TODO
		}
	}()
	ctx.Next()
}

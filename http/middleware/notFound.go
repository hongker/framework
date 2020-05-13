package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/http/response"
)

// NotFoundHandler 404
func NotFoundHandler(ctx *gin.Context) {
	response.Wrap(ctx).Error(404, fmt.Sprintf("404 Not Found: %s", ctx.Request.RequestURI))
}

package http

import (
	"fmt"
	"framework/http/response"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer_Start(t *testing.T) {
	server := NewServer()
	// 添加路由
	server.Router.GET("/test", func(ctx *gin.Context) {
		fmt.Println("hello,world")
		response.Wrap(ctx).Success(response.Data{
			"hello":"world",
		})
	})

	server.Start()
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/http/middleware"
	"github.com/hongker/framework/http/response"
	"github.com/hongker/framework/util/secure"
	"github.com/spf13/viper"
	"testing"
)

func TestMain(m *testing.M) {
	viper.SetDefault("local.redis.cluster", "localhost:6379")
	m.Run()
}
func TestServer_Start(t *testing.T) {
	server := NewServer()

	server.Router.Use(middleware.FaviconFilter, middleware.RequestLog)

	server.Router.GET("test", func(context *gin.Context) {
		response.Wrap(context).Success("hello,world")
	})

	secure.Panic(server.Start())
}

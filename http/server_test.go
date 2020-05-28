package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/app"
	"github.com/hongker/framework/config"
	"github.com/hongker/framework/errors"
	"github.com/hongker/framework/http/middleware"
	"github.com/hongker/framework/http/response"
	"github.com/hongker/framework/util/secure"
	"testing"
)

func TestMain(m *testing.M) {
	secure.Panic(config.ReadFromFile("/usr/app.yaml"))
	secure.Panic(app.InitDB())
	m.Run()
}
func TestServer_Start(t *testing.T) {
	server := NewServer()

	server.Router.Use(middleware.Trace, middleware.FaviconFilter, middleware.RequestLog, middleware.Recover)

	server.Router.GET("test", func(context *gin.Context) {
		response.Wrap(context).Success("hello,world")

	})

	server.Router.GET("error", func(context *gin.Context) {
		secure.Panic(errors.New(1001, "some error"))
		response.Wrap(context).Success("hello,world")

	})

	secure.Panic(server.Start(8082))
}

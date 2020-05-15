package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/app"
	"github.com/hongker/framework/component/session"
	"github.com/hongker/framework/http/response"
	"github.com/hongker/framework/util/secure"
	"github.com/spf13/viper"
	"testing"
)

func TestMain(m *testing.M)  {
	viper.SetDefault("local.redis.cluster", "localhost:6379")
	secure.Panic(app.InitRedisCluster())
	m.Run()
}
func TestServer_Start(t *testing.T) {
	server := NewServer()
	store := session.NewCookieStore("userState")

	redisStore, err := session.NewRedisStore()

	secure.Panic(err)
	// 添加路由
	server.Router.GET("/login", func(ctx *gin.Context) {
		fmt.Println("hello,world")

		s , _ := store.Get(ctx.Request, "user")
		s.Values["name"] = ctx.Query("name")
		if err := s.Save(ctx.Request, ctx.Writer); err != nil {
			response.Wrap(ctx).Error(500, "login failed")
			return
		}
		response.Wrap(ctx).Success(response.Data{
			"hello":"world",
		})
	})

	server.Router.GET("/user", func(ctx *gin.Context) {

		s , _ := store.Get(ctx.Request, "user")
		fmt.Println(s.Values)
	})

	server.Router.GET("/redis/login", func(ctx *gin.Context) {
		fmt.Println("hello,world")

		s , _ := redisStore.Get(ctx.Request, "user")
		s.Values["name"] = ctx.Query("name")
		if err := s.Save(ctx.Request, ctx.Writer); err != nil {
			response.Wrap(ctx).Error(500, "login failed")
			return
		}
		response.Wrap(ctx).Success(response.Data{
			"hello":"world",
		})
	})

	server.Router.GET("/redis/user", func(ctx *gin.Context) {

		s , _ := redisStore.Get(ctx.Request, "user")
		fmt.Println(s.Values)
	})

	secure.Panic(server.Start())
}

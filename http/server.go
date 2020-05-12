package http

import (
	"framework/config"
	"framework/http/middleware"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"errors"
)

// Server Web服务管理器
type Server struct {
	// gin的路由
	Router *gin.Engine

	// not found handler
	NotFoundHandler func(ctx *gin.Context)
}

// NewServer 实例化server
func NewServer() *Server {
	router := gin.Default()

	// use global trace middleware
	router.Use(middleware.Trace)

	return &Server{
		Router:          router,
		NotFoundHandler: middleware.NotFoundHandler,
	}
}

// Start run http server
// args must be less than one,
// eg: Start()
//	   Start(8080)
func (server *Server) Start(args ...int) error {
	// 如果给了端口的参数，则读取该参数，否则使用配置的port值
	port := config.Server().Port
	if len(args) == 1 {
		port = args[0]
		config.Server().Port = port
	} else if len(args) > 1 {
		return errors.New("args must be less than one")
	}

	// 404
	server.Router.NoRoute(server.NotFoundHandler)
	server.Router.NoMethod(server.NotFoundHandler)

	completeHost := net.JoinHostPort("", strconv.Itoa(port))

	return server.Router.Run(completeHost)
}
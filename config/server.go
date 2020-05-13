package config

import "github.com/spf13/viper"

// server 服务相关配置
type server struct {
	// 服务名称
	Name string

	// 服务端口
	Port int

	// http超时时间，单位:秒
	HttpRequestTimeOut int

	// 全局ID的header名称
	TraceHeader string
}


// Server 返回服务配置
func Server() *server {
	return &server{
		Name:               viper.GetString(serverNameKey),
		Port:               viper.GetInt(serverPortKey),
		HttpRequestTimeOut: viper.GetInt(httpRequestTimeoutKey),
		TraceHeader: viper.GetString(traceHeaderKey),
	}
}

package config

import "github.com/spf13/viper"

// ReadFromEnvironment 读取环境变量
func ReadFromEnvironment() {
	viper.AutomaticEnv()
}

// ReadFromFile 通过文件读取，返回error
func ReadFromFile(path string) error {
	viper.SetConfigFile(path)

	return viper.ReadInConfig()
}

const(
	// 服务名称配置项
	serverNameKey = "server.name"
	// 服务名称默认值
	defaultServerName = "app"

	// 服务端口配置项
	serverPortKey = "server.port"
	// 服务端口默认值
	defaultServerPort = 8080

	// http请求超时配置项
	httpRequestTimeoutKey = "server.http-request-timeout"
	// http请求超时默认值
	defaultHttpRequestTimeout = 3

	// 全局ID的header配置项
	traceHeaderKey = "server.trace-header"

	// 全局ID的header默认值
	defaultTraceHeader = "request-trace"

)

// init 初始化默认配置
func init()  {
	viper.SetDefault(serverNameKey, defaultServerName)
	viper.SetDefault(serverPortKey, defaultServerPort)
	viper.SetDefault(httpRequestTimeoutKey, defaultHttpRequestTimeout)
	viper.SetDefault(traceHeaderKey, defaultTraceHeader)
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
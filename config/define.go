package config

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


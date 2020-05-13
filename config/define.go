package config

import (
	"fmt"
	"net"
	"strconv"
)

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

// mysql Mysql的配置项
type mysql struct {
	// host
	Host string

	// port, default 3306
	Port int

	// user, default root
	User string

	// password
	Password string

	// 数据库名称
	Name string

	// 最大空闲连接数
	MaxIdleConnections int

	// 最大打开连接数
	MaxOpenConnections int

	// 最长活跃时间
	MaxLifeTime int
}

// Dsn return mysql dsn
func (conf mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)),
		conf.Name)
}

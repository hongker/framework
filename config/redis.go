package config

import (
	goredis "github.com/go-redis/redis"
	"net"
	"strconv"
	"time"
)

// redis redis配置
type redis struct {
	// 地址
	Host string

	// 端口
	Port int

	// 密码
	Auth string

	// 连接池大小
	PoolSize int

	// 最大尝试次数
	MaxRetries int

	// 超时时间
	IdleTimeout time.Duration
}

// Options get redis options
func (conf *redis) Options() *goredis.Options {
	address := net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port))

	return &goredis.Options{
		Addr:        address,
		Password:    conf.Auth,
		PoolSize:    conf.PoolSize,
		MaxRetries:  conf.MaxRetries,
		IdleTimeout: conf.IdleTimeout,
	}
}

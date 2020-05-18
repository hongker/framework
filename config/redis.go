package config

import (
	goredis "github.com/go-redis/redis"
	"github.com/spf13/viper"
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

	// db
	DB int

	// 连接池大小
	PoolSize int

	// 最大尝试次数
	MaxRetries int

	// 超时时间
	IdleTimeout time.Duration
	
	Cluster string
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
		DB: conf.DB,
	}
}

// Redis 获取redis配置
func Redis() *redis {
	return &redis{
		Host:        viper.GetString(GetKeyWithRunMode(redisHost)),
		Port:        viper.GetInt(GetKeyWithRunMode(redisPort)),
		Auth:        viper.GetString(GetKeyWithRunMode(redisPass)),
		DB : viper.GetInt(GetKeyWithRunMode(redisDB)),
		PoolSize:    10,
		MaxRetries:  3,
		IdleTimeout: 3,
		Cluster:     viper.GetString(GetKeyWithRunMode(redisCluster)),
	}
}

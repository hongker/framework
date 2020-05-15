// app 公共服务
// 提供数据库,redis,权限管理，http等公共服务
// 通过依赖注入的设计模式管理全局变量
package app

import (
	"github.com/hongker/framework/config"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"net"
	"net/http"
	"time"
)

var container = dig.New()

var (
	db *gorm.DB
	enforcer *casbin.Enforcer
	redisConn redis.Cmdable
)
// Container 容器
func Container() *dig.Container {
	return container
}

// DB 返回数据库连接
func DB() *gorm.DB {
	return db
}

// Redis 返回redis连接
func Redis() redis.Cmdable {
	return redisConn
}

// PermissionManager 返回rbac权限管理器
func PermissionManager() *casbin.Enforcer {
	return enforcer
}

// Http 返回http连接池客户端
func Http() (client *http.Client)  {
	if err :=  container.Invoke(func(cli *http.Client) {
		client = cli
	}); err != nil {
		client = &http.Client{
			Transport: &http.Transport{ // 配置连接池
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				IdleConnTimeout: time.Duration(config.Server().HttpRequestTimeOut) * time.Second,
			},
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Duration(config.Server().HttpRequestTimeOut) * time.Second,
		}
	}
	return
}

package app

import (
	"fmt"
	"github.com/casbin/casbin/v2/persist"
	"github.com/hongker/framework/component/mysql"
	"github.com/hongker/framework/component/rbac"
	"github.com/hongker/framework/component/redis"
	"github.com/hongker/framework/component/session"
	"github.com/hongker/framework/config"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	mysqlDialectType = "mysql"
)

// InitPermissionManager 初始化权限管理器
func InitPermissionManager(adapter persist.Adapter) error {
	e, err := rbac.NewManagerWithAdapter(adapter)
	if err != nil {
		return nil
	}

	enforcer = e
	return nil
}

// InitDB 初始化DB
func InitDB() error {
	group := config.MysqlGroup()
	if group.Items == nil {
		return fmt.Errorf("mysql config is empty")
	}

	for name, item := range group.Items {
		dataSourceItems := item.DataSourceItems()

		adapter, err := mysql.NewReadWriteAdapter(mysqlDialectType, dataSourceItems)
		if err != nil {
			return err
		}

		adapter.SetMaxIdleConns(item.MaxIdleConnections)
		adapter.SetMaxOpenConns(item.MaxOpenConnections)
		adapter.SetConnMaxLifetime(time.Duration(item.MaxLifeTime) * time.Second)

		conn, err := gorm.Open(mysqlDialectType, adapter)
		if err != nil {
			return err
		}
		dbGroup[name] = conn
	}

	return nil
}

// InitRedis 初始化redis
func InitRedis() error {
	rdb, err := redis.Connect(config.Redis().Options())
	if err != nil {
		return err
	}

	if err := rdb.Ping().Err(); err != nil {
		return err
	}

	redisConn = rdb
	return nil
}

// InitSessionStore 初始化sessionStore
func InitSessionStore() error  {
	conn, err := redis.Connect(config.Redis().SessionOptions())
	if err != nil {
		return err
	}

	store, err := session.NewRedisStoreWithConn(conn)
	if err != nil {
		return err
	}

	sessionStore = store
	return nil
}

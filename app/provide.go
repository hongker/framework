package app

import (
	"fmt"
	"github.com/casbin/casbin/v2/persist"
	"github.com/hongker/framework/component/mysql"
	"github.com/hongker/framework/component/rbac"
	"github.com/hongker/framework/component/redis"
	"github.com/hongker/framework/config"
	"github.com/jinzhu/gorm"
	"time"
	goredis "github.com/go-redis/redis"
)

const(
	mysqlDialectType = "mysql"
)


// InitPermissionManager 初始化权限管理器
func InitPermissionManager(adapter persist.Adapter) error {
	if db == nil {
		return fmt.Errorf("database not init")
	}
	
	e, err := rbac.NewManagerWithAdapter(adapter)
	if err != nil {
		return nil
	}
	
	enforcer = e
	return nil
}

// InitDB 初始化DB
func InitDB() error {
	dataSourceItems := config.Mysql().DataSourceItems()

	adapter, err := mysql.NewReadWriteAdapter(mysqlDialectType, dataSourceItems)
	if err != nil {
		return err
	}


	adapter.SetMaxIdleConns(config.Mysql().MaxIdleConnections)
	adapter.SetMaxOpenConns(config.Mysql().MaxOpenConnections)
	adapter.SetConnMaxLifetime(time.Duration(config.Mysql().MaxLifeTime) * time.Second)

	conn, err := gorm.Open(mysqlDialectType, adapter)
	if err != nil {
		return err
	}

	db = conn

	return nil
}


// InitRedisCluster 初始化redis集群
func InitRedisCluster() error  {
	// TODO 连接redis
	rdb, err := redis.Connect(&goredis.Options{
		Addr: config.Redis().Cluster,
	})
	if err != nil {
		return err
	}

	if err := rdb.Ping().Err(); err != nil {
		return err
	}

	redisConn = rdb
	return nil
}
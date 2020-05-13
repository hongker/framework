package app

import (
	"fmt"
	"github.com/casbin/casbin/v2/persist"
	"github.com/hongker/framework/component/mysql"
	"github.com/hongker/framework/component/rbac"
	"github.com/hongker/framework/config"
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
	dsn := config.Mysql().Dsn()
	conn, err := mysql.Connect(dsn)
	if err != nil {
		return err
	}
	
	db = conn
	return nil
}

func InitRedis() error  {

}
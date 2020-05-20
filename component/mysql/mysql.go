package mysql

import (
	"github.com/hongker/framework/config"
	"github.com/hongker/framework/errors"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect 连接数据库，返回连接与错误
func Connect(dialectType, dsn string) (*gorm.DB, error) {
	conn, err := gorm.Open(dialectType, dsn)

	if err != nil {
		return nil, errors.MysqlConnectFailed("%s", err.Error())
	}

	conf := config.Mysql()
	// 连接池配置
	conn.DB().SetMaxIdleConns(conf.MaxIdleConnections)
	conn.DB().SetMaxOpenConns(conf.MaxOpenConnections)

	// 超时时间配置
	conn.DB().SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)

	return conn, nil
}

package mysql

import (
	"github.com/hongker/framework/errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect 连接数据库，返回连接与错误
func Connect(dialectType, dsn string) (*gorm.DB, error) {
	conn, err := gorm.Open(dialectType, dsn)

	if err != nil {
		return nil, errors.MysqlConnectFailed("%s", err.Error())
	}

	return conn, nil
}

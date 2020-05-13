package config

import (
	"github.com/spf13/viper"
	"net"
	"strconv"
	"fmt"
)

// mysql Mysql的配置项
type mysql struct {
	// host
	Host string

	// 端口号
	Port int

	// 用户名
	User string

	// 密码
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


// Mysql 返回mysql配置
func Mysql() *mysql {
	return &mysql{
		Host:               viper.GetString(dbHostKey),
		Port:               viper.GetInt(dbPortKey),
		User:               viper.GetString(dbUserKey),
		Password:           viper.GetString(dbPasswordKey),
		Name:               viper.GetString(dbNameKey),
		MaxIdleConnections: viper.GetInt(dbMaxIdleConnectionsKey),
		MaxOpenConnections: viper.GetInt(dbMaxOpenConnectionsKey),
		MaxLifeTime:        viper.GetInt(dbMaxLifeTime),
	}
}
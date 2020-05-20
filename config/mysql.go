package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

// mysql Mysql的配置项
type mysql struct {
	// 表连接
	TablePrefix string

	// data sources
	DataSources []DataSource
	// 最大空闲连接数
	MaxIdleConnections int

	// 最大打开连接数
	MaxOpenConnections int

	// 最长活跃时间
	MaxLifeTime int
}

// DataSource 连接配置
type DataSource struct {
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
}

// DataSourceItems 获取全部dsn资源
func (conf *mysql) DataSourceItems() []string {
	var items []string
	for _, dsn := range conf.DataSources {
		items = append(items, dsn.Dsn())
	}
	return items
}

// Dsn return mysql dsn
func (conf DataSource) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)),
		conf.Name)
}

// Mysql 返回mysql配置
func Mysql() *mysql {
	return &mysql{
		DataSources:        dsn(),
		MaxIdleConnections: viper.GetInt(GetKeyWithRunMode(dbMaxIdleConnectionsKey)),
		MaxOpenConnections: viper.GetInt(GetKeyWithRunMode(dbMaxOpenConnectionsKey)),
		MaxLifeTime:        viper.GetInt(GetKeyWithRunMode(dbMaxLifeTime)),
	}
}

// dsn
func dsn() []DataSource {
	var items []DataSource
	if err := viper.UnmarshalKey(GetKeyWithRunMode(dbDataSourcesKey), &items); err != nil {
		return nil
	}
	return items
}

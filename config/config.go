package config

import "github.com/spf13/viper"

// ReadFromEnvironment 读取环境变量
func ReadFromEnvironment() {
	viper.AutomaticEnv()
}

// ReadFromFile 通过文件读取，返回error
func ReadFromFile(path string) error {
	viper.SetConfigFile(path)

	return viper.ReadInConfig()
}

const(
	// 服务名称配置项
	serverNameKey = "server.name"

	// 服务端口配置项
	serverPortKey = "server.port"

	// http请求超时配置项
	httpRequestTimeoutKey = "server.httpRequestTimeout"

	// 全局ID的header配置项
	traceHeaderKey = "server.traceHeader"

	dbHostKey = "db.host"
	dbPortKey = "db.port"
	dbUserKey = "db.user"
	dbPasswordKey = "db.password"
	dbNameKey = "db.name"
	dbMaxIdleConnectionsKey = "db.maxIdleConnections"
	dbMaxOpenConnectionsKey = "db.maxOpenConnections"
	dbMaxLifeTime = "db.maxLifeTime"

)

// init 初始化默认配置
func init()  {
	// 初始化server默认配置
	viper.SetDefault(serverNameKey, "app")
	viper.SetDefault(serverPortKey, 8080)
	viper.SetDefault(httpRequestTimeoutKey, 3)
	viper.SetDefault(traceHeaderKey, "request-trace")

	// 初始化数据库默认配置
	viper.SetDefault(dbHostKey, "127.0.0.1")
	viper.SetDefault(dbPortKey, 3306)
	viper.SetDefault(dbUserKey, "root")
	viper.SetDefault(dbPasswordKey, "")
	viper.SetDefault(dbNameKey, "")
	viper.SetDefault(dbMaxIdleConnectionsKey, 10)
	viper.SetDefault(dbMaxOpenConnectionsKey, 40)
	viper.SetDefault(dbMaxLifeTime, 10)

}


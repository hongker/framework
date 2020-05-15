package config

import (
	"github.com/hongker/framework/component/runmode"
	"github.com/spf13/viper"
	"strings"
)



// ReadFromEnvironment 读取环境变量
func ReadFromEnvironment() {
	viper.AutomaticEnv()
}

// ReadFromFile 通过文件读取，返回error
func ReadFromFile(path string) error {
	// 设置文件
	viper.SetConfigFile(path)

	return viper.ReadInConfig()
}

const(
	// 运行环境，沿用beego的吧，懒得改了
	runModeKey = "server.runmode"
	// 服务名称配置项
	serverNameKey = "server.name"

	// 服务端口配置项
	serverPortKey = "server.port"

	// http请求超时配置项
	httpRequestTimeoutKey = "server.httpRequestTimeout"

	// 全局ID的header配置项
	traceHeaderKey = "server.traceHeader"

	dbMaxIdleConnectionsKey = "db.maxIdleConnections"
	dbMaxOpenConnectionsKey = "db.maxOpenConnections"
	dbMaxLifeTime = "db.maxLifeTime"

	dbDataSourcesKey = "db.dsn"

	// 应用日志名称
	appKey = "app"

	redisHost = "redis.host"
	redisPass = "redis.pass"
	redisPort = "redis.port"
	redisCluster = "redis.cluster"

)

// GetKeyWithRunMode 根据运行环境获取key
func GetKeyWithRunMode(key string) string  {
	return strings.Join([]string{viper.GetString(runModeKey), key}, ".")
}


// init 初始化默认配置
func init()  {
	// 默认环境为本地环境
	viper.SetDefault(runModeKey, runmode.Local)
	// 初始化server默认配置
	viper.SetDefault(serverNameKey, "app")
	viper.SetDefault(serverPortKey, 8080)
	viper.SetDefault(httpRequestTimeoutKey, 3)
	viper.SetDefault(traceHeaderKey, "request-trace")

	// 初始化数据库默认配置
	//viper.SetDefault(dbHostKey, "127.0.0.1")
	//viper.SetDefault(dbPortKey, 3306)
	//viper.SetDefault(dbUserKey, "root")
	//viper.SetDefault(dbPasswordKey, "")
	//viper.SetDefault(dbNameKey, "")
	//viper.SetDefault(dbMaxIdleConnectionsKey, 10)
	//viper.SetDefault(dbMaxOpenConnectionsKey, 40)
	//viper.SetDefault(dbMaxLifeTime, 10)

}

// UnmarshalAppConfig 解析app配置
func UnmarshalAppConfig(obj interface{}) error {
	return viper.UnmarshalKey(appKey, obj)
}


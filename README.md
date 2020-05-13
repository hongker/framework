# framework
提供快捷的运行web项目的golang脚手架

## 安装
暂时使用github放置，后续转成私有
```
go get github.com/hongker/framework
```
## 组件
- 支持web服务,定时任务
- 支持redis,mysql,elastics
- 支持读取文件配置，环境变量配置等
- 支持全局链路追踪，跨域，权限校验(rbac)等中间件
- 支持日志，事件,http请求等常用组件
- 支持swagger接口文档集成

### Web服务
基于`github.com/gin-gonic/gin`实现的高性能，可扩展的web服务
- 开始使用
```go
func main() {
    server := http.NewServer()
    // 添加路由
	server.Router.GET("/test", func(ctx *gin.Context) {
		response.Wrap(ctx).Success(response.Data{
			"hello":"world",
		})
	})
    // 启动web服务。启动端口是读取的配置文件中的server.port参数，如果不配置，默认是8080
    server.Start()
    // 也可以指定8081端口启动
    // server.Start(8081)
}
```
用法都很简单，想要了解更多请查看:[gin文档](https://github.com/gin-gonic/gin)

### 中间件
使用中间件的方式嵌入一些非业务的逻辑，使得层次更为清晰，简单。   
注：如果需要中间件在执行路由前执行，那么中间件的引入需要在定义路由之前。   
- Trace 全局链路追踪中间件   
一个用户发起的请求，会有唯一的全局traceId.该请求里执行的所有业务操作都可以用该traceId来关联。比如记录日志，发起http请求等，通过该ID形成一条完整的业务链路。   
web服务默认引入了该组件的，生效位置：响应内容的`meta.trace_id`，日志里的`trace_id`,http请求的header头部:`trace_id`   
该组件默认已引入了。

- CORS 跨域中间件
```go
server.Router.Use(middleware.CORS)
```

- Recover 错误处理中间件   
类似于`try .. catch`里的`catch`,拦截panic错误，避免系统错误直接暴露给用户。
```go
server.Router.Use(middleware.Recover)
```
- Permission 权限   
基于rbac的权限校验中间件
```go
// 可以给全部路由设置权限校验
server.Router.Use(middleware.Permission)
// 也可以给某个路由组设置权限校验，代表该组下的所有路由都需要进行权限校验
server.Router.Group("user").Use(middleware.Permission)
// 也可以给某个路由设置权限校验
server.Router.Get("/money", handler.MoneyHandler).Use(middleware.Permission)
```


### Swagger接口文档生成
集成了`github.com/swaggo/gin-swagger`,通过注解自动生成接口文档，使用方式请查看:[gin-swagger文档](https://github.com/swaggo/gin-swagger)

### 配置
- 加载配置
```go
package main
func init() {
    // 一般来说配置文件的加载都是放在main.go里的init函数里执行
    // 两种方式可以二选一，也可以两个都用
    config.ReadFromEnvironment() // 从环境变量中读取配置
    err := config.ReadFromFile(configFilePath) // 通过文件读取配置.如果加载失败，建议中断启动:os.Exit(-1)
}
func main() {
    
}

```

- 初始化配置
通过init方法指定配置的默认值
```go
import (
 "github.com/spf13/viper"
)

func init() {
    // 如果配置文件或者环境变量里没有主动设置someKey这个配置，那么someValue就是someKey这个配置项的默认值
    viper.SetDefault("someKey", "someValue")
}
```

- 读取配置
```go
// 简单的读取
someConfig := viper.GetString("someKey")
// 读取int
someIntConfig := viper.GetInt("someIntKey")
```

更多请查看：[viper文档](https://github.com/spf13/viper)

### 基础服务
基于`github.com/uber/dig`的依赖注入模式，通过统一的`app`模块管理如mysql连接，redis连接等全局变量。
#### 数据库
集成业界好评且强大的`github.com/jinzhu/gorm`,实现对mysql数据的操作。
#### Redis

#### Elastic

### 日志
基于`github.com/sirupsen/logrus`实现的日志组件
```go
func main() {
    log.Info("message", log.Content{
        "content":"this is log content",
        "params": "some other params",
    })
    log.Info("error message", log.Content{
        "content":"this is log content",
        "err": someError.Error(),
    })
}
```
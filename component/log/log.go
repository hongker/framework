package log

import (
	"github.com/hongker/framework/component/trace"
	"github.com/hongker/framework/config"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	//logrus.SetLevel(logrus.WarnLevel)
}

// Content 日志内容
type Content map[string]interface{}

// log 返回日志实例
func log() *logrus.Entry {
	return logrus.WithField("system", Content{
		"trace_id" : trace.Get(),
		"system_name" : config.Server().Name,
	})
}

// Info 信息
func Info(msg string, content Content)  {
	log().WithField("content", content).Info(msg)
}

// Error 错误
func Error(msg string, content Content)  {
	log().WithField("content", content).Error(msg)
}

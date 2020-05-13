package log

import (
	"github.com/hongker/framework/component/trace"
	"testing"
)

func TestMain(m *testing.M)  {
	m.Run()
}
// TestInfo 测试Info
func TestInfo(t *testing.T) {
	trace.Set(trace.NewUUID())
	defer trace.GC()
	Info("test", Content{
		"id": 1,
	})
}

// TestInfo 测试Error
func TestError(t *testing.T) {
	trace.Set(trace.NewUUID())
	defer trace.GC()
	Error("test", Content{
		"id": 1,
	})
}

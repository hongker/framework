// trace 基于协程的服务全局链路追踪组件
// 一般在http的中间件，定时任务里初始化,业务里仅仅需要GET即可
// 使得同一个协程下的所有业务都使用一个ID标记，用于追踪业务执行链路
// 可用于日志分析，用户请求溯源
// Usage:
// 		trace.Set(trace.NewUUID())
//  	defer trace.GC() // 注：使用defer延迟调用，表示最后执行id的回收释放容量，否则容易导致内存溢出
//		fmt.Println(trace.Id())
package trace

import (
	"github.com/hongker/framework/util/strings"
	"github.com/petermattis/goid"
	"sync"
)

const (
	// 前缀，用于区分id类型
	tracePrefix = "trace:"
)

var (
	// 保存协程的全局唯一ID数组
	traceIds = map[int64]string{}

	// 读写锁，保证线程安全
	rwm sync.RWMutex
)

// NewUUID 新生成一个全局的唯一ID
func NewUUID() string {
	return tracePrefix + strings.UUID()
}

// Set 设置全局唯一ID,协程安全
func Set(id string) {
	goID := grtId()

	// 通过读写锁保证协程安全
	rwm.Lock()
	defer rwm.Unlock()

	traceIds[goID] = id
}

// Get 获取当前协程的唯一ID，如果没有设置当前协程的唯一ID，会返回空字符串
func Get() string {
	goID := grtId()

	// 读锁不会阻塞，提高读取的性能
	rwm.RLock()
	defer rwm.RUnlock()

	return traceIds[goID]
}

// GC 回收唯一ID,释放map容量
func GC() {
	goID := grtId()
	rwm.Lock()
	defer rwm.Unlock()

	delete(traceIds, goID)
}

// grtId 获取协程ID
func grtId() int64 {
	return goid.Get()
}

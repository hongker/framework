package bytes

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

// Adapter buffer pool
type Adapter struct {
	pool sync.Pool
}

const (
	// 默认长度
	adapterBufferPoolSize = 4096
)

// NewAdapter 生成一个bytes适配器
func NewAdapter() *Adapter {
	return &Adapter{pool: sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(make([]byte, adapterBufferPoolSize))
	}}}
}

// Read 读取流数据
func (adapter *Adapter) Read(reader io.Reader) (string, error) {
	buffer := adapter.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			adapter.pool.Put(buffer)
			buffer = nil
		}
	}()
	_, err := io.Copy(buffer, reader)

	if err != nil {
		return "", fmt.Errorf("failed to read respone:%s", err.Error())
	}

	return buffer.String(), nil
}

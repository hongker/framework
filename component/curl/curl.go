// curl 提供http请求执行组建
// 通过Execute方法执行http请求，返回解析后的响应内容
// Usage:
//		request,_ := http.NewRequest(http.MethodGet, "http://baidu.com", nil)
//		resp, err := Execute(request)
package curl

import (
	"errors"
	"fmt"
	"framework/app"
	"framework/component/trace"
	"framework/config"
	"framework/util/bytes"
	"net/http"
)

// DefaultAdapter
var adapter = bytes.NewAdapter()

// Execute 执行http请求，并返回解析后的响应内容字符串和错误信息
func Execute(req *http.Request) (string, error) {
	// 设置全局链路ID
	req.Header.Set(config.Server().TraceHeader, trace.Get())

	resp, err := app.Http().Do(req)
	if err != nil {
		return "", err
	}

	if resp == nil {
		return "", errors.New("no response")
	}

	// 关闭response
	defer func() {
		_ = resp.Body.Close()
	}()

	// 判断验证码是否正确
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response status code is:%d", resp.StatusCode)
	}

	// 通过buffer池读取response的内容，避免如因io.ReadAll的内存溢出问题
	return adapter.Read(resp.Body)
}


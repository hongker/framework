// curl 提供http请求执行组建
// 通过Execute方法执行http请求，返回解析后的响应内容
// Usage:
//		request,_ := http.NewRequest(http.MethodGet, "http://baidu.com", nil)
//		resp, err := Execute(request)
package curl

import (
	b "bytes"
	"errors"
	"fmt"
	"github.com/hongker/framework/app"
	"github.com/hongker/framework/component/trace"
	"github.com/hongker/framework/config"
	"github.com/hongker/framework/util/bytes"
	"github.com/hongker/framework/util/json"
	"io"
	"net/http"
)

// DefaultAdapter
var adapter = bytes.NewAdapter()

// NewRequest 返回原生httpRequest实例
// Usage:
//	getRequest, err := curl.NewRequest("GET", "http://www.baidu.com?id=1", nil)
func NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

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

// Get 发送get请求
// url: 请求地址
// headers: 头部信息，比如: Content-Type:application/json，或者自定义的token字段等等。如果不需要传nil
func Get(url string, headers map[string]string) (string, error)  {
	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// 设置头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return Execute(req)
}

// Post 发送post请求
// url: 请求地址
// data: 请求参数，可以是数组或者对象
// headers: 头部信息，比如: Content-Type:application/json，或者自定义的token字段等等。如果不需要传nil
func Post(url string, data interface{}, headers map[string]string) (string, error) {
	jsonStr, err := json.Encode(data)
	if err != nil {
		return "", err
	}

	req, err := NewRequest(http.MethodPost, url, b.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return "", err
	}

	// 设置头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}


	return Execute(req)
}


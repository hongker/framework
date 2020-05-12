package strings

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"unsafe"
)

// UUID 返回唯一ID
func UUID() string {
	return uuid.NewV4().String()
}

// Implode 根据连接符号连接数组
func Implode(items interface{}, separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(items), "[]"), " ", separator, -1)
}

// Explode 分割字符词，返回字符词数组
func Explode(str, separator string) []string {
	return strings.Split(str, separator)
}

// ToByte 字符串转字节
func ToByte(s string)  []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Default 返回默认值
func Default(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}
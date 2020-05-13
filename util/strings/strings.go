package strings

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"unsafe"
)

//Md5 return the encrypt string by md5 algorithm
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash return hash string
func Hash(s string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(s))
	result := Sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}

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

// ToInt 字符串转int
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// Default 返回默认值
func Default(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}
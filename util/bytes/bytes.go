package bytes

import "unsafe"

// ToString 字节转字符串
func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

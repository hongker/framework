package errors

import (
	"fmt"
	"github.com/hongker/framework/util/json"
	"net/http"
)

// Error
type Error struct {
	// error code
	Code int `json:"code"`
	// error message
	Message string `json:"message"`
}

const (
	MysqlConnectFailedCode = 1001
	RedisConnectFailedCode = 1002
)

// Error strings
func (e *Error) Error() string {
	s, _ := json.Encode(e)
	return s
}

// New
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Parse tries to parse a JSON strings into an error. If that
// fails, it will set the given strings as the error detail.
func Parse(errStr string) *Error {
	e := new(Error)

	if err := json.Decode([]byte(errStr), e); err != nil {
		e.Code = http.StatusInternalServerError
		e.Message = err.Error()
	}
	return e
}

// Unauthorized generates a 401 error.
func Unauthorized(format string, v ...interface{}) *Error {
	return New(http.StatusUnauthorized, fmt.Sprintf(format, v...))
}

// Forbidden generates a 403 error.
func Forbidden(format string, v ...interface{}) *Error {
	return New(http.StatusForbidden, fmt.Sprintf(format, v...))
}

// NotFound generates a 404 error.
func NotFound(format string, v ...interface{}) *Error {
	return New(http.StatusNotFound, fmt.Sprintf(format, v...))
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(format string, v ...interface{}) *Error {
	return New(http.StatusMethodNotAllowed, fmt.Sprintf(format, v...))
}

// Timeout generates a 408 error.
func Timeout(format string, v ...interface{}) *Error {
	return New(http.StatusRequestTimeout, fmt.Sprintf(format, v...))
}

// InternalServerError generates a 500 error.
func InternalServer(format string, v ...interface{}) *Error {
	return New(http.StatusInternalServerError, fmt.Sprintf(format, v...))
}

// MysqlConnectFailed
func MysqlConnectFailed(format string, v ...interface{}) *Error {
	return New(MysqlConnectFailedCode, fmt.Sprintf(format, v...))
}

// RedisConnectFailed
func RedisConnectFailed(format string, v ...interface{}) *Error {
	return New(RedisConnectFailedCode, fmt.Sprintf(format, v...))
}

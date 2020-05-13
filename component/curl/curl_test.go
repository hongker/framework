package curl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestExecute(t *testing.T) {
	request, err := NewRequest(http.MethodGet, "http://baidu.com", nil)
	assert.Nil(t, err)
	assert.NotNil(t, request)
	resp, err := Execute(request)
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestGet(t *testing.T) {
	resp, err := Get("http://baidu.com", nil)
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestPost(t *testing.T) {
	data := map[string]interface{}{
		"id": 1,
		"name":"test01",
	}
	resp, err := Post("http://baidu.com", data, map[string]string{"Content-Type":"application/json"})
	assert.Nil(t, err)
	fmt.Println(resp)

}

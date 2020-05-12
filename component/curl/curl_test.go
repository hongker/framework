package curl

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExecute(t *testing.T) {
	request,_ := http.NewRequest(http.MethodGet, "http://baidu.com", nil)
	resp, err := Execute(request)
	fmt.Println(resp, err)
}

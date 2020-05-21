package response

import (
	"fmt"
	"testing"
)

func TestPool_Get(t *testing.T) {
	p := newPool()
	r := p.Get()
	r.Code = 200
	fmt.Println(r)
}

func BenchmarkPool_Get(b *testing.B) {
	p := newPool()
	for i:=0;i<b.N;i++ {
		r := p.Get()
		r.Code = 200
	}
}

func BenchmarkNewResponse(b *testing.B)  {
	for i:=0;i<b.N;i++ {
		r := newResponse(200, "1")
		r.Code = 200
	}
}
package trace

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// TestTrace
func TestTrace(t *testing.T)  {
	uuid := NewUUID()
	Set(uuid)
	defer GC()

	assert.Equal(t, uuid, Get())
}

// BenchmarkGet Get压测
func BenchmarkGet(b *testing.B) {
	uuid := NewUUID()
	Set(uuid)
	defer GC()
	for i:=0; i<b.N;i++ {
		Get()
	}
}

// BenchmarkSet 压测Set
func BenchmarkSet(b *testing.B) {
	for i:=0; i<b.N;i++{
		uuid := NewUUID()
		Set(uuid)
	}
	defer GC()
}
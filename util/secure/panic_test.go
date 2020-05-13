package secure

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	var err error
	// 不会触发panic
	Panic(err)

	err = fmt.Errorf("test")
	// 会触发panic
	Panic(err)
}

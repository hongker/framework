package number

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i:=0;i<100;i++ {
		if n := RandInt(1,3); n != 1 {
			fmt.Println(n)
		}
	}
}
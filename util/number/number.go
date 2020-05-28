package number

import (
	"math"
	"math/rand"
	"time"
)

//Min return min number
func Min(a, b int) int {
	if a < b {
		return a
	}
	if b < a {
		return b
	}
	return a
}

//Max return max number
func Max(a, b int) int {
	if a > b {
		return a
	}
	if b > a {
		return b
	}
	return a
}

// Div return a/b
func Div(a, b int) int {
	if b == 0 || a == 0 {
		return 0
	}

	return int(math.Ceil(float64(a) / float64(b)))
}

// Round rounding-off method
func Round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// RandInt 取随机数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// DefaultInt
func DefaultInt(val , defaultVal int) int {
	if val == 0 {
		return defaultVal
	}

	return val
}
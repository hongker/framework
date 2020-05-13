package secure

// Panic 检查error,如果err不为nil,才panic,可以有效的减少业务代码里过多if判断
func Panic(err error)  {
	if err != nil {
		panic(err)
	}
}

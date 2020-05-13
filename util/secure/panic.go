package secure

// Panic 检查error
func Panic(err error)  {
	if err != nil {
		panic(err)
	}
}

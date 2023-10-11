package memory

func MyFunc() int {
	var fn func() int
	fn = func() int { return 42 }
	return fn() + 1
}

package memory

func Inc(a [3]uint16) [3]uint16 {
	var res [3]uint16

	res[0] = a[0] + 1
	res[1] = a[1] + 1
	res[2] = a[2] + 1

	return res
}

//go:noinline
func Inc3(a int) int {
	return a + 1
}

func Inc2(a [3]uint16) [3]uint16 {
	var res [3]uint16

	res[0] = a[0] + 2
	res[1] = a[1] + 2
	res[2] = a[2] + 2

	return res
}

func IncPtr(a *[3]uint16) {
	for i := 0; i < len(a); i++ {
		(*a)[i] = (*a)[i] + 1
	}

	return
}

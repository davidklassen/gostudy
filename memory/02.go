package memory

func Assignment() {
	println("Assignment:")

	var a int
	println("a value:", a, "- a address:", &a)
	a = 42
	println("a value:", a, "- a address:", &a)
	b := a
	println("b value:", b, "- b address:", &b)
	b = a + 1
	println("b value:", b, "- b address:", &b)
	{
		// variable shadowing
		a := a
		println("masking a value:", a, "- masking a address:", &a)
	}

	println("---")
}

func Address() {
	println("Address()")

	a := 42
	println("a value:", a, "- a address:", &a)
	b := 24
	println("b value:", b, "- b address:", &b)
	c := &a
	println("c value:", c, "- c address:", &c)
	c = &b
	println("c value:", c, "- c address:", &c)

	println("---")
}

func Dereference() {
	println("Dereference()")

	a := 42
	println("a value:", a, "- a address:", &a)
	b := &a
	println("b value:", b, "- b address:", &b, "- *b value:", *b)
	c := &a
	println("c value:", c, "- c address:", &c, "- *c value:", *c)
	a += 1
	println("c value:", c, "- c address:", &c, "- *c value:", *c)
	d := &b
	println("d value:", d, "- d address:", &d, "- *d value:", *d, "- **d value:", **d)

	println("---")
}

func IndirectAssignment() {
	println("IndirectAssignment()")

	a := 42
	println("a value:", a, "- a address:", &a)
	b := &a
	println("b value:", b, "- b address:", &b)
	*b = 24
	println("b value:", b, "- b address:", &b)
	println("a value:", a, "- a address:", &a)

	println("---")
}

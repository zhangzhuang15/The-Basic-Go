package main

const (
	a = 1
	b
	c = iota
	d
	e = 20
)

func main() {
	println("a: ", a)
	println("b: ", b)
	println("c: ", c)
	println("d: ", d)
	println("e: ", e)
}

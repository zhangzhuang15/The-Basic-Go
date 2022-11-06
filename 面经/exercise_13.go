package main

import "fmt"


func main() {

	a := new(struct{})
	b := new(struct{})

	c := new(struct{})
	d := new(struct{})

	// WTF! a == b  is false 😱
	println(a, b, a == b)

	fmt.Println(c, d, c == d)

}

// why ?
//
// a and b not escape from the stack to heap, and then
// compiler make optimization, it finds a and b are two
// variable.
//
// c and d is as input of fmt.Println, yes, as you can guess,
// c and d escape from stack to heap !
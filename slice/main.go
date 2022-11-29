package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func exercise1() {
	s := make([]int, 3)
	s[0], s[1], s[2] = 0, 1, 2

	var d []int

	copy(d, s)

	// Amazing! Nothing
	fmt.Println("d: ", d)
}

func exercise2() {
	s := make([]int, 3, 10)
    s[0], s[1], s[2] = 0, 1, 2

	b := append(s, 3)
	d := append(s, 4)

	fmt.Println("b: ", b) // [0, 1, 2, 4]
	fmt.Println("d: ", d) // [0, 1, 2, 4]
	fmt.Println("s: ", s) // [0, 1, 2]
}

func exercise3() {
	var innerAppend func([]int, int)
	var innerAppendPointer func(*[]int, int)

	innerAppend = func(src []int, v int) {
		src = append(src, v)
	}

	innerAppendPointer = func(src *[]int, v int) {
        *src = append(*src, v)
	}

	s := make([]int, 3)
	
	s[0], s[1], s[2] = 0, 1, 2

	innerAppend(s, 4)
	fmt.Println("s: ", s) // [0, 1, 2]

	innerAppendPointer(&s, 5)
	fmt.Println("s: ", s) // [0, 1, 2, 5]
}

func exercise4() {
	s := make([]int, 3, 4)
	s[0], s[1], s[2] = 0, 1, 2

	b := s[1: 4]
	b[2] = 10

	// b 和 s 公用一个底层数组

	fmt.Println("s: ", s)
    fmt.Println("b: ", b)

	sort.Slice(s, func(i, j int) bool { return s[i] > s[j]})

	fmt.Println("s: ", s)
	fmt.Println("b: ", b)

	s = append(s, 5)

	fmt.Println("s: ", s)
	fmt.Println("b: ", b)

	b = append(b, 11, 12, 13)

	fmt.Println("s: ", s)
	fmt.Println("b: ", b)

	fmt.Println("s: ", unsafe.Pointer(&s[0]))
	fmt.Println("b: ", unsafe.Pointer(&b[0]))
	
}

func main() {
	exercise4()
}
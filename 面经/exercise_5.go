package main

import "fmt"

func main() {
	fmt.Println(func() {} == func() {})
}

// func() {} 函数之间无法比较，函数只能和 nil 比较！

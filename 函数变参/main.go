package main

import "fmt"

// Javascript定义函数变参用 ...args，
// 将一个列表拆解成函数变参也是 ...args

// python定义函数变参用 *args定义，
// 将一个列表拆解成函数变参也是 *args

// go语言定义函数变参用 ...类型,
// 将一个切片拆解成函数变参是 args...
// NOTE: ...的位置不同！

func run(args ...interface{}) {
	fmt.Println("args: ", args)
}

func main() {
	var args = []interface{}{5, "Mock", 111.002}

	run(2, "Jack", 200.02)

	run(args...)
}

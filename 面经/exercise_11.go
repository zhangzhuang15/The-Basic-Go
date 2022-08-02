package main

import "fmt"

func main() {

	var name string = "Jack"

	fmt.Printf("name address: %p\n", &name)

	name, ok := "Peter", true

	if ok {
		fmt.Printf("name address: %p\n", &name)
	}

}

// 两个 name 的地址一样嘛？
// 当心哦，这是 go 的一个坑

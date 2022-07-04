package main

import "fmt"

// 一个struct数据直接赋值给另一个struct类型的变量，
// 会发生什么呢？
// 会是深拷贝嘛？

type Man struct {
	name string
	age  int
	son  *Man
}

func main() {
	son := Man{
		"Jack",
		10,
		nil,
	}

	physicalFather := Man{
		"Peter",
		34,
		&son,
	}

	lawFather := physicalFather

	fmt.Println("age value equal ? ", lawFather.age == physicalFather.age)
	fmt.Println("age address equal ? ", &lawFather.age == &physicalFather.age)
	fmt.Println("son value equal ? ", lawFather.son == physicalFather.son)
	fmt.Println("son address equal ? ", &lawFather.son == &physicalFather.son)
}

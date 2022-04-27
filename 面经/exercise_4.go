package main

import "fmt"

type Man struct {
	name string
}

func (man *Man)Tell() {
	if man == nil {
		fmt.Println("name is empty")
	} else {
		fmt.Println(man.name)
	}
}

func main() {
	// m == nil , 调用 Tell 不会报错嘛？
	var m *Man
	var n Man = Man{ "jack" }

	m.Tell()
	n.Tell()
}

// func (man *Man)Tell() {}
// 等效于
// func Tell(man *Man) {}

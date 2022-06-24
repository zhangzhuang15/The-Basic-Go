package main

import "fmt"

func main() {
	m := [...]int {
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m))
}

// TODO: 输出多少呢？是3嘛？
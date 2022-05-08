package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"Tom": 11}

	if v, ok := m["Tom"]; ok {
		fmt.Println("m['Tom'] exists: ", v)
	}

	v := m["Tom"]
	fmt.Println("m['Tom']: ", v)

	// 不存在的key，会返回零值
	v = m["Jim"]
	fmt.Println("m['Jim']: ", v)

	// map中不存在的key，返回零值，在这里就是 0，
	// 可如果这个 key 存储的就是 0，
	// 我们将无法断定这个key是否存在。
	// 因此要采用 v, ok := m[key] 的形式判断 key是否存在
}

package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 3, 5, 6}

	for key := range a {
		fmt.Printf("%d\t", key)
	}

	fmt.Println()

	for _, value := range a {
		fmt.Printf("%d\t", value)
	}

	fmt.Println()

	// 你也可以不用填满数据, 空缺的默认用零值填充
	aa := [4]int{2, 3}
	for key, value := range aa {
		fmt.Printf("%d\t%d\n", key, value)
	}

	fmt.Println()

	// 当你不想自己计算数组中究竟有多少个元素时，可以这样写
	b := [...]int{1, 4, 9, 10}
	for k := range b {
		fmt.Printf("%d\t", k)
	}

	fmt.Println()

	// 当你想初始化一个非常大的数组时，可以这样搞
	c := [...]int{3: 6, 99: 100}
	for key, value := range c {
		fmt.Printf("%d %d\n", key, value)
	}
}

// 数组的玩法 get 到了嘛？
// 注意哦 []int 是切片类型！

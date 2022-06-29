package main

import "fmt"

func main() {
	c1 := make(chan int, 1)

	select {
	// c1 管道中没有数据，就会阻塞
	// NOTE: 如果不想有阻塞，需要设置default
	case value := <- c1:
		fmt.Println(value)
	}
}
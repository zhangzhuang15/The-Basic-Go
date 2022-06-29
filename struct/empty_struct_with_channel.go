package main

import (
	"fmt"
	"time"
)

// struct{} 类型数据不占据内存，
// 配合 channel 使用，可以作为一种信号，打破阻塞或者引发阻塞

func main() {
	var ch = make(chan struct{})

	go func() {
		fmt.Println("child go routine gets blocked")
		<-ch
		fmt.Println("child go routine continues")
	}()

	time.Sleep(5 * time.Second)
	ch <- struct{}{}
	time.Sleep(2 * time.Second)
	fmt.Println("main go routine ends")
}

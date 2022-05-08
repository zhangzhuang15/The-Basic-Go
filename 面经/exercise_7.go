package main

import (
	"context"
	"fmt"
	"time"
)

// goroutine 泄漏的例子

func main() {
	ch := make(chan struct{})

	// timeout 只有 1s
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	go func(){
		time.Sleep(2 * time.Second)
		// 这里发生阻塞
		ch <- struct{}{}
		fmt.Println("goroutine over")
	}()

	select {
	case <- ch:
		fmt.Println("ok")
	case <- ctx.Done(): // 这一条会先被执行
		fmt.Println("timeout")
	}

	time.Sleep(5 * time.Second)
}

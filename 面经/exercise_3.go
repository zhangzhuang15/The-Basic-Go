package main

import "fmt"

// import "C"


func main() {
	var ch chan struct{}
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
		}
	}()
	<-ch
}

// 会出现 panic 嘛？

// 打开第5行的注释会发生 panic 嘛？

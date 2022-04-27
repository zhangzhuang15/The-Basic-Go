package main

import "fmt"

func main() {

	defer func(){
		fmt.Println("panic message: ", recover())
	}()
	panic("hello world")
}

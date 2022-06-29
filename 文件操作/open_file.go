package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("demo.txt", os.O_RDONLY | os.O_CREATE, 0664)
	defer file.Close()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("file name: ", file.Name())
}

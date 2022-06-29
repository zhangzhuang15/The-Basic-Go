package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("demo.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("file name: ", file.Name())
}

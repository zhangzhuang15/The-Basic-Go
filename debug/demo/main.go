package main

import "fmt"

func main() {
	s := []int {1, 3, 5, 3, 4, 2}

	for item := range s {
		fmt.Println("value: ", item)
	}
}

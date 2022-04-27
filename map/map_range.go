package main

import "fmt"

func main() {
	var m map[string]int = map[string]int { "Jack": 1, "Jason": 17, "Tom": 25, "Cook": 34, "Apple": 29}

	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}

package main

import "fmt"

func main() {
	s := make([]int, 1)
	s = append(s, 1)
	fmt.Println("s[0]: ", s[0])
}

// s[0]是多少呢？

package main

import (
	"fmt"
	u "jack/utils"

	pete "github.com/peter"
)

func main() {
	var a int = 10
	var b int = 30

	fmt.Printf("result: %d\n", u.Add(a, b))

	pete.Hello()
}

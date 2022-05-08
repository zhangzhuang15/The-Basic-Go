package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for index := 0; index < 10; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("output: ", index)
		}()
	}
	wg.Wait()
}

// 输出的都是 10， 为什么？
// 如何解决呢？



// 修改为 go func(value) {
//       ....
//      }(index)

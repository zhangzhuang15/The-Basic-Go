package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	fmt.Print("A, ")
	m.Lock()

	go func() {
		time.Sleep(2000 * time.Millisecond)
		m.Unlock()
	}()

	// block here until m.Unlock() is invoked
	m.Lock()
	fmt.Print("B ")
}

// so, what's output?

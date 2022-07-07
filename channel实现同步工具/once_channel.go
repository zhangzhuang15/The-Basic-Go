package main

import (
	"fmt"
	"time"
)

// Once用来确保同一个函数在多个协程中只执行一次
type Once chan struct{}

func NewOnce() Once {
	once := make(chan struct{}, 1)
	once <- struct{}{}
	return once
}

func (once Once) Do(f func()) {
	_, ok := <-once

	// once close之后，ok取false
	if !ok {
		return
	}

	f()
	close(once)
}

func main() {
	once := NewOnce()

	f := func() {
		fmt.Println("hello world")
	}

	go func() {
		once.Do(f)
	}()

	go func() {
		once.Do(f)
	}()

	time.Sleep(2 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

// 信号量实现

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	s <- struct{}{}
}

func (s Semaphore) UnLock() {
	<- s
}

func main() {
	s := NewSemaphore(3)

	for i := 0; i < 2; i += 1 {
		go func(id int){
			defer s.UnLock()
			s.Lock()
			fmt.Printf("child routine %d\n", id)
		}(i)
	}

	defer s.UnLock()
	s.Lock()
    fmt.Println("main goroutine")
	time.Sleep(800 * time.Millisecond)
}
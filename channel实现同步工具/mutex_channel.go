package main

import (
	"fmt"
	"time"
)

// 独占锁实现

type Mutex chan struct{}

func NewMutex() Mutex {
	return make(Mutex, 1)
}

func (m Mutex) Lock() {
	m <- struct{}{}
}

func (m Mutex) UnLock() {
	<-m
}

func (m Mutex) TryLock(t time.Duration) bool {
	timer := time.NewTimer(t)
	select {
	case m <- struct{}{}:
		return true
	case <-timer.C:
		return false
	}
}

func main() {
	m := NewMutex()

	go func() {
		defer m.UnLock()
		m.Lock()
		fmt.Println("child goroutine")
	}()

	defer m.UnLock()
	fmt.Println("main goroutine wanna say something")
	m.Lock()
	fmt.Println("main goroutine say: hello world")
}

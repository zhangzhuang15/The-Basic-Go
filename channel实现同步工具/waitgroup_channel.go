package main

import (
	"fmt"
)


type WaitGroup struct {
	mutex chan struct{}
	waiters chan int
}

func NewWaitGroup() WaitGroup {
	g := WaitGroup{
		mutex: make(chan struct{}, 1),
		waiters: make(chan int, 1),
	}

	g.waiters <- 0

	return g
}

func (waitGroup WaitGroup) Add(waiters int) {
	if waiters < 0 {
		panic("waiters must >= 0!")
	}
	_waiters := <- waitGroup.waiters
	_waiters += waiters
	waitGroup.waiters <- _waiters
}

func (waitGroup WaitGroup) Done() {
	waiters := <- waitGroup.waiters
	if waiters == 0 {
		waitGroup.waiters <- 0
		return
	}
	waiters--
	waitGroup.waiters <- waiters
	if waiters == 0 {
		waitGroup.mutex <- struct{}{}
	}
}

func (waitGroup WaitGroup) Wait() {
	waiters := <- waitGroup.waiters
	waitGroup.waiters <- waiters
	if waiters == 0 {
		return
	}
	<- waitGroup.mutex
}


func main() {
	g := NewWaitGroup()

	g.Add(3)

	for i := 0; i < 3; i++ {
		go func(id int){
			defer g.Done()
			fmt.Printf("child routine %d\n", id)
		}(i)
	}

	g.Wait()

	fmt.Println("main routine")
}
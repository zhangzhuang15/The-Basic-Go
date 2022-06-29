package main

import (
	"fmt"
	"time"
)

// 读写锁实现

type RWMutex struct {
	write chan struct{}
	readers chan int
}

func NewRWMutex() RWMutex {
	return RWMutex{
		write: make(chan struct{}, 1),
		readers: make(chan int, 1),
	}
}

func (rwMutex RWMutex) Lock() {
	rwMutex.write <- struct{}{}
}

func (rwMutex RWMutex) UnLock() {
	<- rwMutex.write
}

func (rwMutex RWMutex) RLock() {
	var rs int

	select {
	case rwMutex.write <- struct{}{}:
	case rs = <- rwMutex.readers:
	}

	rs++
	rwMutex.readers <- rs
}

func (rwMutex RWMutex) RUnLock() {
	rs := <- rwMutex.readers
	rs--
	if rs == 0 {
		<- rwMutex.write
		return
	}
	rwMutex.readers <- rs
}


func main() {
	rwMutex := NewRWMutex()

	go func(){
		defer rwMutex.UnLock()
		rwMutex.Lock()
		fmt.Println("special write goroutine")
		time.Sleep(2000 * time.Millisecond)
	}()

	for i := 0; i < 4; i++ {
		go func(id int){
			defer rwMutex.RUnLock()
			rwMutex.RLock()
			fmt.Printf("child goroutine %d\n", id)
		}(i)
	} 

	

    time.Sleep(4000 * time.Millisecond)
}
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

	// NOTE: select 中没有default时，所有case都阻塞情况下，select就会阻塞，
	// 直到某个case结束阻塞
	select {
	// 如果有写协程，阻塞
	case rwMutex.write <- struct{}{}:
	// 如果没有读协程，阻塞
	// NOTE: 如果写协程存在，第一个读协程就会阻塞在这里，
	// 当写协程释放锁后，第一个读协程就会从37行代码逃出
	// 阻塞状态，并在下边送入数据到readers channel，
	// 后续的读协程就不会再在这里阻塞了
	case rs = <- rwMutex.readers:
	}

	rs++
	rwMutex.readers <- rs
}

func (rwMutex RWMutex) RUnLock() {
	rs := <- rwMutex.readers
	rs--
	if rs == 0 {
		// NOTE: 当读协程数量归0时，要释放写锁，
		// 同时不能往 readers 中送入数据了，否则
		// 会出现一个读协程和一个写协程能同时运行的状况
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
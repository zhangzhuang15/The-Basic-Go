package main

func main() {
	var ch chan int
    defer func(){
    	println("error!")
	}()
    ch <- 5
}

// 写入 nil channel 会阻塞

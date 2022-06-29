package main

func main() {
	ch := make(chan int, 1)

	defer func(){
		println("error!")
	}()
	close(ch)

	close(ch)
}

// close 一个已经关闭的channel 会panic
package main


func main() {
	var ch chan int

	defer func(){
		println("error!")
	}()

	close(ch)
}

// 关闭 nil channel 会panic
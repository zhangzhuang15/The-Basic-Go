package main


func main() {
	var ch chan int
    defer func(){
    	println("error!")
	}()
	if value, ok := <- ch; ok {
		println("value: ", value)
	} else {
		println("error")
	}
}

// 读取 nil channel 会阻塞

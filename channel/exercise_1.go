package main


func main() {
	ch := make(chan int)

	go func(){
		ch <- 3
		// time.Sleep(500 * time.Millisecond)
		// close(ch)
	}()

	for value := range ch {
		println("value: ", value)
	}
}

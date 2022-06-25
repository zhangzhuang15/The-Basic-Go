package main


func main() {
	ch := make(chan int)

	go func(){
		ch <- 3
	}()

	for value := range ch {
		println("value: ", value)
	}
}

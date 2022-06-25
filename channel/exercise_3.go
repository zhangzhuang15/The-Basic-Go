package main

func main() {
	var ch chan int
    defer func(){
    	println("error!")
	}()
    ch <- 5
}

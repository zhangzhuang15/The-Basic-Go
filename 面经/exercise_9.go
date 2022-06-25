package main

func main() {
	type Box [2]int

	box1 := Box{1, 2}

	box2 := Box{1, 2}

	println(box1 == box2)
	println(&box1 == &box2)
}

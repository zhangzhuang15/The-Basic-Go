package main

import "fmt"

// struct{} 不占用内存，
// struct{} 和 map 结合， 可以创造 hashSet

type HashSet map[string]struct{}

func (set *HashSet)Has(s string) bool {
	_, ok := (*set)[s]
	return ok
}

func (set *HashSet)Add(s string) {
	if _, ok := (*set)[s]; !ok {
		(*set)[s] = struct{}{}
	}
}

func (set *HashSet)Delete(s string) {
	if _, ok := (*set)[s]; ok {
		delete(*set, s)
	}
}

func main() {
	var hashset = HashSet{}

	hashset.Add("Tom")

	fmt.Println("has Tom ?  ", hashset.Has("Tom"))

	hashset.Delete("Tom")

	fmt.Println("has Tom ?  ", hashset.Has("Tom"))
}

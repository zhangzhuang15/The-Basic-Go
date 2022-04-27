package main

import "fmt"

type set interface{
	set1(s string)
	set2(s string)
}

type test struct {
	s string
}

func (t *test) set1(s string) {
	t.s = s
}

func (t test) set2(s string) {
	t.s = s
}

func main() {
	var (
		t1 test
		t2 = new(test)
	)

	t1.set1("1")   // 等效于 (&t1).set1("1")
	fmt.Print(t1.s)
	t1.set2("2")   // t1 按照值传递到 set2方法中，因此 t1.s不会被修改
	fmt.Print(t1.s)

	t2.set1("3")
	fmt.Print(t2.s)
	t2.set2("4")   // 等效于(*t2).set2("4"), 同样是值传递，不会修改 t2.s
	fmt.Print(t2.s)

	fmt.Print("	\t")

	_, ok1 := (interface{}(t1)).(set) // t1 是 test 类型，只拥有 set2 方法
	_, ok2 := (interface{}(t2)).(set) // t2 是 *test 类型，拥有 set1 set2 方法

	fmt.Println(ok1, ok2)
}

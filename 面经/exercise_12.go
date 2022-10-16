package main

import "fmt"

type M string

func (m M) Speak() {
	fmt.Println("hello ", m)
}

func (m *M) Say() {
	fmt.Println("I am OK, ", *m)
}

func main() {
	const a M = "Jack"
	var b M
	var c = new(M)

	b.Speak()
	b.Say()
	c.Speak()
	c.Say()
	a.Speak()
	// Error! a 不可寻址，因为它是const类型
	// a.Say()
}

// M类型的数据，可以使用(m M)定义的方法，
// 如果M类型的数据是可以寻址的，也可以使用(m *M)定义的方法

// *M类型的数据，可以使用(m M)定义和(m *M)定义的方法，
// 在编译的时候，会自动识别，加上解指针操作

// (m M) 和 （m *M） 定义的方法，方法名不可以重复

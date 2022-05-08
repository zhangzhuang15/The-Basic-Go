package main

import (
	"fmt"
	"log"
)

func work() {
	defer func() {
		c := recover()
		fmt.Println(c)
	}()
	panic("fafa")
}

func main() {
	defer func() {
		log.Println("outgoing")
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	//log.Fatalln("this is a test error")
	//panic("right")
	//work()
	//fmt.Println("to the end")
	//os.Exit(200)
}
// 执行代码，发现defer执行了
// 打开29行注释，增加 import "os" , 执行代码，发现 defer没有执行
// 打开25行注释，执行代码，发现 defer没有执行，查询API可知，log.Fatalln输出日志后，会执行os.Exit结束进程
// 关闭25行注释，打开26行注释，执行代码，发现defer执行了，可是第28行代码并没有执行
// 关闭26行注释，打开27行注释，执行代码，发现defer执行了，第28行代码也执行了
// 查看 panic 文档说明，可以知道：
//
//  func A() {
//      defer func D()
//      func B() {
//             defer func C()
//
//             panic()
//      }
//  }
//
//  B()中遇到 panic，跳到B()的defer C 执行，执行完毕后，回到 func A(), 再次触发 panic，func D执行，
//  之后再返回到 A() 的上级函数中。。。。
//  如果在 C 中使用 recover 了，那么 A 中不会再次触发 panic, A中继续执行 B 后边的代码
//
//  总结defer：
//      函数正常结束时，defer会执行；
//      进程结束时， defer不会执行；
//      函数中出现panic时，defer会执行。
//                      1）如果defer中不使用 recover， panic会向上层函数传播；
//                      2）如果defer中使用 recover，回到上层函数，继续执行后续代码；
//

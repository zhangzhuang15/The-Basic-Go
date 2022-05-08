package main

import "runtime"

// 在函数执行中，获取函数的信息

func work() {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		// 本函数所在文件的绝对路径
		println("file name: ", file)
		// 本函数调用 runtime.Caller 位于源码第几行
		println("line: ", line)

		f := runtime.FuncForPC(pc)
		functionName := f.Name()
		// 本函数名
		println("func name: ", functionName)
	}
}

func manage() {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		// 本函数的上层函数main所在文件的绝对路径
		println("file name: ", file)
		// 本函数在 main 函数被调用时位于源码第几行
		println("line: ", line)

		f := runtime.FuncForPC(pc)
		functionName := f.Name()
		// 本函数的上层函数函数名
		println("func name: ", functionName)
	}
}

func main() {
	work()
	manage()
}

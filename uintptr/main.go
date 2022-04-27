package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 0x0000 001C
	var a uint32 = 28

	// 转化为通用指针
	var a_ptr = unsafe.Pointer(&a)

	// uintptr 将通用指针转化为地址值，加1后，转换回通用指针
	var a_new_ptr = unsafe.Pointer(uintptr(a_ptr) + 1)

	// 通用指针转换为具体指针，该指针指向 001C 的 00
	var a_final_ptr = (*byte)(a_new_ptr)

	// 重新赋值，得到 0x0000 011C
	*a_final_ptr = 1

	// 11c
	fmt.Printf("%x", a)

}

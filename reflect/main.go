package main

import (
	"fmt"
	"reflect"
)

type Man struct {
	Age int  `age:"年纪" required:"yes"`
	Name string `name:"姓名"`
}

func main() {
	man := Man{ 14, "Jack" }

	// rfv 用于获取 值信息
	rfv := reflect.ValueOf(man)
	// rfk 用于获取 类型信息
	rfk := reflect.TypeOf(man)

	// 通过 Type 也可以拿到值的类型信息
	fmt.Println("rfv.Type: ", rfv.Type())
	fmt.Println("rfv.Kind: ", rfv.Kind())
	// Field获取struct中字段的值，Field(0)是第一个字段，Field(1)是第二个字段，
	// String() 可以将 Value 类型值转化为 string 类型值
	fmt.Println("rfv.Field(1): ", rfv.Field(1).String())

	fmt.Println("")

    // rfk只能拿到类型信息，类型值拿不到
	fmt.Println("rfk.Kind: ", rfk.Kind())
	fmt.Println("rfk.Name: ", rfk.Name())
	fmt.Println("rfk.Size: ", rfk.Size())
	fmt.Println("rfk.Field(0).Type: ", rfk.Field(0).Type)
	fmt.Println("rfk.Field(1).Offset: ", rfk.Field(1).Offset)
	fmt.Println("rfk.Field(0).Tag: ", rfk.Field(0).Tag)
	fmt.Println("rfk.Field(0).Tag.Get('age'): ", rfk.Field(0).Tag.Get("age"))
    fmt.Println("rfk.Field(0).Tag.Get('required'): ", rfk.Field(0).Tag.Get("required"))


	var m int = 0
	// Zero 得到某个 Type 数据的零值
	n := reflect.Zero(reflect.TypeOf(m)).Interface()
	fmt.Println("m == n: ", reflect.DeepEqual(m, n))

}
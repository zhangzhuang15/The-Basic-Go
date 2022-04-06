### 如何进行类型转换？
```go
type MyInt int 

var m int = 1

var n interface{} = 3

// 具体类型可以这样显示转换
var m_myint MyInt = (MyInt)m

// 接口类型可被认为是通用类型，可以这样转换
var n_myint MyInt = n.(MyInt)
```

<br>

### 下面的切片声明是否正确？
```go

m := []int { 1, 2, 3}      // ✅

m := []int { 1, 2, 3,}     // ✅

m := []int {              //  ✅
    1, 2, 3,
    4, 5,
} 

m := []int {              // ❌
    1, 2, 3,
    4, 5
}

m := []int {             // ✅
    1, 2, 3,
    4, 5}
```

<br>

### go中什么样的数据不能被`json.Marshal`序列化?
* func 
* channel
* complex
* 循环引用的数据
> 结构体中首字母小写的字段，不会被序列化

<br>

### go各个类型数据的默认值
* 引用类型数据 ->  `nil`
  1. slice
  2. map
  3. channel
  4. func
  5. (p *P) func

* 指针类型 ->  `nil`
* 值类型
  1. int  -> `0`
  2. string -> `""`
  3. float -> `0.0`
  4. bool -> `false`
  5. array -> `0`

> interface{} 默认值也是 `nil`

### main函数有什么要求？
* 无输入参数
* 无返回值
* 必须位于package main中
* package main包中只能有一个main函数  
* main函数不能被显示调用
> 1. init函数也要无输入参数，无返回值；  
> 2. init函数不能被显示调用；   
> 3. 一个package中，init函数可以有多个；  
> 4. 先初始化导入包的变量和常量，再执行init函数，之后初始化本包的变量和常量，再执行init函数，最后执行main函数。

<br>

### 使用channel的时候要注意哪些问题？
* 对一个关闭的通道再发送值就会导致panic
* 对一个关闭的通道进行接收会一直获取值直到通道为空
* 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值
* 关闭一个已经关闭的通道会导致panic
* 对一个nil channel 读写数据会发生阻塞
* 关闭nil channel 会发生panic

<br>

### 所有的数据类型都可以定义方法么？
不可以。只有自定义类型才可以。像 int 就不可以，但是可以声明 type MyInt int后，对MyInt定义方法。


<br>


### slice map 的使用
* map必须初始化后再使用；
* slice可以未初始化使用，在append函数中完成初始化；


<br>

### interface的理解
* 两个interface的方法列表一样，但是顺序不同，二者等效；
* interface A 中的方法列表是 interface B 方法列表的子集，则 A = B 赋值运算成立；
* 接口调用要在运行时才能确定；
* 接口赋值在编译期就可以确定；

<br>

### const 常量对数据类型有什么限制
*  bool
*  int 
*  float
*  complex
*  string
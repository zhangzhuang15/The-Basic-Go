先看效果  
在本目录下，执行 `go test test -v`
> 这将执行 echo_hello_test.go  echo_world_test.go  
> 第一个test是go的子命令  
> 第二个test是mod名，详见go.mod

执行 `go test test/utils -v`
> 这将执行 utils/echo_world_test.go
---
具体解释  
这个文件夹下的文件是怎么来的？
* 在`测试`文件夹下执行 `go mod init test`；  
* 编写了你现在看到的.go文件；
> mod 的名字是 test；  
> utils文件夹作为一个package，它里面的.go文件都要声明 package test；  
> 作为单元测试的文件命名规则是 *_test.go， 文件中的单元测试函数必须以Test开头；   
> 测试文件和被测试文件同属于一个package下；

go test [package_name] 就可以执行某个package下所有的测试文件；  


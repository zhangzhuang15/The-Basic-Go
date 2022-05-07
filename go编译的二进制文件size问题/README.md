编译一段最简单的示例代码hello world，C语言编译出来的二进制文件要远小于go语言编译的结果。  

为什么会这样呢？  
* c语言在编译的时候，采用动态链接的编译方式，因此会比较小，特别是常用的头文件，它的实现位于libc的动态库中，不会被拷贝到编译出来的二进制文件当中，使得二进制文件体积更小，一般是48KB(macOS上)/16KB(Linux上）。
    > 使用静态链接的编译方式，得到的二进制文件体积为889KB左右；  
    > 再用 strip 处理后体积来到 691KB 左右； 
    > 注意这里并没有开启 gcc 的 `-O`参数；
* go语言在编译的时候，采用静态链接的编译方式，虽然体积很大，但是并不会依赖具体的动态库，可移植性非常好，同时它自身的package之间存在嵌套引入，即便你引入了一个package，可实际上这个package实现的时候需要引入其他的package，造成编译出来的二进制文件体积很大。当go采用内置的`println`函数，并且使用`go build -ldflags="-w -s" -o main main.go`，得到的二进制文件会加以缩小，达到大约780KB左右。
    > 和静态链接编译下的C相比，go还是非常紧凑的。
* 如果使用Rust编译的话，二进制文件体积会更加乐观
    * `cargo build` 得到 447KB;
    * `cargo build --release` 得到 244KB;
    > Rust和Go一样，采用静态链接的编译方式，将自身的一些标准库编译进去。但是从结果来看，Rust比Go好，说明Rust对C的封装抽象更少，而且没有垃圾回收运行时，也省去了一部分代码。

**总结**  
站在一个偏OS的应用层语言，Go已经非常出众了，不比C逊色；   
站在一个偏OS的系统层语言，Rust和C更出众，因为他们足够小，更重要的是没有垃圾回收。

---

<br>

## 尝试
安装好 go,gcc,rust后，可以尝试：
1. `go build -o maingo main.go`  
2. `go build -o maingo_small -ldflags="-w -s" main.go`  
3. `gcc -o mainc main.c`  
4. `gcc -static -o mainc_static main.c`   
5.  * `cd hello-world && cargo build` 
    * `cd hello-world && cargo build --release`

比对一下 二进制文件的大小吧
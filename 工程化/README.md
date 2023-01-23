## GOPATH
存储下载的第三方库 
module模式下载的第三方库存储于`$GOPATH/pkg/mod`

## GOROOT
存储GO源码，标准库源码，go二进制指令  
- go指令存储于`$GOROOT/bin`
- 标准库.a静态文件存储于`$GOROOT/pkg`  
- 标准库源码存储于 `GOROOT/src`

## 代码组织方式
一个项目、一个工具库基本上对应一个 module；
一个module下面由若干个 package 组成；

可以和其他语言类比一下：
go语言的module，javascript中叫做package，rust中叫做 crate；
go语言的package，javascript中叫做模块，rust中叫做 mod；

在javascript的package管理中，有一个叫做monorepo的概念，就是使用一个github仓库，管理若干个package代码，而不是一个package代码托管在一个独立的github仓库中；

go语言也可以做到这点，用一个github仓库，管理多个module，这就是workspace；

### 如何使用module
[前往](./module/README.md)

### 如何使用workspace
本项目就使用了workspace, 在项目根目录下存在go.mod和go.work.

go.mod
```shell
module theBasicGo

go 1.18
```

go.work
```shell
go 1.18

use (
 ./go_install
)
```
* go.mod 记录了根module；
* go.work 使用 `use` 指出子module分布在哪里；
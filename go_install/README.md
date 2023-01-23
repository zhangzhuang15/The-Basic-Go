无论是本地module，还是远程module，
执行 `go install`后，go会下载源码并编译成二进制程序，安装在用户本地，供用户使用；

本文件夹就是一个本地module，
module名 `"github.com/zhangzhuang/hello"`,
在安装之后，就可以得到同名二进程程序 `hello`,
在shell中执行：
```bash
$ hello 
hello world
```

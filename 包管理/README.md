### 本目录中的文件是怎么来的？
1. 创建文件夹`包管理`
2. 在终端中，进入该文件夹
3. 执行 `go mod init jack`, 生成了 `go.mod` 文件
4. 其余文件，都是手动创建添加的
   
<br>

---

### package名 和 文件夹名 的关系

* 与`go.mod`同目录下的`.go文件`, 都属于 `package main`, 入口文件是 `main.go`, 你的代码的入口是`main.go` 中的 `main`函数；
  
* 同一个文件夹下的`.go`文件，必须属于同一个`package`, `package名` 和 文件夹名可以不同；
  > 为了使得项目结构更加清晰，package名和文件夹名保持一致

* 如果一个文件夹下存在`.go`文件，那么这个文件夹肯定对应<span style="color: red">**唯一**</span>的`package`;

* 使用`import` 引用 `package`时，路径表示的是`文件夹`，不是`package`.以本目录下的main.go为例解释一下。
  ```go
  // main.go

  import （
        "fmt"
        u "jack/utils"
  )
  // jack 是 mod 名，也就是整个项目package管理系统的根目录；
  // utils 是指 utils 文件夹，不是指 package utils;
  // 想要验证这一点并不难：
  // 1. 查看utils文件夹，将 utils 文件夹下的文件都改为 package myutil;
  // 2. 在 main.go 中，将 'jack/utils' 改为 'jack/myutil'，发现找不到package；
  
  // package alias 别名
  // 当文件夹名和文件夹下的package名不一样的时候，
  // 无法直接使用package名，还是以 main.go 和 utils 文件夹为例：
  // 1. 将 utils 下所有文件的package归属改为 package myutil;
  // 2. 在 main.go 中将 u "jack/utils" 改为 “jack/utils”；
  // 结果就是没有办法在 main.go 中使用 myutil.Add使用 utils/add.go定义的函数；
  //
  // 解决办法就是使用别名：
  // 接着上边的操作结果，将 main.go 改回为 u "jack/utils";
  // 1. u 就是 myutil 的别名；
  // 2. 如果 文件夹名和package名一样，但是名字太长时，也可以用别名处理，简化package的引用；
  ```
  ```go
  // main.go
  import (
      "fmt"
      "jack/utils"
  )
  // 很抱歉，无法使用 utils 访问 myutil 下公开的函数、变量；
  // 这是因为 文件夹名 和 package 名不一致造成的；
  // 要么像上边那样起一个别名 u；
  // 要么修改 package myutil 为 package utils ;
  ```


### 使用本地module替代远程module  
  ```bash
  一般开发中，我们会下载远程的module依赖，之后在自己的代码中import他们。

  那么问题来了，能不能import一个本地的module呢？
  ```
  请看go.mod文件。
  ```
  module jack

  go 1.17


  require github.com/peter v0.0.1
  replace github.com/peter => "./peter"
  ```
  > require那一行指定了我们想拉取一个远程的module使用；  
  > replace那一行指定了我们使用./peter中的module替换掉远程的module；  
  > 再看看peter文件夹中有什么？go.mod ！没错，peter文件夹中正是一个module；

  再看看 main.go 文件是怎么使用的？
  ```go
  import (
	  "fmt"
	  u "jack/utils"

	  p "github.com/peter"
  )

  func main() {
      ...

	  p.Hello()
  }
  ```
  或者
  ```go
  import (
	  "fmt"
	  u "jack/utils"

	  "github.com/peter"
  )

  func main() {
      ...

	  peter.Hello()
  }
  ```

  就是这么简单!  
  在./peter/peter.go文件中，你会看到
  * package peter
  * func Hello()
  
  *好消息是 go 1.18给出了workspace的方法，也能达到这种效果啦*

### 发布你自己的module  
  ```bash
  没错，你可以将自己的module发布到go.dev上边去！

  具体操作：
   1、 在本地建立一个module，还记得 go mod init 指令嘛？module先随便起一个名字，后边会修正的。

   2、 由于是开发一个工具包，因此 module 下的.go文件不属于 package main，
       不妨假设它们 package yyp；
       一定要记住这点假设，下边会用到！！！！

   3、 将你的module推到github上，并打一个tag。
       注意： 
       如果你的github仓库地址是 https://github.com/xxxx/yyy.git 
       那么go.mod中的 module名字应该是 github.com/xxxx/yyy 或者 
       在yyy后边加一些额外的路径，比方说 github.com/xxxx/yyy/e3、github.com/xxxx/yyy/v4之类的, e3可能表示子文件夹名e3，v4则可能是指分支名。

       你打的tag，对应的就是 go.mod 文件中 require 语句的版本号，
       比如说 require github.com/xxxx/yyy v0.0.1  中的 v0.0.1。

   4、 推到GitHub上之后，你需要随便建立一个文件夹，比如叫做temp，在temp中执行
       go mod init temp 初始化为一个module， 
       之后执行 go get github.com/xxxx/yyy 
       (github.com/xxxx/yyy假设就是你的包名)，
       就会激活go.dev下载你的包信息

   5、 登陆到go.dev 就能看到你的包啦。
   （可能无法立即看到，等一小段时间就能看到）

   6、 使用你发布的包，只需要
       import "github.com/xxxx/yyy"

       引用包时，不是使用 yyy, yyy只是 module 名的一部分；

       更具体地说，yyy表示的是 module 被存储的文件夹，
       想想看，如果你git clone github.com/xxxx/yyy.git会发生什么？
       是的，会在本地创建一个 yyy 文件夹，项目放在这个文件夹下；

       而该module下真正的 package 名是 yyp.
       所以要使用 yyp 引用包中的函数、数据。
       这个例子中，package名是 yyp，比较短，
       如果package名很长，你也可以用 package 别名去处理；
  ```
  如果你发布的module版本号大于等于2.0，需要做一些调整：
  * module名要改为 `github.com/xxxx/yyy/v2`
    > /v2这种写法是官方强制要求的；
  * 引用你的module时要改为 `import "github.com/xxxx/yyy/v2"`;
  
  实际案例可参考：https://github.com/panjf2000/ants
  
  <br>

  如果你采取同一个仓库，管理多个module，比如在仓库 `github.com/xxxx/yyy.git` 的 `amp` 文件夹下，你有一个module，这个module对应的package名是 msq, 你就需要做如下调整了：
  * `amp`文件夹的module名要写成 `github.com/xxxx/yyy/amp`;
    > 如果版本号大于等于2.0，还要把类似 /v2 的标记在末尾补上；
  * 引用的时候，也要用 `import github.com/xxxx/yyy/amp`

  // TODO: 补充实际案例和Go官网介绍

<br>

---

### 外界对package的访问权
* package中的函数、变量，如果首字母大写，即对外公开，外部可以通过import后访问；
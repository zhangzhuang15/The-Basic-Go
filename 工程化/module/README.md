### 本目录中的文件是怎么来的？
1. `mkdir` 创建文件夹 `module`
2. 在`zsh`中，进入该文件夹
3. 执行 `go mod init jack`, 生成了 `go.mod` 文件
4. 其余文件，都是手动创建添加的
   
<br>

---

### package 和 文件夹 的关系
* 一个文件夹对应一个package;
* 一个文件夹下的所有go文件必须属于同一个package;
* 理论上，package的名字可以和文件夹名不一样，但在实践上，必须保持一致；
* 如果项目是二进制应用，那么项目根目录(go.mod所在文件夹)对应 main package;
* 如果项目是库应用，那么项目根目录也对应一个package，这个package名字不能叫做main；
* module名是项目根目录的别名；
* 使用`import` 引用 `package`时，路径表示的是`文件夹`，不是`package`.
  ```go
  // main.go

  import （
        "fmt"
        u "jack/utils"
  )
  // jack 是 module 名，也就是整个项目的根目录；
  // utils 是指 utils 文件夹，不是指 package utils;
  // 验证这一点：
  // 1. 查看utils文件夹，将 utils 文件夹下的文件都改为 package myutil;
  // 2. 在 main.go 中，将 'jack/utils' 改为 'jack/myutil'，发现找不到package；
  
  // 当文件夹名和package名不一样的时候，
  // 无法直接使用package名调用其中公开的数据、函数，
  // 验证这一点：
  // 1. 将 utils 下所有文件的package 改为 package myutil;
  // 2. 在 main.go 中将 u "jack/utils" 改为 “jack/utils”；
  // 3. 在 main.go 中使用 myutil.Add；
  // 编译和运行都会报错
  //
  // 使用 package alias 可以解决此问题，
  // u "jack/utils" 就是 package alias,
  // u.Add 可以成功通过编译和运行；
  // 如果 文件夹名和package名一样，但是名字太长时，
  // 也可以用别名处理，简化package的引用；
  ```

### 用本地module替代远程module  
一般开发中，我们会下载远程的module依赖，之后在自己的代码中import他们。如果你开发了一个module，想知道别人下载后，import该module的使用效果，该怎么办？真的要先上传么？

不需要，可以用本地module替换远程module，做一个mock；

go.mod
```shell
module jack

go 1.18

require github.com/peter v0.0.1
replace github.com/peter => "./peter"
```

* require 指定了我们想拉取一个远程的module使用；  
* replace 指定了我们使用./peter中的module替换掉远程的module；  
* peter文件夹中有什么？go.mod ！没错，peter文件夹中正是一个module；

main.go
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

  // module github.com/peter 是 peter 文件夹的别名，
  // 该文件夹下的 package 名是 peter, 因此使用 peter.Hello,
  // 实际上go fmt 会自动帮助我们为 "github.com/peter"加上
  // package alias, 所以不用担心 package 名是否和peter保持一致
  func main() {
      ...

	  peter.Hello()
  }
```

### 发布module  
```bash
具体操作：
   1、创建一个github仓库，下载到本地，假设仓库地址为
   https://github.com/xxxx/yy.git

   2、进入根目录，执行 `go mod init github.com/xxxxx/yy`

   3、假设开发一个工具库，因此package名不能叫做main，不妨叫yyp；

   4、打上tag v1.0.4, 并推到远程仓库。
   你打的tag，对应 go.mod 中 require 语句的版本号，
   比如说 require github.com/xxxx/yyy v0.0.1 中的 v0.0.1。

   5、随便建立一个文件夹，比如叫做temp。
   在temp中，
   执行go mod init temp 初始化为一个module， 
   执行 go get github.com/xxxxx/yy,
   就会激活go.dev下载你的包信息

   6、登录到go.dev 就能看到你的包啦。
   （可能无法立即看到，等一小段时间就能看到）

   7、使用你发布的包，只需要
   import "github.com/xxxxx/yy"
   引用包时，不是使用 yy, 要使用 yyp；   
```
  <br>

  如果你发布的module版本号大于等于2.0，需要做一些调整：
  * tag要打成 v2.0.4 之类的，大版本号必须是2
  * go.mod 的 module 改为 `github.com/xxxxx/yy/v2`
    > /v2这种写法是官方强制要求的；
  * 引用你的module时要改为 `import "github.com/xxxxx/yy/v2"`;
  
  实际案例可参考：https://github.com/panjf2000/ants
  
  <br>

  如果你采取同一个仓库，管理多个module.
  仓库依旧是 `https://github.com/xxxxx/yy.git`,
  根目录的go.mod 
  ```shell
  module github.com/xxxxx/yy

  go 1.18.1
  ```
  go.work
  ```shell
  go 1.18.1

  use (
    ./amp
  )
  ```
  根目录下有个 amp 目录，里面的 go.mod:
  ```shell
  module github.com/xxxxx/yy/amp
  
  go 1.18.1
  ```
  发版打上tag amp/v1.0.4
  使用的时候，只需 `import "github.com/xxxxx/yy/amp"`;

  如果module是v2版本，
  发版打上tag amp/v2.0.3,
  使用时则需要 `import "github.com/xxxxx/yy/amp/v2"`;

  在使用该module的项目go.mod中，如果指定:
  ```require github.com/xxxxx/yy/amp/v2.0.3```
  那么go就会选择 `github.com/xxxxx/yy.git` tag号amp/v2.0.3的代码，将其中的 amp 文件夹下载到本地，作为依赖使用；
  [更多信息参考](http://docscn.studygolang.com/doc/modules/managing-source)

<br>

---

### 外界对package的访问权
* package中的函数、变量，如果首字母大写，即对外公开，外部可以通过import后访问；
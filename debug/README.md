## Go 代码如何调试

#### 使用 pycharm或者vscode扩展调试
* 直观
* 简单
* 无法跟踪到Go内部代码实现
  
1. 安装 go 编译器和 go vscode官方插件；
   > 安装 go vscode官方插件时，会被自动导向下载相关的一些工具，其中就包括dlv调试工具，但是该工具只在vscode调试的时候发挥作用，如果你想使用该工具，请将其加入到环境变量path中，或者使用 brew 单独再安装一次。
2. 安装 dlv工具；
3. 打开go文件，打上断点；
4. 点击vscode侧边栏的调试面板，选择debug即可；
   
---

#### dlv工具
一个专门调试 Go 程序的工具。

安装：
`brew install delve`

使用：
1. 进入go的项目包（本例是进入demo文件夹）
2. `dlv debug`
   > 项目包路径中不能包含中文😱
3. 进入调试界面，`breakpoint main.go:7`
   > 在 main.go 文件的第7行打上断点。
4. 更多的调试指令，用`help` 查看。
   > 相信我，提示非常友好🤩


#### 只生成汇编代码
`go tool compile -S -N -l main.go`
> 汇编代码会直接输出到终端


#### 反汇编
`go tool objdump ./main`
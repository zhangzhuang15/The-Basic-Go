package main

import (
	"fmt"
	"runtime"
)

func main() {
    var system string = runtime.GOOS
	fmt.Printf("hello! This is %s", system)
}


// 🌟 go version > 1.5
// 如果你要编译的是:
//
// x86_64\amd64\x64 windows 的程序， 执行 GOOS=windows GOARCH=amd64 go build -o main_windows_amd64 main.go
//
// GOOS是操作系统，GOARCH是CPU架构
// 相关组合有
//
// GOOS       GOARCH            补充
// -----------------------------------------------
// darwin     arm64       对应apple silicon macOS
// darwin     amd64       对应intel x86_64  macOS
// darwin     386         对应intel i386    macOS
// +++        +++         ++++++++++++++++++++++
// linux      arm64
// linux      amd64
// linux      386
// +++        +++         +++++++++++++++++++++++
// windows    amd64
// windows    386         i386是 32位体系哦
// +++        +++         +++++++++++++++++++++++
// freebsd    amd64
// freebsd    386
// +++        +++         +++++++++++++++++++++++

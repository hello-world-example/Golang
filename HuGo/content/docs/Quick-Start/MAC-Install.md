# MAC 下安装



## 安装

- https://golang.org/ 官网可能会被屏蔽
- 可以从 https://studygolang.com/dl 下载安装包
- 这里下载 **go1.15.2.darwin-amd64.pkg**  安装包
- 双击安装即可



``` bash
$ which go
/usr/local/go/bin/go

$ cat /etc/paths.d/go
/usr/local/go/bin

# 版本
$ go version
go version go1.15.2 darwin/amd64

# 查看帮助
$ go help

```



## 环境配置

```bash
# GOPATH
$ go env GOPATH
/Users/kail/go

# Go version >= 1.13，无需配置环境变量，通过以下命令即可配置
$ go env -w GOPATH=/Users/kail/go2

# 国内代理配置
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```



## Hello World

> **注意**： Go 的入口必须是 **main 包 main 方法**，**包名和文件夹名可以不一样**
>
> `go run HelloWorld.go` 运行程序

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```



## 卸载

```bash
# 删除 go 的安装目录
$ rm -rf /usr/local/go

# MAC 下才有该文件， 文件内容是 go 命令所在的路径
$ rm -f /etc/paths.d/go
```



## Read More

- [Go 安装包下载](https://studygolang.com/dl)
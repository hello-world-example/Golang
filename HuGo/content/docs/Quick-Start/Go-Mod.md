# Go Mod



> [拜拜了，GOPATH君！新版本Golang的包管理入门教程](https://zhuanlan.zhihu.com/p/60703832)



## 准备工作

```bash
# GO111MODULE 为 on 或者 auto
$ go env -w GO111MODULE=on

# 国内代理配置
$ go env -w GOPROXY=https://goproxy.cn,direct
```



## Hello World

> 文件名： HelloWorld.go

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, world!")
}
```



## 初始化项目

```bash
$ go mod init HelloWorld

# 当前目录下会多一个 go.mod 文件，内容如下
$ cat go.mod
module HelloWorld

go 1.15

```



## 添加依赖

> - 无需使用 `go get` 命令下载依赖
> - 直接 `go run HelloWorld.go` 即可
> - 运行之后 `cat go.mod`，多了一行 `require github.com/astaxie/beego v1.12.2`

```go
package main

import (
  "fmt"
  "github.com/astaxie/beego"
)

func main() {
  fmt.Println("Hello, world!")
  beego.Run()
}
```



## 自定义模块

新建 `module/tools/HelloUtil.go` 文件，内容如下

```go
package tools

import "fmt"

func PrintHello() {
	fmt.Println("Hello ~~~")
}
```

在 `HelloWorld.go` 中使用本地模块

```go
package main

// 【注意】这里不是 "./module/tools"
import (
	"HelloWorld/module/tools"
)

func main() {
	tools.PrintHello()
}
```



## FAQ

### 依赖的包的本地地址

- 使用 Go 的包管理方式，依赖的第三方包被下载到了 `$GOPATH/pkg/mod` 路径下
- 类似于 `github.com/astaxie/beego@v1.11.1`



### 依赖包的版本控制

- `$GOPATH/pkg/mod`里可以保存相同包的不同版本
- 版本是在 `go.mod` 中指定的，如果没有指定，go 命令会自动下载代码中的依赖的最新版本
- 指定版本时可以用 `latest`，这样它会自动下载指定包的最新版本



### GO111MODULE 变量

- go 会根据 `GO111MODULE` 的值而采取不同的处理方式
- 默认情况下，`GO111MODULE=auto` 自动模式
  - **auto** ： 项目在 `$GOPATH/src`里会使用 `$GOPATH/src` 的依赖包；在 `$GOPATH/src` 外，就使用 `go.mod` 里 `require` 的包
  - **on** ： 无论在 `$GOPATH/src` 里还是在外面，都会使用 `go.mod` 里 `require` 的包
  - **off** ：使用  `$GOPATH/src` 的依赖包



### 地址失效了怎么办

- 在 go.mod 文件里用 replace 替换包，如：
- `replace golang.org/x/text => github.com/golang/text latest`
- go 会下载 `http://github.com/golang/text` 的最新版本到 `$GOPATH/pkg/mod/golang.org/x/text` 下


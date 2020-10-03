# 依赖模块



## 导包

```go
import mrand "math/rand" // 【重命名】 在当前文件里有效

import . "fmt"           // 【静态导入】可以直接使用 Printf(...)

import _ "math/rand"     // 【匿名导入】不使用包里的函数，仅是为了调用包里的init初始化函数
```



## 自定义模块

>  [go语言中的模块（包）](https://juejin.im/post/6844903520366247944)



### 项目结构

```bash
/
|--test.go
|--module/
|--|--tools/
|--|--|--helloUtil.go
```



### helloUtil.go

```go
package tools;

import "fmt";

func init()  {
	fmt.Println("init...")
}

// 大写字母开头， public 变量
var Hello string = "Hello"

// 小写字母开头： private 变量
var world string = " World";

// public 方法
func PrintHello(){
	fmt.Println(Hello + world);
}

// private 方法
func sayHello(){
	fmt.Println(Hello + world);
}
```



### test.go

```go
package main

import (
	// 相对 文件集 路径
	"./module/tools"
	. "fmt"
)

func main() {

	// 跟文件名没有关系，直接使用文件中的变量
	Println(tools.Hello)

	// 跟文件名没有关系，直接调用文件中的方法
	tools.PrintHello()	
}
```



### 注意事项

- 同一个目录，包名必须一样
- `public` 大写字母开头，`private` 小写字母开头



## 安装三方库

```bash
# 安装第三方库
$ go get -u github.com/PuerkitoBio/goquery

# 安装之后，源码会在 GOPATH/src 下
$ ll src/github.com/PuerkitoBio/goquery

# 如果 go get -u 无法安装或比较慢，可【手动安装】
$ git clone https://github.com/PuerkitoBio/goquery src/github.com/PuerkitoBio/goquery
```



## 三方库示例

```go
package main

import (
	gQuery "github.com/PuerkitoBio/goquery"
	. "fmt"
)

func main() {

	document, _ := gQuery.NewDocument("https://godoc.org/github.com/PuerkitoBio/goquery")
	
	selections := document.Find("#pkg-overview")

	selections.Each(func(i int, selection *gQuery.Selection) {

		text := selection.Text()

		Println(text)

	})

}
```



## Read More

- [GoDoc](https://godoc.org/) (Search for Go Packages)
- [Golang标准库文档](https://studygolang.com/pkgdoc)
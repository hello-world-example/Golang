
# 快速开始




## Hello World
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
>  **注意**： Go 的入口必须是 **main 包 main 方法**，**包名和文件夹名可以不一样**。

## 包

``` go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

按照惯例，**使用导入路径的最后一个目录**。例如，导入`math/rand` 包，可以使用`rand.Intn(10)` 调用其中的方法。

## 函数

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
函数可以没有参数或接受多个参数。在这个例子中，`add` 接受两个 `int` 类型的参数(**参数类型在变量名之后**)。

### 函数参数
当两个或**多个连续的函数命名参数是同一类型**，则除了**最后一个类型之外**，其他都可以省略。如：
```go
func add(x, y int) int {
	return x + y
}
```

### 函数返回值
函数可以返回任意数量的返回值。
```go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

## 变量

```go
package main

import "fmt"

var c, python, java bool // 定义了一个变量的列表；跟函数的参数列表一样，类型在后面

var i, j int = 1, 2 // 变量初始化

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```
`var` 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。就像在这个例子中看到的一样，`var` 语句可以定义在包或函数级别。
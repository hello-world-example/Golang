# 基础语法



## 变量

```go
// 定义了一个变量的列表；跟函数的参数列表一样，类型在后面
var c, python, java bool 

// 变量 批量 初始化
var i, j int = 1, 2 

// 直接初始化
a := 3
// 初始化后赋值，这个不允许，a := 4
a = 4
```



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



### 引用传递

```

```










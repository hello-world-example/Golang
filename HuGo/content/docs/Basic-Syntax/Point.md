# 指针与引用

## 概述

- Go 语言的参数只能是值传递
- 值传递是一个拷贝的过程
- `&` 取址符，返回变量的内存地址
- `*` 取值符，返回指针指向的变量的值



## 值传递

```go
func append(L list.List) {
	L.PushBack(1)
	L.PushBack(2)
}

func main() {
	L := list.List{}
  // 值传递
	append(L)
  // 这里输出0
  fmt.Println(L.Len())
}
```



## 指针传递

```go
// L 内存地址对应的值
func append(L *list.List) {
	L.PushBack(1)
	L.PushBack(2)
}

func main() {
	L := list.List{}
  // L 内存地址
	append(&L)
  // 这里输出 2
  fmt.Println(L.Len())
}
```





## Read More

- [Golang - 指针与引用](https://www.cnblogs.com/anthony-dong/p/12249394.html)
package main;

import "fmt";

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
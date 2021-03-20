package main

import "fmt"

//[primary] 定义一个包内全局字符串变量，下面语法正确的是 （）
//A. var str string
//B. str := ""       [函数内局部变量]
//C. str = ""
//D. var str = ""    [赋值操作]
//参考答案 AD

// 定义一个包内全局字符串变量,要使用var
var str string

//var str = ""
//str := ""
//str = ""
func main() {
	fmt.Println(str)
}

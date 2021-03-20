package main

//[intermediate] 关于init函数，下面说法正确的是（）
//A. 一个包中，可以包含多个init函数
//B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数
//     [如引用了第三方包，mysql，这些包运行前要执行初始化操作，三方包的初始化比自己项目的高，求到别人了]
//C. main包中，不能有init函数
//D. init函数可以被其他函数调用
//
//参考答案：AB

//C 这个可以有
//D init函数是自动被调用的

func main() {

}

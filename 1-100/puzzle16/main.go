package main

//[intermediate] golang中没有隐藏的this指针，这句话的含义是（）
//A. 方法施加的对象显式传递，没有被隐藏起来 [都要用struct.method或*struct.method显示调用]
//B. golang沿袭了传统面向对象编程中的诸多概念，比如继承、虚函数和构造函数
//C. golang的面向对象表达更直观，对于面向过程只是换了一种语法形式来表达
//D. 方法施加的对象不需要非得是指针，也不用非得叫this
//参考答案：ACD

//B 大神都说了，go没有继承，只有组合

func main() {

}

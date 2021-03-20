package main

//[primary] 关于类型转化，下面语法正确的是（）
//A.
//type MyInt int
//var i int = 1
//var j MyInt = i  自定义MyInt和int是不同类型

//B.
//
//type MyInt int
//var i int = 1
//var j MyInt = (MyInt)i   Type MyInt is not an expression

//C.
//
//type MyInt int
//var i int = 1
//var j MyInt = MyInt(i)

//D.
//
//type MyInt int
//var i int = 1
//var j MyInt = i.(MyInt)  i不是interface接口类型,不能断言

//参考答案：C

//A
//type MyInt int
//var i int = 1
//var j MyInt = i

//B.
//type MyInt int
//var i int = 1
//var j MyInt = (MyInt)i

//C
//type MyInt int
//var i int = 1
//var j MyInt = MyInt(i)

//D.
//type MyInt int
//var i int = 1
//var j MyInt = i.(MyInt)

func main() {

}

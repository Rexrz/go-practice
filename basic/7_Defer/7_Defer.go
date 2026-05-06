package main

import "fmt"

/*
学习defer的调用顺序
栈的调用顺序 LIFO
*/
func func1() {
	fmt.Println("func1执行")
}

func func2() {
	fmt.Println("func2执行")
}

func func3() {
	fmt.Println("func3执行")
}

/*
	学习defer和return之间的调用顺序
*/

func deferAndReturn() string {
	defer func1()
	return "deferAndReturn中return方法执行"
}

// 视频中的方法
/*
当 return 后面跟着一个表达式（比如函数调用、数学运算等）时，Go 语言会优先把这个表达式计算/执行完毕，得出确切的结果并赋值后，才会去触发 defer 栈里的函数。

简单总结：如果是 return 一个函数调用，这个被调用的函数一定会在 defer 之前执行。
*/
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

func main() {

	//defer func1()
	//defer func2()
	//defer func3()

	//result := deferAndReturn()
	//fmt.Println(result)

	returnAndDefer()
}

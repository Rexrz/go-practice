package main

import "fmt"

/*
	学习go语言中的函数
	1.多返回值
		1.1 函数返回多个返回值，返回值为匿名
		1.2 函数返回多个返回值，返回值有形参
*/

// 函数一：函数返回单个返回值
func plus(a int, b int) int {
	return a + b
}

// 函数二：函数返回多个返回值
func minus(a int, b int) (minuend int, subtraction int, result int) {
	return a, b, a - b
}

func main() {
	a, b := 1, 1
	//plusResult := plus(a, b)
	//fmt.Println(plusResult)

	minuend, subtraction, result := minus(a, b)
	fmt.Println("被减数 =", minuend, "减数 =", subtraction, "差 =", result)
}

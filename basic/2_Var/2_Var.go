package main

import "fmt"

// 学习go语言的变量声明

/*
 一、局部变量的声明
	1. 方法一，声明一个变量，默认的值是0
	2. 方法二，声明一个变量，初始化一个值
	3. 方法三，在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	4. 方法四（常用），省去var关键字，直接自动匹配

 二、全局变量的声明
	上述局部变量的前三种都可以进行全局变量的声明，第四种不可以

 三、多变量的声明
	1. 单行写法
	2. 多行写法
*/

// 声明全局变量
var globalA int = 42
var globalB string = "24"

func main() {

	// 局部变量的声明
	// 方法一
	var a int
	fmt.Println("a =", a)
	// 打印一个变量的类型
	fmt.Printf("type of a = %T\n", a)

	// 字符串是空串
	var aStr string
	fmt.Println("aStr =", aStr)
	fmt.Printf("type of aStr = %T\n", aStr)

	// 方法二
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 方法三
	var c = true
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 方法四
	d := "hello"
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	// 打印全局变量
	fmt.Println("globalA = ", globalA)
	fmt.Printf("type of globalA = %T\n", globalA)

	fmt.Println("globalB = ", globalB)
	fmt.Printf("type of globalB = %T\n", globalB)

	// 多变量的声明
	// 单行直接声明多变量
	var j, k int = 1, 2
	fmt.Println("j =", j, ",k =", k)

	// 多行声明变量
	var (
		m = "hello"
		n = true
	)
	fmt.Println(m, n)

}

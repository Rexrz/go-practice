package main

import "fmt"

/*
	学习go语言中的指针
*/

func main() {
	var t1, t2 = 10, 20
	// 定义一个交换方法
	swap(t1, t2)
	fmt.Println("使用值交换的swap方法：")
	fmt.Println("t1 =", t1, "t2 =", t2)

	swapUserPointer(&t1, &t2)
	fmt.Println("使用指针交换的swap方法：")
	fmt.Println("t1 =", t1, "t2 =", t2)
}

/*
*
传递指针，则可以交换两个变量的值
*/
func swapUserPointer(t1 *int, t2 *int) {
	var temp int
	temp = *t1
	*t1 = *t2
	*t2 = temp
}

/*
*
传递值，不能交换两个变量的值
*/
func swap(t1 int, t2 int) {
	var temp int
	temp = t1
	t1 = t2
	t2 = temp
}

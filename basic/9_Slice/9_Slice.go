package main

import "fmt"

/*
	学习go语言中的切片（动态数组）
*/

func printSlice(slice []int) {
	for index, value := range slice {
		//fmt.Println(value)
		fmt.Println("index:", index, "value:", value)
	}
}

func main() {
	// 切片在方法形参的传递过程中是不校验长度的
	//var slice []int
	slice := []int{1, 2, 3}
	printSlice(slice)

	fmt.Println()
	// 切片传递的是“引用”及地址值
	updateSlice(slice)
	printSlice(slice)

	// 切片的声明方式有四种
	// 方式一：声明slice是一个切片，并且初始化，默认值是1，2，3；长度是len
	slice1 := []int{1, 2, 3}
	printSlice(slice1)

	// 方式二：声明slice是一个切片，但是没有给slice分配空间
	fmt.Println()
	var slice2 []int
	isNil(slice2)

	// 方式三：使用make关键字进行切片的初始化操作，开辟3个空间，每个元素的默认值都为0
	fmt.Println()
	slice2 = make([]int, 3)
	isNil(slice2)

	// 方式四：通过 := 的方式进行推导
	fmt.Println()
	slice3 := make([]int, 3)
	printSlice(slice3)
}

func isNil(slice2 []int) {
	if slice2 == nil {
		fmt.Println("该切片为空")
	} else {
		fmt.Println("该切片有分配地址")
	}
}

func updateSlice(slice []int) {
	slice[1] = 100
}

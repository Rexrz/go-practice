package main

import "fmt"

/*
	学习go语言中的数组的使用
*/

func main() {
	// 声明一个数组
	// 不初始化的时候，该数组所有元素都为该数组的数据结构类型的默认值
	myArray := [10]int{}

	myArray2 := [10]int{1, 2, 3, 4}

	// 遍历方式一
	for i := 0; i < 10; i++ {
		fmt.Println(myArray[i])
	}

	// 遍历方式二
	for index, value := range myArray2 {
		fmt.Println("index:", index, "value:", value)
	}

	// 数组作为形参进行传递的时候，其严格匹配数组类型的（数组长度）
	printArray(myArray2)

	fmt.Println()
	// 数组传递皆为值传递，无法修改数组中元素的值
	fmt.Println("尝试修改数组中元素的值")
	updateArray(myArray)
	printArray(myArray)
}

func updateArray(array [10]int) {
	array[3] = 100
}

func printArray(array [10]int) {
	for _, value := range array {
		fmt.Println("value:", value)
	}
}

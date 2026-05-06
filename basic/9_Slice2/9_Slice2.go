package main

import "fmt"

/*
	学习go中切片的截取和追加
*/

func main() {
	// part1 切片的追加
	var slice = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)

	// 向切片进行追加元素1，len = 4，cap = 5
	slice = append(slice, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)

	// 向切片进行追加元素2，len = 5，cap = 5
	slice = append(slice, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)

	// 向初始化容量已经满的切片进行追加，此时切片会进行扩容，扩容的cap容量是初始化容量的两倍
	slice = append(slice, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)

	fmt.Println("======================================================")
	// part2 切片的截取
	// 使用 : 进行截取为索引的左闭右开
	s2 := slice[1:4]
	fmt.Println(s2)

	// 截取出来的新切片本质上是有一个新指针指向原来的切片
	// 修改原来切片中的元素，截取出来的s2切片的对应元素的值也会发生变化
	slice[1] = 100
	fmt.Println(s2)

	// 可以使用copy进行深拷贝
	fmt.Println("======================================================")
	s3 := make([]int, len(s2), cap(s2))
	copy(s3, s2)
	// 修改s2的其中一个元素的值，s2对应元素发生修改，但是s3还是为copy时的值
	// copy相当于开辟了一个全新的空间，使用一个全新的指针指向这个空间
	s2[1] = 200
	fmt.Println(s2)
	fmt.Println(s3)
}

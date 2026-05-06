package main

/*
	学习不同的导包方式
*/

import (
	"fmt"
	. "go-practice/basic/4_Init/init2"

	// import _ "包"
	// 给包起一个匿名，无法使用导入包的方法但是会执行当前包的init方法
	_ "go-practice/basic/4_Init/init1"
	// 给包起一个别名，可以直接使用别名调用方法
	//test "go-practice/basic/init2"
)

func main() {
	// 使用init1包下的api
	//init1API := init1.Init1API()
	//fmt.Println(init1API)

	// 使用init2包下的api
	init2API := Init2API()
	fmt.Println(init2API)
}

package main

/*
	学习有关init和main函数的执行顺序

	在导入不同的包的时候，会先接着导入包中的导入包进行导入；
	在导入包结束后，会先执行导入的包的init方法

	直到所有的init方法执行完毕，则执行当前文件的main方法
*/

import (
	"fmt"
	"go-practice/basic/4_Init/init1"
	"go-practice/basic/4_Init/init2"
)

func main() {
	// 使用init1包下的api
	init1API := init1.Init1API()
	fmt.Println(init1API)

	// 使用init2包下的api
	init2API := init2.Init2API()
	fmt.Println(init2API)
}

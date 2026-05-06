package init1

import "fmt"

func Init1API() string {
	return "调用执行init1API方法"
}

func init() {
	fmt.Println("正在执行init1包下的init函数")
}

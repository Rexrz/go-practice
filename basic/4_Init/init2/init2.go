package init2

import "fmt"

func Init2API() string {
	return "调用执行init2API方法"
}

func init() {
	fmt.Println("正在执行init2包下的init函数")
}

package main

import (
	"fmt"
	"sync"
)

/*
*
知乎一面 3个协程分工打印1-30；最终输出有序
*/
func main() {
	// 创建控制
	// 打印的最大数值
	const maxNum = 30
	// 协程数量
	const numGoroutines = 3

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// 创建Channel
	chans := make([]chan struct{}, numGoroutines)
	for i := range chans {
		chans[i] = make(chan struct{}, 1)
	}

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// 通过循环开启三个协程
			defer wg.Done()
			// 每个协程打印对应的数字
			// id + 1（1，2，3）
			// +1 +3（4，5，6）
			// +1 +3 +6
			for num := id + 1; num <= maxNum; num += numGoroutines {
				<-chans[id]
				fmt.Printf("goroutine %d -> %d\n", id, num)

				nextID := (id + 1) % numGoroutines
				chans[nextID] <- struct{}{}
			}
		}(i)

	}
	// 向第一个协程发送信号
	chans[0] <- struct{}{}
	// 等待
	wg.Wait()
}

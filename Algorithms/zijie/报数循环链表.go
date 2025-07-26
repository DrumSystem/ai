package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	r := ring.New(100)
	for i := 1; i <= 100; i++ {
		r.Value = i
		r = r.Next()
	}

	counter := 1
	deadCount := 0

	for 100 - deadCount >= n {
		if counter == n {
			// 保存要删除的节点
			toRemove := r
			// 移动到下一个节点
			r = r.Next()
			// 连接前驱和后继节点
			toRemove.Prev().Link(toRemove.Next())
			deadCount++
			counter = 0
			counter++
		}else {
			r = r.Next()
			counter++
		}
	}

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	current := 1 // 当前应该打印的数字
	max := 100
	var wg sync.WaitGroup
	wg.Add(3)

	printNumber := func(num int, next int) {
		defer wg.Done()
		for {
			mu.Lock()
			for current <= max && current%3 != num%3 {
				cond.Wait()
			}
			if current > max {
				mu.Unlock()
				cond.Broadcast()
				return
			}
			fmt.Printf("goroutine%d: %d\n", num%3+1, current)
			current++
			mu.Unlock()
			cond.Broadcast()
		}
	}

	go printNumber(1, 2)
	go printNumber(2, 3)
	go printNumber(0, 1)

	wg.Wait()
	fmt.Println("All goroutines finished")
}
package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func do(taskCh chan int) {
	for {
		select {
		case t, beforeClosed := <-taskCh:
			if !beforeClosed {
				fmt.Println("taskCh has closed!")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		//default:
		//	return
		}
	}
}

func sendTasks() {
	taskCh := make(chan int, 10)
	go do(taskCh)
	for i := 0; i < 1000; i++ {
		taskCh <- i
	}
	close(taskCh)
}

func TestDo(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasks()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
	//sort.Slice()
}
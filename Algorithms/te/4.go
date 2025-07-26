package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	//key := "ljw"
	var wg sync.WaitGroup
	wg.Add(3)
	search := []string{"baidu", "google", "bing"}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)


	for i := 0; i < len(search); i++ {
		childContext, childCancel := context.WithTimeout(ctx, time.Duration(i) * time.Second)
		go func(search string, ctx , parentContext context.Context, childCancel, parentCancel context.CancelFunc) {
			defer wg.Done()
			//time.Sleep(1 * time.Second)
			select {
			case <-childContext.Done():
				fmt.Println(search)
				parentCancel()
			case <- parentContext.Done():
				parentCancel()
				fmt.Println("all late")
			}
		}(search[i], childContext ,ctx, childCancel, cancel)
	}
	wg.Wait()
	fmt.Println("all goroutine finished")
}
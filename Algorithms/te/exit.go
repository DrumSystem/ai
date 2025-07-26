package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	urls := []string{"ljw", "ljp"}
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	res := make(chan struct{}, 1)

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result, _ := handleUrl(ctx, url, res)
			select {
			case <-ctx.Done():
				fmt.Println("超时取消，无图片成功")
			case <-res:
				fmt.Println(result)
				cancel()
			}
		}(urls[i])
	}

	wg.Wait()
	fmt.Println("all goroutine finished")
}

func handleUrl(ctx context.Context ,url string, res chan struct{}) (string, error) {
	randDelay := time.Duration(rand.Intn(5)) * time.Second
	timer := time.NewTimer(randDelay)
	select {
	case <-timer.C:
		res <- struct{}{}
	case <- ctx.Done():
		return "",  errors.New( url + "取消下载")
	}
	return fmt.Sprintf("%s 下载成功", url), nil
}

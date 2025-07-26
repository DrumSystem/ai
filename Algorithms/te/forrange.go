package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	dogch := make(chan struct{}, 1)
	catch := make(chan struct{}, 1)
	fishch := make(chan struct{}, 1)

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-dogch
			fmt.Println("dog")
			catch <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-catch
			fmt.Println("cat")
			fishch <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-fishch
			fmt.Println("fish")
			dogch <- struct{}{}
		}
	}()
	dogch <- struct{}{}
	wg.Wait()
	fmt.Println("all finished")

}
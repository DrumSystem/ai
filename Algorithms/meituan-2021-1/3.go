package main

import "fmt"

func main() {
	var n int
	var x, y int
	var value, actually int
	value = 0
	actually = 0
	fmt.Scanln(&n)
	for i := 1; i < n; i++ {
		fmt.Scanln(&x, &y)
		if x < y {
			value += y
		} else {
			value += x
			actually += x - y
		}
	}
	fmt.Println(value, actually)
}

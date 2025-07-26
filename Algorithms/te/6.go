package main

import "fmt"

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:3]
	fmt.Println(len(b), cap(b))
	b = append(b, 11, 12, 13, 14)
	fmt.Println(len(b), cap(b))
	fmt.Println(a, b)
}

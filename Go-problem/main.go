package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		fmt.Println(&v)
		newArr = append(newArr, &v)
	}
	fmt.Println(newArr)
	for i, v := range newArr {
		fmt.Println("<<", v)
		fmt.Println(*newArr[i])
	}
}
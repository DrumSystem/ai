package main

import (
	"fmt"
	"sort"
)

/*100个人围成一圈，每个人有一个编码，编号从1开始到100。
他们从1开始依次报数，报到为M的人自动退出圈圈，然后下一个人接着从1开始报数，直到剩余的人数小于M。
请问最后剩余的人在原先的编号为多少？
3
59，81
*/

func findLast(num []int, n int) []int {
	var tmp []int
	for i, _ := range num {
		if i+1 == n {
			tmp = append(tmp, num[i+1:]...)
			tmp = append(tmp, num[:i]...)
			return findLast(tmp, n)
		}
	}
	return num
}

func main() {
	var n int
	fmt.Scanln(&n)
	num := make([]int, 100)
	for i := 0; i < 100; i++ {
		num[i] = i+1
	}

	res := findLast(num, n)
	sort.Ints(res)
	for i, v := range res {
		fmt.Print(v)
		if i != len(res)-1 {
			fmt.Print(",")
		}
	}
}
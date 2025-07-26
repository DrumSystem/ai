package main

import "fmt"
var parent []int

func findR(x int) int {
	if x == parent[x] {
		return x
	}
	return findR(parent[x])
}

func unionR(x, y int)  {
	parent[findR(y)] = findR(x)
}

func main() {
	var n int
	fmt.Scanln(&n)

	var r int
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&r)
			tmp[i][j] = r
			if r == 1 {
				unionR(i, j)
			}
		}
	}
	fmt.Println(tmp)
	count := make(map[int]int)
	for _, v := range parent {
		count[v]++
	}
	res := 0
	for _, v := range count {
		if res < v {
			res = v
		}
	}
	fmt.Println(res)
}

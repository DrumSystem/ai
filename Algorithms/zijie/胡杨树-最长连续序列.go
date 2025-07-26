package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, m, k int
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	total := make([]int, n+1)
	for i := 1; i <= n; i++ {
		total[i] = 0
	}

	var dead int
	for i := 0; i < m; i++ {
		fmt.Scan(&dead)
		total[dead] = 1
	}

	fmt.Scanln(&k)

	var res, sumLeft, sumRight, left int
	left = 1

	for i := 1; i <=n ; i++ {
		sumRight += total[i]
		for ; sumRight - sumLeft > k ;  {
			sumLeft += total[left]
			left++
		}

		res = max(res, i-left+1)
		fmt.Println(i, left, res)

	}
	fmt.Println(res)


}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	graph := make([]int, n)
	sumEdge := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&graph[i])
		sumEdge += graph[i]
	}
	edge := sumEdge / 2

	if sumEdge % 2 == 1  {
		fmt.Println(-1)
		return
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, edge+1)
	}

	used := make([][]bool, n)
	for i := range used {
		used[i] = make([]bool, edge+1)
	}


	for j := graph[0]; j <= edge; j++ {
		dp[0][j] = graph[0]
		used[0][j] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= edge; j++ {
			if j < graph[i] {
				dp[i][j] = dp[i-1][j]
			}else {
				if dp[i-1][j] < dp[i-1][j-graph[i]]+graph[i] {
					used[i][j] = true
				}
				dp[i][j] = maxG(dp[i-1][j], dp[i-1][j-graph[i]]+graph[i])
			}
		}
	}
	var left ,right []int
	if dp[n-1][edge] == edge {
		for i := n-1; i >= 0 ; i-- {
			if used[i][edge] {
				left = append(left, i)
				edge = edge - graph[i]
			}else {
				right = append(right, i)
			}
		}
	}
	var res [][]int
	if dp[n-1][sumEdge / 2] == sumEdge / 2 {
		fmt.Println(sumEdge / 2)
		for _, l := range left {
			for _, r := range right {
				tmp := minG(graph[l], graph[r])
				graph[l] -= tmp
				graph[r] -= tmp
				for tmp > 0 {
					tmp--
					//fmt.Println( l+1, r+1)
					tmp := []int{l+1, r+1}
					sort.Ints(tmp)
					res = append(res, tmp)
				}
			}
		}
	}else {
		fmt.Println(-1)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	for i := 0; i < len(res); i++ {
		fmt.Println(fmt.Sprintf("%d %d", res[i][0], res[i][1]))
	}



}

func maxG(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minG (x, y int) int {
	if x > y {
		return y
	}
	return x
}



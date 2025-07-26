package main

import "fmt"

func main() {
	play := make([]int, 10)
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Scan(&play[i])
		sum += play[i]
	}
	targetSum := sum / 2
	dp := make([]int, targetSum+1)

	for i := 0; i < 10; i++ {
		for j := targetSum; j >= play[i] ; j-- {
			dp[j] = maxP(dp[j], dp[j-play[i]] + play[i])
		}
	}
	//fmt.Println(dp[targetSum])
	fmt.Println(sum - 2 * dp[targetSum])
}

func maxP(x, y int) int {
	if x > y {
		return x
	}
	return y
}
package main

import "fmt"

// 店铺dp[i][j]表示i-j区间内能比对手获得的最大美味值

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}
var p []int
//var n int
var dp[][]int
func main() {

	var size int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&size)
		p = append(p, size)
	}

	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var maxVal int
	for i := 0; i < n; i++ {
		//玩家A从0开始选择
		maxVal = max2(maxVal, allocation((i+1)%n, (i-1+n)%n)+ p[i])
	}

	fmt.Println(maxVal)
}

func allocation(l, r int) int {

	if dp[l][r] != -1 {
		return dp[l][r]
	}


	//玩家B选最大的
	if p[l] > p[r] {
		l = (l+1) % n
	}else {
		r = (r-1+n) % n
	}

	//玩家a从剩下的选，看选左边还是右边最大
	if l == r {
		dp[l][r] = p[l]
	}else {
		leftMax := allocation((l+1) % n, r) + p[l]
		rightMax := allocation(l, (r-1+n) % n) + p[r]
		dp[l][r] = max2(leftMax, rightMax)
	}
	return dp[l][r]

}

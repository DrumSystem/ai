package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s [][]int


	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		var tmp []int
		lines := strings.Split(scan.Text(), " ")
		for i := 0; i < len(lines); i++ {
			v, _ := strconv.Atoi(lines[i])
			tmp = append(tmp, v)
		}
		s = append(s, tmp)
	}

	fmt.Println(s)

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s[0]))
	}
	dp[0][0] = s[0][0]
	for i := 1; i < len(s); i++ {
		dp[i][0] = dp[i-1][0] + s[i][0]
	}

	for j := 1; j < len(s[0]); j++ {
		dp[0][j] = dp[0][j-1] + s[0][j]
	}
	fmt.Println(dp)

	for i := 1; i < len(s); i++ {
		for j := 1; j <len(s[0]); j++ {
			dp[i][j] = min(dp[i-1][j] + s[i-1][j], dp[i][j-1]+s[i][j-1])
		}
	}
	fmt.Println(dp[len(s)-1][len(s[0])-1])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
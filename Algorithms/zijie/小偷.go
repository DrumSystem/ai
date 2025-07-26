package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxM(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := strings.Trim(scan.Text(), "[]")
	s = strings.ReplaceAll(s, " ", "")

	money := strings.Split(s, ",")

	dp := make([]int, len(money))
	moneys := make([]int, len(money))

	for i := 0; i < len(money); i++ {
		m, _ := strconv.Atoi(money[i])
		moneys[i] = m
	}
	dp[0], dp[1] = moneys[0], maxM(moneys[0], moneys[1])

	for i := 2; i < len(moneys); i++ {
		dp[i] = maxM(dp[i-1], dp[i-2] + moneys[i])
	}

	fmt.Println(dp[len(moneys)-1])

}
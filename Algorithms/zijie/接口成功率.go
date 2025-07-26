package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	avg, _ := strconv.Atoi(scan.Text())

	scan.Scan()
	s := strings.Split(scan.Text(), " ")
	var time []int
	for i := 0; i < len(s); i++ {
		a, _ := strconv.Atoi(s[i])
		time = append(time, a)
	}

	sum	:= make([]int, len(time))
	sum[0] = time[0]
	for i := 1; i < len(time); i++ {
		sum[i] = time[i] + sum[i-1]
	}

	fmt.Println(time, sum, avg)
	durMap := make(map[int]string)

	var l, maxLen int
	for r := 0; r < len(sum); r++ {

		for (l > 0 && sum[r] - sum[l-1] > avg * (r-l+1)) || (l == 0 && sum[r] > avg * (r-l+1)){
			l++
		}
		if r-l+1 > maxLen {
			maxLen = max1(maxLen, r-l+1)
			durMap[maxLen] = fmt.Sprintf("%d-%d", l, r)
		} else if r-l+1 == maxLen {
			durMap[maxLen] += fmt.Sprintf(" %d-%d", l, r)
		}
	}

	for k, v := range durMap {
		if	k == maxLen {
			fmt.Print(v)
			fmt.Print(" ")
		}
	}

}
package main

import (
	"fmt"
	"strings"
)

func maxL(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var s string
	fmt.Scanln(&s)

	s = strings.Trim(s, `"`)

	index := make(map[byte]int)

	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {

		k, flag := index[s[right]]
		if flag && right > k {
			left = k + 1
		}

		index[s[right]] = right
		maxLen = maxL(maxLen, right - left + 1)
	}

	fmt.Println(maxLen)
}

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	s = strings.ToLower(s)

	var res []string


	for i := 1; i < len(s); i++ {
		tmp := s[i-1]
		count := 1
		for i < len(s) && s[i-1] == s[i]  {
			i++
			count++
		}

		if count == 1 {
			ans := findAfter(s[i+1:], tmp)
			res = append(res, ans)
		}else {
			res = append(res, string(s[i-1]) + strconv.Itoa(count))
		}

	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][1] == res[j][1] {
			return res[i][0] < res[j][0]
		}
		return res[i][1] > res[j][1]
	})

	fmt.Println(res)
}

func findAfter(s string, target byte) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			count++
		}

	}
	fmt.Println(count)
	return string(target) + strconv.Itoa(count)

}
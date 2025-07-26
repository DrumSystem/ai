package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := strings.Split(scan.Text(), " ")
	s := line[0]
	m, _ := strconv.Atoi(line[1])

	n := len(s) / m + 1
	count := 0
	i, j := 0, 0
	res := make([][]string, m)
	for i := range res {
		res[i] = make([]string, n)
	}
	for count < len(s) {
		// 向下移动
		for i < m && j < n && count < len(s) {
			res[i][j] = string(s[count])
			count++
			i++

		}
		j++
		i--

		for i >= 0 && j < n && count < len(s) {
			res[i][j] = string(s[count])
			count++
			i--
		}
		i++
		j++

		for i < m && j < n && count < len(s) {
			res[i][j] = string(s[count])
			count++
			i++

		}
		j++
		i--

		for i >= 0 && j < n && count < len(s) {
			res[i][j] = string(s[count])
			count++
			i--
		}
		i++
		j++

	}
	fmt.Println(res)

}

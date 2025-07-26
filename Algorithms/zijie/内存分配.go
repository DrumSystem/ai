package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line1 := strings.Split(scan.Text(), ",")
	count := make(map[int]int)
	var resSize []int

	for i := 0; i < len(line1); i++ {
		tmp := strings.Split(line1[i], ":")
		size, _ := strconv.Atoi(tmp[0])
		num, _ := strconv.Atoi(tmp[1])
		resSize = append(resSize, size)
		count[size] = num
	}
	sort.Ints(resSize)
	scan.Scan()
	line2 := strings.Split(scan.Text(), ",")
	nums := make([]int, len(line2))
	for i := 0; i < len(line2); i++ {
		v, _ := strconv.Atoi(line2[i])
		nums[i] = v
	}


}

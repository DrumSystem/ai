package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var path string
var res []int
var used []bool

func maxN(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := strings.Split(scan.Text(), ",")
	maxV := 0
	nums := make([]int, len(line))
	used = make([]bool, len(line))
	set := make(map[int]bool)
	for i := 0; i < len(line); i++ {
		v, _ := strconv.Atoi(line[i])
		exist, _ := set[v]
		if v < 1 || v > 9 || exist {
			fmt.Println(-1)
			return
		}
		nums[i] = v
		set[v] = true
		maxV = maxN(maxV, v)
	}
	ex2 := set[2]
	ex5 := set[5]
	ex6 := set[6]
	ex9 := set[9]

	if len(set) != 4 || (ex2 && ex5) || (ex6 && ex9){
		fmt.Println(-1)
		return
	}

	if ex2 && !ex5 {
		nums = append(nums, 5)
	}
	if !ex2 && ex5 {
		nums = append(nums, 2)
	}

	if ex6 && !ex9 {
		nums = append(nums, 9)
	}
	if !ex6 && ex9 {
		nums = append(nums, 6)
	}

	sort.Ints(nums)

	dfsNum(nums, 0)
	sort.Ints(res)
	fmt.Println(res[0:10])

}

func dfsNum(nums []int, startIndex int)  {

	if len(path) == 4 {
		return
	}

	for i := startIndex; i < 4; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = path + strconv.Itoa(nums[i])
		v, _ := strconv.Atoi(path)
		res = append(res, v)
		dfsNum(nums, 0)
		path = path[0:len(path)-1]
		used[i] = false
	}

}

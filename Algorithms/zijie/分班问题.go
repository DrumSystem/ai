package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find1(x int, parent []int) int {
	if parent[x] == x {
		return x
	}
	return parent[find1(parent[x], parent)]
}

func union1(x, y int, parent []int)  {
	parent[find1(y, parent)] = find1(x, parent)
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	scan.Scan()
	s := strings.Split(scan.Text(), " ")

	parent := make([]int, len(s)+1)
	for i := 1; i < len(s)+1; i++ {
		parent[i] = i
	}

	pre := strings.Split(s[0], "/")
	preChild, _ := strconv.Atoi(pre[0])

	for i := 1; i < len(s); i++ {
		tmp := strings.Split(s[i], "/")
		flag := tmp[1]
		child, _ := strconv.Atoi(tmp[0])
		if flag == "Y" {
			union1(child, preChild, parent)
		}
		preChild = child
	}

	res := make(map[int][]int)
	for k, v := range parent {
		if k > 0 {
			res[v] = append(res[v], k)
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
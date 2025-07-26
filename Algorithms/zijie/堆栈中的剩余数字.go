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
	l := strings.Split(scan.Text(), " ")
	nums := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		num, _ := strconv.Atoi(l[i])
		nums[i] = num
	}

	var stack []int
	for i := 0; i < len(nums); i++ {
		stack = push(stack, nums[i])
	}
	fmt.Println(stack)
}

func push(stack []int, sum int) []int {
	//if len(stack) == 0 {
	//	stack = append(stack, sum)
	//	return stack
	//}
	n := sum
	for i := len(stack) - 1; i >= 0 ; i-- {
		sum -= stack[i]
		if sum == 0 {
			stack = stack[:i]
			stack = push(stack, 2 * n)
			return stack
		}else if sum < 0 {
			break
		}
	}
	stack = append(stack, n)
	return stack
}


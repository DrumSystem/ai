package main

import (
	"fmt"
	"strings"
)

func main() {

	var n, k int
	fmt.Scanln(&n)
	fmt.Scanln(&k)


	nums := make([]int, n)
	f := make([]int, n)


	for i := 0; i < n; i++ {
		nums[i] = i+1
	}
	f[0] = 1
	for i := 1; i < n ; i++ {
		f[i] = f[i-1] * i
	}
	fmt.Println(nums, f)

	k--
	var result strings.Builder
	for i := n-1; i >= 0 ; i-- {
		index := k / f[i]
		result.WriteByte(byte('0' + nums[index]))
		copy(nums[index:], nums[index+1:])
		k = k % f[i]
	}
	fmt.Println(result.String())

}
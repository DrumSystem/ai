package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanln(&t)
	score := make([]int, t)
	totalScore := 0
	maxScore := 0
	for i := 0; i < t; i++ {
		fmt.Scan(&score[i])
		totalScore += score[i]
		if maxScore < score[i] {
			maxScore = score[i]
		}
	}
	sort.Ints(score)
	maxK := totalScore / maxScore
	for k := maxK; k > 0 ; k-- {
		if totalScore % k != 0 {
			continue
		}
		s := totalScore / k
		used := make([]bool, t)
		// 检查是否可以分为k个集合，每个集合的和为s
		if backTracking(score, s, k, 0, 0, used) {
			fmt.Println(s)
			break
		}
	}
}

func backTracking(score []int, targetSum, k, sum, startIndex int, used []bool) bool {
	if k == 0 {
		return true
	}

	if targetSum == sum {
		return backTracking(score, targetSum, k-1, 0, 0, used)
	}

	for i := startIndex; i < len(score); i++ {
		if used[i] || sum + score[i] > targetSum {
			continue
		}
		used[i] = true
		if backTracking(score, targetSum, k, sum + score[i], i+1, used) {
			return true
		}
		used[i] = false
	}
	return false
}

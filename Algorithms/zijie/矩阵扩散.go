package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var m, n int

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := strings.Split(scan.Text(), ",")
	m, _ = strconv.Atoi(line[0])
	n, _ = strconv.Atoi(line[1])
	i, _ := strconv.Atoi(line[2])
	j, _ := strconv.Atoi(line[3])
	k, _ := strconv.Atoi(line[4])
	l, _ := strconv.Atoi(line[5])
	fmt.Println(line, m, n, i, j, k, l)

	nums := make([][]int, m)
	for i := range nums {
		nums[i] = make([]int, n)
	}
	nums[i][j] = 1
	nums[k][l] = 1

	visited := make([][]bool, m)
	for i := range nums {
		visited[i] = make([]bool, n)
	}
	count := 2
	time := 0
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if !visited[i][j] && nums[i][j] == 1 {
	//			visited[i][j] = true
	//			dfsNums(nums, visited, i, j)
	//		}
	//	}
	//}
	visited[i][j] = true
	visited[k][l] = true
	tmp := [][]int {{i, j}, {k, l}}
	fmt.Println(tmp, nums)

	for count < m * n {
		if count == m * n {
			break
		}
		length := len(tmp)
		for i := 0; i < length; i++ {
			count, tmp = dfsNums(nums, visited, tmp[i][0], tmp[i][1], count, tmp)
		}
		tmp = tmp[length:]
		time++
	}
	fmt.Println(time)



}

func dfsNums(nums [][]int, visited [][]bool, x, y, count int, tmp [][]int) (int, [][]int) {
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(dir); i++ {
		nextX := x + dir[i][0]
		nextY := y + dir[i][1]

		if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
			continue
		}

		if !visited[nextX][nextY] && nums[nextX][nextY] == 0 {
			visited[nextX][nextY] = true
			count++
			nums[nextX][nextY] = 1
			tmp = append(tmp, []int{nextX, nextY})
		}

	}
	return count, tmp
}
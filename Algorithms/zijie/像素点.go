package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				visited[i][j] = true
				res++
				dfs2(visited, grid, i, j)
			}
		}
	}
	fmt.Println(res)

}

func dfs2(visited [][]bool, grid [][]int, x, y int)  {

	dir := [][]int{{0,1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < 4; i++ {
		nextX := x + dir[i][0]
		nextY := y + dir[i][1]

		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}

		if !visited[nextX][nextY] && grid[nextX][nextY] == 1  {
			visited[nextX][nextY] = true
			dfs2(visited, grid, nextX, nextY)
		}
	}

}

package main

import "fmt"

type Tree struct {
	Val int
	Left *Tree
	Mid *Tree
	Right *Tree
}

func createTree(tree []int) (*Tree, int) {

	root := &Tree{Val: tree[0]}
	i, level := 1, 1
	queue := []*Tree{root}
	n := len(tree)
	for i < n && len(queue) > 0 {
		if len(queue) == 0 {
			level++
		}
		node := queue[0]
		queue = queue[1:]

		if i < n && tree[i] < node.Val - 500 && node.Left == nil{
			node.Left = &Tree{Val: tree[i]}
			queue = append(queue, node.Left)
			i++
		}


		if i < n && tree[i] > node.Val + 500 && node.Right == nil{
			node.Right = &Tree{Val: tree[i]}
			queue = append(queue, node.Right)
			i++
		}

		if i < n && tree[i] >= node.Val - 500 && tree[i] <= node.Val + 500 && node.Mid == nil{
			node.Mid = &Tree{Val: tree[i]}
			queue = append(queue, node.Mid)
			i++
		}

	}
	return root, level
}

func dfs(root *Tree) int {
	if root == nil {
		return 0
	}

	leftLevel := dfs(root.Left)
	midLevel := dfs(root.Mid)
	rightLevel := dfs(root.Right)

	return maxLs(maxLs(leftLevel, midLevel), rightLevel) + 1
}

func main() {

	var n int
	fmt.Scanln(&n)

	tree := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tree[i])
	}
	root, _ := createTree(tree)
	//fmt.Println(level)
	fmt.Println(dfs(root))

}

func maxLs(x, y int) int {
	if x > y {
		return x
	}
	return y
}

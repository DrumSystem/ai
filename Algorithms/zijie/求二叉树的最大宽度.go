package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	val int
	index int
	left *TreeNode
	right *TreeNode
}

func maxW(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func buildTree(s string) *TreeNode {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.TrimPrefix(s, "root=")
	parts := strings.Split(strings.Trim(s, "[]"), ",")
	n := len(parts)

	if n == 0 {
		return nil
	}

	v, _ := strconv.Atoi(parts[0])
	root := &TreeNode{val: v, index: 1}

	var queue []*TreeNode
	queue = append(queue, root)

	i := 1

	for len(queue) > 0 && i < n {

		node := queue[0]
		queue = queue[1:]

		if i < n && parts[i] != "null" {
			leftVal, _ := strconv.Atoi(parts[i])
			node.left = &TreeNode{val: leftVal, index: 2 * node.index}
			queue = append(queue, node.left)
		}
		i++

		if i < n && parts[i] != "null" {
			rightVal, _ := strconv.Atoi(parts[i])
			node.right = &TreeNode{val: rightVal, index: 2 * node.index + 1}
			queue = append(queue, node.right)
		}
		i++

	}
	return root

}

func bfs(root *TreeNode) int {

	maxWidth := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		leftIndex := queue[0].index
		rightIndex := queue[levelSize-1].index
		currentWidth := rightIndex - leftIndex + 1
		if currentWidth > maxWidth {
			maxWidth = currentWidth
		}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

	}
	return maxWidth

}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()

	root := buildTree(s)
	//printTreeBFS(root)

	fmt.Println(bfs(root))
}

func printTreeBFS(root *TreeNode) {
	if root == nil {
		fmt.Println("Tree is empty!")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				fmt.Print("null ")
				continue
			}
			fmt.Print(node.index, " ")
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}
		fmt.Println() // 换行表示一层结束
	}
}
/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode2 struct {
	Val int
	Left *TreeNode2
	Right *TreeNode2
}

func buildTree2(tree []int) *TreeNode2 {
	root := &TreeNode2{Val: tree[0]}
	n := len(tree)
	queue := []*TreeNode2{root}
	i := 1
	for i < n && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]


		if i < n {
			if tree[i] != -1 {
				node.Left = &TreeNode2{Val: tree[i]}
				queue = append(queue, node.Left)
			}else {
				node.Left = &TreeNode2{Val: 0}
				queue = append(queue, node.Left)
			}
		}
		i++

		if i < n {
			if tree[i] != -1 {
				node.Right = &TreeNode2{Val: tree[i]}
				queue = append(queue, node.Right)
			}else {
				node.Right = &TreeNode2{Val: 0}
				queue = append(queue, node.Right)
			}

		}
		i++
	}
	return root
}

func maxT(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfsV(root *TreeNode2) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	leftVal := dfsV(root.Left)
	RightVal := dfsV(root.Right)
	return maxT(leftVal, RightVal) + root.Val
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := strings.Split(scan.Text(), " ")
	var tree []int
	for i := 0; i < len(line); i++ {
		v, _ := strconv.Atoi(line[i])
		tree = append(tree, v)
	}

	root := buildTree2(tree)
	//fmt.Println(tree)
	fmt.Println(dfsV(root))


}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxT(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfsArray(tree []int, index int) int {
	// 检查是否超出数组范围或遇到空节点
	if index >= len(tree) || tree[index] == -1 {
		return 0
	}

	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	leftTime := dfsArray(tree, leftIndex)
	rightTime := dfsArray(tree, rightIndex)

	return maxT(leftTime, rightTime) + tree[index]
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := strings.Split(scan.Text(), " ")
	var tree []int
	for i := 0; i < len(line); i++ {
		v, _ := strconv.Atoi(line[i])
		tree = append(tree, v)
	}

	// 计算传递时间
	totalTime := dfsArray(tree, 0)
	fmt.Println(totalTime)
}
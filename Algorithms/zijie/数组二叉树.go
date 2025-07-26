package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TreeNum struct {
	val int
	left *TreeNum
	right *TreeNum
	parent *TreeNum
	index int
}

var ress []int

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	l := strings.Split(scan.Text(), " ")
	var tree []int
	for i := 0; i < len(l); i++ {
		v, _ := strconv.Atoi(l[i])
		tree = append(tree, v)
	}
	minNode := bTree(tree)
	dfsNode(minNode)
	fmt.Println(ress)

}

func dfsNode(minNode *TreeNum) {
	ress = append(ress, minNode.val)
	if minNode.parent == nil {
		return
	}
	if minNode.parent != nil {
		dfsNode(minNode.parent)
	}

}

func bTree(tree []int) *TreeNum {
	root := &TreeNum{val: tree[0], parent: nil, index: 0}
	queue := []*TreeNum{root}
	n := len(tree)
	minVal := &TreeNum{val: math.MaxUint32}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		leftIndex := 2 * node.index + 1
		if leftIndex < n && tree[leftIndex] != -1 {
			left := &TreeNum{val: tree[leftIndex], parent: node, index: leftIndex}
			node.left = left
			queue = append(queue, node.left)
		}

		rightIndex := 2 * node.index + 2
		if rightIndex < n && tree[rightIndex] != -1 {
			right := &TreeNum{val: tree[rightIndex], parent: node, index: rightIndex}
			node.right = right
			queue = append(queue, node.right)
		}

		if node.left == nil && node.right == nil {
			if minVal.val > node.val {
				minVal = node
			}
		}

	}
	return minVal
}
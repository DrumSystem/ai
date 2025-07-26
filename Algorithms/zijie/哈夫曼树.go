package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode2 struct {
	Val int
	Left *TreeNode2
	Right *TreeNode2
	Height int
}

type minHeap []*TreeNode2

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	if h[i].Val  == h[j].Val {
		return h[i].Height < h[j].Height
	}
	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{})  {
	*h = append(*h, x.(*TreeNode2))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func createTree2(tree []int, h *minHeap) *TreeNode2 {
	var parent *TreeNode2

	for h.Len() > 1 {
		left := heap.Pop(h).(*TreeNode2)
		right := heap.Pop(h).(*TreeNode2)
		parent = &TreeNode2{Val: left.Val + right.Val, Left: left, Right: right, Height: left.Height + 1}
		heap.Push(h, parent)
	}
	return parent
}

func dfsH(root *TreeNode2, res *strings.Builder) string {

	if root == nil {
		return ""
	}
	dfsH(root.Left, res)
	res.WriteString(strconv.Itoa(root.Val) + " ")
	dfsH(root.Right, res)
	return res.String()
}


func main() {
	var n int
	fmt.Scanln(&n)

	tree := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tree[i])
	}

	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, &TreeNode2{Val: tree[i], Height: 0	})
	}


	root := createTree2(tree, h)

	var res strings.Builder
	result := dfsH(root, &res)

	fmt.Println(result)


}
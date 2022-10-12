package goweb

import (
	"strings"
)

type node struct {
	pattern string //待匹配的路由 
	part string //路由的一部分
	wildChild bool // 是否精确匹配， 判断是否含有 * || ：
	children []*node
}

//找到第一个匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || n.wildChild {
			return child
		}
	}
	return nil
}

//寻找所有匹配的节点， 用于查找
func (n *node) matchChildren(part string) []*node {
	var nodes []*node
	for _, child := range n.children {
		if child.part == part || n.wildChild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	if height == len(parts) {
		n.pattern = pattern
		//fmt.Println("insert children", n.pattern)
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{
			part: part,
			wildChild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 查找匹配的节点
func (n *node) search(parts []string, height int) *node {
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	if children == nil {

	}

	for _, child := range children {
		//递归的查询
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil

}
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 使用双链表和map实现

type LRUCache struct {
	Cap int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	K, V int
}

func Construct(cap int) LRUCache {
	return LRUCache{
		cap,
		make(map[int]*list.Element),
		list.New(),
	}
}

func (c *LRUCache) get(key int) int {
	if el, ok := c.Keys[key]; ok  {
		c.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (c *LRUCache) put(key, val int) {
	if el, ok := c.Keys[key]; ok  {
		el.Value = pair{key, val}
		c.List.MoveToFront(el)
	} else {
		el := c.List.PushFront(pair{key, val})
		c.Keys[key] = el
	}

	if c.List.Len() > c.Cap {
		el := c.List.Back()
		c.List.Remove(el)
		delete(c.Keys, el.Value.(pair).K)
	}
}

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s1 := strings.ReplaceAll(scan.Text(), "[", "")
	s1 = strings.ReplaceAll(s1, "]", "")
	s1 = strings.ReplaceAll(s1, `"`, "")
	s1 = strings.ReplaceAll(s1, " ", "")
	commands := strings.Split(s1, ",")
	scan.Scan()
	s2 := strings.ReplaceAll(scan.Text(), "[", "")
	s2 = strings.ReplaceAll(s2, "]", "")
	s2 = strings.ReplaceAll(s2, " ", "")
	nums := strings.Split(s2, ",")

	fmt.Println(commands)
	fmt.Println(nums)

	lru := Construct(0)
	var res strings.Builder
	j := 0
	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "LRUCache":
			cap, _ := strconv.Atoi(nums[j])
			lru = Construct(cap)
			res.WriteString("null, ")
			//res = append(res, "null, ")
			j++
		case "get":
			key, _ := strconv.Atoi(nums[j])
			val := lru.get(key)
			res.WriteString(strconv.Itoa(val) + ", ")
			//res = append(res, strconv.Itoa(val) + ", ")
			j++
		case "put":
			key, _ := strconv.Atoi(nums[j])
			val, _ := strconv.Atoi(nums[j+1])
			lru.put(key, val)
			res.WriteString("null, ")
			//res = append(res, "null, ")
			j += 2
		}
	}
	fmt.Println("[" + strings.TrimSuffix(res.String(), ", ") + "]")


}
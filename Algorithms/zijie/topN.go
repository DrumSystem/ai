package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type URL struct {
	Count int
	Url string
}

type Heap []URL

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Url < h[j].Url
	}
	return h[i].Count > h[j].Count
}

func (h Heap) Swap(i, j int)  {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{})  {
	*h = append(*h, x.(URL))
}

func (h * Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main() {

	scan := bufio.NewScanner(os.Stdin)
	urlCount := make(map[string]int)

	for scan.Scan() {
		line := scan.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		if n, err := strconv.Atoi(line); err == nil {
			if n <= 0 || n > len(urlCount) {
				continue // 无效的N值
			}

			h := &Heap{}
			heap.Init(h)
			for url, count := range urlCount{
				heap.Push(h, URL{count, url})
			}

			var res []string
			for i := 0; i < n && h.Len() > 0; i++ {
				item := heap.Pop(h).(URL)
				res = append(res, item.Url)
			}
			fmt.Println(strings.Join(res, ","))
		}else {
			urlCount[line]++
		}
	}
}

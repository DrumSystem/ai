package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Sport struct {
	id int
	val int
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()
	ss := strings.Split(s, " ")
	var sports []*Sport
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.Atoi(ss[i])
		sport := &Sport{id: i, val: v}
		sports = append(sports, sport)
	}

	win := make([]*Sport, len(sports))
	copy(win, sports)
	for len(win) > 4 {
		var tmp	[]*Sport
		for i := 1; i < len(win); i += 2 {
			if win[i].val > win[i-1].val {
				tmp = append(tmp, win[i])
			}else if win[i].val < win[i-1].val {
				tmp = append(tmp, win[i-1])
			}else {
				if win[i].id < win[i-1].id {
					tmp = append(tmp, win[i])
				}else {
					tmp = append(tmp, win[i-1])
				}
			}
		}
		if len(win) % 2 == 1 {
			tmp = append(tmp, win[len(win)-1])
		}
		win = win[:0]
		win = append(win, tmp...)
	}
	//fmt.Println(win)
	var lose []*Sport
	var tmp	[]*Sport
	for i := 1; i < len(win); i += 2 {
		if win[i].val > win[i-1].val {
			tmp = append(tmp, win[i])
			lose = append(lose, win[i-1])
		}else if win[i].val < win[i-1].val {
			tmp = append(tmp, win[i-1])
			lose = append(lose, win[i])
		}else {
			if win[i].id < win[i-1].id {
				tmp = append(tmp, win[i])
				lose = append(lose, win[i-1])
			}else {
				tmp = append(tmp, win[i-1])
				lose = append(lose, win[i])
			}
		}
	}
	if len(win) % 2 == 1 {
		tmp = append(tmp, win[len(win)-1])
	}
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].val == tmp[j].val {
			return tmp[i].id < tmp[j].id
		}
		return tmp[i].val > tmp[j].val
	})
	sort.Slice(lose, func(i, j int) bool {
		if lose[i].val == lose[j].val {
			return lose[i].id < lose[j].id
		}
		return lose[i].val > lose[j].val
	})
	fmt.Println(tmp[0].id, tmp[1].id, lose[0].id)

}

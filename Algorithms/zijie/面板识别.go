package main

import (
	"fmt"
	"math"
	"sort"
)

type Lamp struct {
	id int
	centerX int
	centerY int
	height int
}

func main() {
	var n int
	var id, x1, y1, x2, y2 int
	fmt.Scanln(&n)
	var info []*Lamp
	for i := 0; i < n; i++ {
		fmt.Scan(&id, &x1,  &y1, &x2, &y2)
		lamp := &Lamp{id, (x1 + x2) / 2, (y1 + y2)/2, y2 - y1}
		info = append(info, lamp)
	}
	sort.Slice(info, func(i, j int) bool {
		return info[i].centerY < info[j].centerY
	})
	for _, v := range info {
		fmt.Println(v.id, v.centerY, v.centerX)
	}

	base := info[0]
	var sameLamp []*Lamp
	var res []int
	sameLamp = append(sameLamp, info[0])

	for i := 1; i < n; i++ {
		curLamp := info[i]
		if math.Abs(float64(curLamp.centerY - base.centerY)) <= float64(base.height) / 2 {
			sameLamp = append(sameLamp, curLamp)
		}else {
			sort.Slice(sameLamp, func(i, j int) bool {
				return sameLamp[i].centerX < sameLamp[j].centerX
			})
			for _, v := range sameLamp {
				res = append(res, v.id)
			}
			sameLamp = sameLamp[:0]
			base = curLamp
			sameLamp = append(sameLamp, base)
		}
	}

	if len(sameLamp) > 0 {
		sort.Slice(sameLamp, func(i, j int) bool {
			return sameLamp[i].centerX < sameLamp[j].centerX
		})
		for _, v := range sameLamp {
			res = append(res, v.id)
		}
	}

	fmt.Println(res)
}

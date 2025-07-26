package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	fmt.Scanln(&s)
	avg := len(s) / 4
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for k, v := range sMap {
		if v > avg {
			tMap[k] = v - avg
		}
	}

	check := func() bool {
		for _, v := range sMap {
			if v > avg {
				return  false
			}
		}
		return true
	}

	ansL, ansR := 0, 0
	length := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]]--
		for check() && l <= r {
			if r - l + 1 < length {
				length = r - l + 1
				ansL = l
				ansR = l + length
			}
			sMap[s[l]]++
			l++
		}
	}
	fmt.Println(s[ansL:ansR])


}




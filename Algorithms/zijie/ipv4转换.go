package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scanln(&s)
	sli := strings.Split(s, "#")
	if len(sli) != 4 {
		fmt.Println("invalid IP")
		return
	}

	for i := 0; i < len(sli); i++ {
		if !isNumeric(sli[i]) {
			fmt.Println("invalid IP")
			return
		}
		if len(sli[i]) == 0 {
			fmt.Println("invalid IP")
			return
		}

		if len(sli[i]) > 1 && sli[i][0] == '0' {
			fmt.Println("invalid IP")
			return
		}
	}

	//firstSeg, err := strconv.ParseInt(sli[0], 10, 64)
	firstSeg, err := strconv.Atoi(sli[0])
	if err != nil || (firstSeg < 1 || firstSeg > 128){
		fmt.Println("invalid IP")
		return
	}

	for i := 1; i < 4; i++ {
		num, err := strconv.Atoi(sli[i])
		if err != nil || (num < 0 || num > 255) {
			fmt.Println("invalid IP")
			return
		}
	}
	var res int
	for i := 0; i < 4; i++ {
		ipv4, _ := strconv.Atoi(sli[i])
		res = res * 256 + ipv4
	}
	fmt.Println(res)


}

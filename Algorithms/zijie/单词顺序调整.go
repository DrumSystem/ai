package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var s string

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s = scan.Text()

	words := strings.Split(s, " ")
	for i, word := range words {
		tmp := []byte(word)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		words[i] = string(tmp)
	}
	fmt.Println(words)


}

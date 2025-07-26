package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "root = 1,2,3"
	s = strings.TrimPrefix(s, "root = ")
	fmt.Println(s)
}

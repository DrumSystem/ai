package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	m, _ := strconv.Atoi(scan.Text())
	scan.Scan()
	fileIds := strings.Split(scan.Text(), " ")
	scan.Scan()
	fileSize := strings.Split(scan.Text(), " ")

	value := make(map[int]int)
	ori := make(map[int]int)
	for i := 0; i < len(fileIds); i++ {
		id, _ := strconv.Atoi(fileIds[i])
		size, _ := strconv.Atoi(fileSize[i])
		value[id] += size
		ori[id] = size
	}
	res := 0
	for k, v := range value {
		if v > m + ori[k] {
			res += m + ori[k]
		}else {
			res += v
		}
	}
	fmt.Println(res)
}

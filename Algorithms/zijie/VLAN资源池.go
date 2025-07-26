package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseVlanPool(res string) []int {
	s := strings.Split(res, ",")
	var vlanPool []int
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "-") {
			vlanItems := strings.Split(s[i], "-")
			start, _ := strconv.Atoi(vlanItems[0])
			end, _ := strconv.Atoi(vlanItems[1])
			for j := start; j <= end; j++ {
				vlanPool = append(vlanPool, j)
			}
		}else{
			vlan, _ := strconv.Atoi(s[i])
			vlanPool = append(vlanPool, vlan)
		}
	}
	return vlanPool
}

func DeleteSlice(a []int, elem int) []int {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

func formatVlanPool(vlanPool []int) string {

	if len(vlanPool) == 0 {
		return ""
	}
	var result strings.Builder
	start := vlanPool[0]
	prev := vlanPool[0]
	for i := 1; i < len(vlanPool); i++ {
		if vlanPool[i] == prev + 1 {
			prev = vlanPool[i]
		}else {
			if start == prev {
				result.WriteString(strconv.Itoa(start))
			}else {
				result.WriteString(strconv.Itoa(start) + "-" + strconv.Itoa(prev))
			}
			result.WriteString(",")
			start = vlanPool[i]
			prev = vlanPool[i]
		}

	}

	if start == prev {
		result.WriteString(strconv.Itoa(start))
	}else {
		result.WriteString(strconv.Itoa(start) + "-" + strconv.Itoa(prev))
	}

	return result.String()
}

func main() {
	//scan := bufio.NewScanner(os.Stdin)
	//scan.Scan()
	//resource := scan.Text()
	//scan.Scan()
	//allRes := scan.Text()

	var res string
	var allRes int
	fmt.Scanln(&res)
	fmt.Scanln(&allRes)

	vlanPool := parseVlanPool(res)
	sort.Ints(vlanPool)
	vlanPool = DeleteSlice(vlanPool, allRes)
	fmt.Println(formatVlanPool(vlanPool))

}

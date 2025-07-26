package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	name string
	//score int
	totalScore int
	subject map[string]int
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	l1 := strings.Split(scan.Text(), " ")
	n, _ := strconv.Atoi(l1[0])
	m, _ := strconv.Atoi(l1[1])
	scan.Scan()
	sub := strings.Split(scan.Text(), " ")
	var res []*Student
	for i := 0; i < n; i++ {
		scan.Scan()
		stu := strings.Split(scan.Text(), " ")
		student := &Student{
			name: stu[0],
			totalScore: 0,
			subject: make(map[string]int),
		}
		for j := 1; j <= m; j++ {
			score, _ := strconv.Atoi(stu[j])
			student.subject[sub[j-1]] = score
			student.totalScore += score
		}
		res = append(res, student)
	}
	scan.Scan()
	querySub := scan.Text()
	if querySub == ""{
		sort.Slice(res, func(i, j int) bool {
			if res[i].totalScore == res[j].totalScore {
				return res[i].name < res[j].name
			}
			return res[i].totalScore > res[j].totalScore
		})
	}else {
		sort.Slice(res, func(i, j int) bool {
			if res[i].subject[querySub] == res[j].subject[querySub] {
				return res[i].name < res[j].name
			}
			return res[i].subject[querySub] > res[j].subject[querySub]
		})
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].name, res[i].totalScore)
	}




}

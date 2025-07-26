package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type People struct {


	id int
	time int
	dis int
	actualNum string
	registerNum string
	order int

}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	n, _ := strconv.Atoi(scan.Text())
	var ans []*People
	pMap := make(map[int][]*People)
	for i := 0; i < n; i++ {
		scan.Scan()
		record := strings.Split(scan.Text(), ",")
		id, _  := strconv.Atoi(record[0])
		time, _  := strconv.Atoi(record[1])
		dis, _  := strconv.Atoi(record[2])
		actualNum := record[3]
		registerNum := record[4]
		peo := &People{id: id, time: time, dis: dis, actualNum: actualNum, registerNum: registerNum, order: i}
		pMap[id] = append(pMap[id], peo)


	}


	for _, records := range pMap {
		sort.Slice(records, func(i, j int) bool {
			return records[i].time < records[j].time
		})
		if len(records) > 1 {
			for i := 0; i < len(records); i++ {
				for j := 0; j < len(records); j++ {
					if j == i {
						continue
					}
					if (math.Abs(float64(records[i].time - records[j].time)) <= 60 && math.Abs(float64(records[i].dis - records[j].dis)) > float64(5)) ||
						records[i].actualNum != records[i].registerNum {
						ans = append(ans, records[i])
						j = len(records)
					}
				}

			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].order < ans[j].order
	})
	var res strings.Builder
	for i, record := range ans {
		res.WriteString(strconv.Itoa(record.id) + "," + strconv.Itoa(record.time) +  "," + strconv.Itoa(record.dis) +  "," +
			record.actualNum +  "," +record.registerNum)
		if i != len(ans) - 1 {
			res.WriteString(";")
		}
	}
	if len(ans) == 0 {
		fmt.Println("null")
	}else {
		fmt.Println(res.String())
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//f4 06 01 00 01 00 02 0f
//f5 07 01 00 01 00 02 10 11
//f5 05 01 00 01 00 02 [10 11 0f]
//
//f4
//07
//02
//00
//01
//01
//00
//02
//0f
//
//f4 07 02 00 01 01 00 02 0f
//f5 0d 02
//00 01 00 02 10 11
//01 01 00 02 11 13
//
//f5 0d 02
//00 01 00 02 10 11
//01 01 00 02 10 11
//01 00 02 11 13
//查询2个数组(编号1和2)(这里应该是0， 1) 的1段数据，起始数据标识是0，个数2，关键数据0x11；
//根据示例的服务器数据，返回的响应帧包含 2个数组的数据；
//其中数组1（数组0），1个数据段，起始标识1（0），数据个数1（2），数据内容0x13（10， 11）；
//数组2（数组1），1个数据段，起始标识0，数据个数2，数据内容0x12 0x14 （11， 13）

func main() {
	var data1 = []string{"10", "11","0f", "13", "14", "15", "16", "0f", "18", "19" }
	var data2 = []string{"11", "13","0f", "17", "19", "21", "23", "0f", "27", "29" }
	var data3 = []string{"12", "14","0f", "18", "20", "22", "24", "0f", "28", "30" }
	ori := [][]string{data1, data2, data3}
	var s string
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s = scan.Text()
    frame := strings.Split(s, " ")
	index := 2
	arrs, _ := strconv.ParseInt(frame[2], 16, 64)
	var arrId []int
	var arrId16 []string
	for i := index + 1; i <= index + int(arrs); i++ {
		id, _ := strconv.ParseInt(frame[i], 16, 64)
		arrId = append(arrId, int(id))
		arrId16 = append(arrId16, frame[i])
	}
	index += int(arrs) + 1
	duan, _ := strconv.ParseInt(frame[index], 16, 64)
	duanId := make(map[int][]string)
	//var duanId16 [][]int
	count := int64(-1)
	for i := index + 1; i <= index + int(duan) * 2; i += 2 {
		sid, _ := strconv.ParseInt(frame[i], 16, 64)
		count, _ = strconv.ParseInt(frame[i+1], 16, 64)
		duanId[0] = append(duanId[0], fmt.Sprintf("%s %s %s %s",frame[index], frame[i], frame[i+1], strings.Join(ori[0][sid:sid+count], " ")))
		duanId[1] = append(duanId[1], fmt.Sprintf("%s %s %s %s",frame[index], frame[i], frame[i+1], strings.Join(ori[1][sid:sid+count], " ")))
		duanId[2] = append(duanId[2], fmt.Sprintf("%s %s %s %s",frame[index], frame[i], frame[i+1], strings.Join(ori[2][sid:sid+count], " ")))
		//duanId = append(duanId, []int{int(sid), int(count), int(sid)+int(count)-1})
	}
	//keyNum := frame[index + int(duan) * 2+1]

	var res strings.Builder
	res.WriteString("f5" + " ")
	frameL := 1 + int(arrs) * (1 + 1 + int(duan) * (2 + int(count)))
	//fmt.Println(frameL)
	res.WriteString(fmt.Sprintf("%02x", frameL) + " ")
	res.WriteString(frame[2] + " ")
	for i := 0; i < len(arrId) ; i++ {
		res.WriteString(arrId16[i] + " ")
		for j := 0; j < len(duanId[0]); j++ {
			res.WriteString(strings.Join(duanId[arrId[i]], " ") + " ")
		}
	}
	fmt.Println(res.String())






}

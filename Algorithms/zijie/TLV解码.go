package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*题目描述

TLV编码是按[Tag Length Value]格式进行编码的，一段码流中的信元用Tag标识，Tag在码流中唯一不重复，Length表示信元Value的长度，Value表示信元的值。

码流以某信元的Tag开头，Tag固定占一个字节，Length固定占两个字节，字节序为小端序。

现给定TLV格式编码的码流，以及需要解码的信元Tag，请输出该信元的Value。

输入码流的16进制字符中，不包括小写字母，且要求输出的16进制字符串中也不要包含小写字母；码流字符串的最大长度不超过50000个字节。*/


func main() {
	//var tag int
	//var s string
	//fmt.Scan(&tag)
	//fmt.Scanln(&s)
	//fmt.Println(s)

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	tag := scan.Text()
	scan.Scan()
	s := scan.Text()


	hexPre := strings.Split(s, " ")
	//fmt.Println(hexPre, s)
	index := 0
	for ; index < len(hexPre)-2; {
		//res, _ := strconv.Atoi(hexPre[index])
		res := hexPre[index]
		length, _ := strconv.ParseInt(hexPre[index+2] + hexPre[index+1], 16, 64)
		l := int(length)
		if res == tag {
			var result strings.Builder
			for i := index+3; i < index+3+l; i++ {
				result.WriteString(hexPre[i])
				result.WriteByte(' ')
			}
			fmt.Println(result.String())
			break
		}else {
			index += 2 + l + 1
		}
	}
}
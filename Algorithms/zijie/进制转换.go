package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encodeNumber(numStr string) string {
	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		return "" // 根据实际需求处理错误，题目保证输入合法
	}

	var encoded []byte
	for {
		// 取低7位
		sevenBits := num & 0x7F
		num >>= 7
		if num != 0 {
			// 还有后续字节，最高位置1
			sevenBits |= 0x80
		}
		encoded = append(encoded, byte(sevenBits))
		if num == 0 {
			break
		}
	}

	// 转换为16进制字符串
	var hexStr strings.Builder
	for _, b := range encoded {
		hexStr.WriteString(fmt.Sprintf("%02X", b))
	}
	return hexStr.String()
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(encodeNumber(input))
}
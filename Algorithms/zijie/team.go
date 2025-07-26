package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
//总共有 n 个人在机房，每个人有一个标号（1<=标号<=n），他们分成了多个团队，需要你根据收到的 m 条消息判定指定的两个人是否在一个团队中，具体的：
//
//消息构成为 a b c，整数 a、b 分别代表两个人的标号，整数 c 代表指令
//c == 0 代表 a 和 b 在一个团队内
//c == 1 代表需要判定 a 和 b 的关系，如果 a 和 b 是一个团队，输出一行’we are a team’,如果不是，输出一行’we are not a team’
//c 为其他值，或当前行 a 或 b 超出 1~n 的范围，输出‘da pian zi’

func find(x int, parent []int) int {
	if parent[x] == x {
		return x
	}
	return parent[find(parent[x], parent)]
}

func union(x, y int, parent []int)  {
	parent[find(y, parent)] = find(x, parent)
}

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := strings.Split(scan.Text(), " ")
	n, _ := strconv.Atoi(line[0])
	m, _ := strconv.Atoi(line[1])
	rel := make([][]int, n+1)
	for i := range rel {
		rel[i] = make([]int, n+1)
	}

	parent := make([]int, n+1)
	for i := range	parent {
		parent[i] = i
	}
	for i:= 0; i < m; i++ {
		scan.Scan()
		msgs := strings.Split(scan.Text(), " ")
		a, _ := strconv.Atoi(msgs[0])
		b, _ := strconv.Atoi(msgs[1])
		val, _ := strconv.Atoi(msgs[2])
		if a < 1 || a > n || b > n || b < 1 {
			fmt.Println("Null")
			continue
		}

		if val != 0 && val != 1 {
			fmt.Println("da pian zi")
			continue
		}

		/*switch val {
			case 0:
				// 建立双向关系
				rel[a][b] = 1
				rel[b][a] = 1
			case 1:
				if rel[a][b] == 1 || rel[b][a] == 1 {
					fmt.Println("We are a team")
				} else {
					fmt.Println("We are not a team")
				}
			default:
				fmt.Println("da pian zi")
		}*/

		switch val {
		case 0:
			union(a, b, parent)
		case 1:
			if find(a, parent) == find(b, parent) {
				fmt.Println("We are a team")
			} else {
				fmt.Println("We are not a team")
			}
		default:
			fmt.Println("da pian zi")

		}
	}

}
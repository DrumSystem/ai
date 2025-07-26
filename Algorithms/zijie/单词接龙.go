package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var start, n int
	var s, target string




	fmt.Scanln(&start)
	fmt.Scanln(&n)

	var words []string
	wMap := make(map[byte][]string)

	for i := 0; i < n; i++ {
		fmt.Scanln(&s)
		wMap[s[0]] = append(wMap[s[0]], s)
		if i == start {
			target = s
			continue
		}
		words = append(words, s)

	}

	var res strings.Builder
	res.WriteString(target)

	lastChar := target[len(target)-1]

	for  {
		nextWord := findWord(wMap, lastChar)
		if nextWord == "" {
			break
		}
		res.WriteString(nextWord)
		lastChar = nextWord[len(nextWord)-1]
		removeWord(wMap, nextWord)

	}
	fmt.Println(res.String())

}

func findWord(wMap map[byte][]string, lastChar byte) string {
	candidates := wMap[lastChar]

	if len(candidates) == 0 {
		return ""
	}

	sort.Slice(candidates, func(i, j int) bool {
		leni, lenj := len(candidates[i]), len(candidates[j])
		if leni != lenj {
			return leni > lenj
		}
		return candidates[i] < candidates[j]
	})
	return candidates[0]
}

func removeWord(wMap map[byte][]string, nextWord string)  {
	words := wMap[nextWord[0]]
	for i, w := range words {
		if w == nextWord {
			wMap[nextWord[0]] = append(words[:i], words[i+1:]...)
			break
		}
	}
}

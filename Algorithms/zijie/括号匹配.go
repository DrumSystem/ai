package main

import (
	"container/list"
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)

	stack := list.New()

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '{', '(':
			stack.PushBack(s[i])
		case ']', '}', ')':
			if stack.Len() == 0 {
				fmt.Println("false")
				return
			}

			top := stack.Back().Value.(byte)
			stack.Remove(stack.Back())

			if !((top == '[' && s[i] == ']') ||
				(top == '{' && s[i] == '}') ||
				(top == '(' && s[i] == ')')) {
				fmt.Println("false")
				return
			}
		}
	}

	if stack.Len() == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
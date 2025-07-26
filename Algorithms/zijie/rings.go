package main

import (
	"container/ring"
	"fmt"
)



func main() {
	r := ring.New(100)
	r1 := ring.New(3)

	for i := 1; i <= 100; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 4; i <= 6; i++ {
		r1.Value = i
		r1 = r1.Next()
	}
	pre := r.Prev()
	r.Link(pre.Move(2))


	r.Do(func(i interface{}) {
		fmt.Print(i)
		fmt.Print(",")
	})


}
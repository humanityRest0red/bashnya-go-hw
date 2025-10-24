package main

import (
	"fmt"
	"hw3/stack"
)

func main() {
	s := stack.New(1.3, 2.3, -3.14)
	// for i := range 10 {
	// 	s.Push(i)
	// }
	fmt.Printf("%##v\n", s)
}

// elem, _ := s.Pop()
// println(elem)

// 	d := deque.Deque{}
// 	d.PushBack(3)
// 	d.PushBack(1)
// 	println(d.PopFront())
// }

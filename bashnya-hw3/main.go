package main

import (
	"bstree"
	"fmt"
)

// func main() {
// 	s := stack.New(1.3, 2.3, -3.14)
// 	// for i := range 10 {
// 	// 	s.Push(i)
// 	// }
// 	fmt.Printf("%##v\n", s)
// }

// func main() {
// elem, _ := s.Pop()
// println(elem)

// 	d := deque.Deque{}
// 	d.PushBack(3)
// 	d.PushBack(1)
// 	println(d.PopFront())
// }

// func main() {
// 	b := bstree.New(2, 3)
// 	for v := range b.Iterator() {
// 		print(v)
// 	}
// }

func main() {
	bb := bstree.New(1, 2, 3)
	// bb := bstree.BSTree[int]{}
	test(*bb)
	fmt.Printf("%p\n", &bb)
}

func test(b bstree.BSTree[int]) {
	fmt.Printf("%p\n", &b)
}

package main

import "bstree"

func main() {
	// 	d := deque.Deque{}
	// 	d.PushBack(3)
	// 	d.PushBack(1)
	// 	println(d.PopFront())

	b := bstree.New(1, 2, 3, 4, 10, 5, 0, -1)
	for v := range b.Iterator() {
		print(v, " ")
	}

	//		s := stack.New(1.3, 2.3, -3.14)
	//		s.Clear()
	//		s.Push(2.1)
	//		val, _ := s.Pop()
	//		fmt.Println(val)
}

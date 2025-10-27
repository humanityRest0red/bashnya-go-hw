package bstree

import (
	"cmp"
	"time"
)

type BSTree[T cmp.Ordered] struct {
	root *node[T]
}

func New[T cmp.Ordered](elements ...T) *BSTree[T] {
	b := BSTree[T]{}
	for _, elem := range elements {
		b.Insert(elem)
	}
	return &b
}

func (b *BSTree[T]) Insert(elem T) {
	if b.root == nil {
		b.root = NewNode(elem)
		return
	}

	current := b.root
	for current != nil {
		if elem < current.key {
			if current.left == nil {
				current.left = NewNode(elem)
				current.left.parent = current
				return
			} else {
				current = current.left
			}
		} else if elem >= current.key {
			if current.right == nil {
				current.right = NewNode(elem)
				current.right.parent = current
				return
			} else {
				current = current.right
			}
		}
	}
}

func (b *BSTree[T]) Remove(elem T) {
	// current := b.root
	// for current != nil {
	// 	if elem < current.key {
	// 		current = current.left
	// 	} else if elem > current.key {
	// 		current = current.right
	// 	} else {
	// 		parent := current.parent
	// 		left := current.left
	// 		right := current.right

	// 	}
	// }
}

func (b *BSTree[T]) Find(elem T) bool {
	current := b.root
	for current != nil {
		if elem < current.key {
			current = current.left
		} else if elem > current.key {
			current = current.right
		} else {
			return true
		}
	}

	return false
}

func (b *BSTree[T]) Depth() uint {
	return b.root.Depth()
}

func (b *BSTree[T]) Iterator() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)

		if b.root == nil {
			return
		}

		current := b.root

		on_left_side := false
		if current.left != nil {
			for current.left != nil {
				current = current.left
			}
			on_left_side = true
		}
		for {
			ch <- current.key
			if current.right != nil {
				current = current.right
				if current.left != nil {
					for current.left != nil {
						current = current.left
					}
					on_left_side = true
				}
			} else if on_left_side && current.parent != nil {
				current = current.parent
				on_left_side = false
			} else {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}()

	return ch
}

func (a *BSTree[T]) Eq(b *BSTree[T]) bool {
	ch1 := a.Iterator()
	ch2 := b.Iterator()

	v1, ok1 := <-ch1
	v2, ok2 := <-ch2
	for ok1 || ok2 {
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
	}
	return true
}

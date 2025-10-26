package bstree

import "cmp"

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

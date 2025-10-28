package bstree

import "cmp"

type node[T cmp.Ordered] struct {
	key    T
	left   *node[T]
	right  *node[T]
	parent *node[T]
}

func NewNode[T cmp.Ordered](elem T) *node[T] {
	return &node[T]{key: elem, left: nil, right: nil, parent: nil}
}

func (n *node[T]) Depth() uint {
	if n == nil {
		return 0
	}
	return 1 + myUintMax(n.left.Depth(), n.right.Depth())
}

func myUintMax(a, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}

// функция для поиска минимального узла в поддереве
func minimum[T cmp.Ordered](node *node[T]) *node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

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
	current := b.root
	for current != nil {
		if elem < current.key {
			current = current.left
		} else if elem > current.key {
			current = current.right
		} else {
			b.deleteNode(current)
			break
		}
	}
}

func (b *BSTree[T]) deleteNode(node *node[T]) {
	if node == nil {
		return
	}

	// Случай 1: узел — лист
	if node.left == nil && node.right == nil {
		if node.parent == nil {
			b.root = nil
		} else if node == node.parent.left {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		return
	}

	// Случай 2: один ребёнок
	if node.left == nil {
		b.transplant(node, node.right)
	} else if node.right == nil {
		b.transplant(node, node.left)
	} else {
		// Случай 3: два ребёнка
		successor := minimum(node.right)
		if successor.parent != node {
			b.transplant(successor, successor.right)
			successor.right = node.right
			successor.right.parent = successor
		}
		b.transplant(node, successor)
		successor.left = node.left
		successor.left.parent = successor
	}
}

func (b *BSTree[T]) transplant(u, v *node[T]) {
	if u.parent == nil {
		b.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
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

		if current.left != nil {
			for current.left != nil {
				current = current.left
			}
		}
		for {
			ch <- current.key
			if current.right != nil {
				current = current.right
				if current.left != nil {
					for current.left != nil {
						current = current.left
					}
				}
			} else if current.parent != nil && current == current.parent.left {
				current = current.parent
			} else {
				break
			}
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

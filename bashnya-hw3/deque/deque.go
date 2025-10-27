package deque

import "errors"

var ErrEmptyDeque = errors.New("empty deque")

type Deque[T comparable] struct {
	data []T
}

func New[T comparable](elements ...T) *Deque[T] {
	return &Deque[T]{
		data: elements,
	}
}

func (d *Deque[T]) PushFront(elem T) {
	new_data := make([]T, d.Size()+1)
	new_data[0] = elem

	for i := range d.Size() {
		new_data[i+1] = d.data[i]
	}

	d.data = new_data
}

func (d *Deque[T]) PushBack(elem T) {
	d.data = append(d.data, elem)
}

func (d *Deque[T]) PopFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, ErrEmptyDeque
	}
	elem := d.data[0]
	d.data = d.data[:1]

	return elem, nil
}

func (d *Deque[T]) PopBack() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, ErrEmptyDeque
	}
	last_i := d.Size() - 1
	elem := d.data[last_i]
	d.data = d.data[:last_i]

	return elem, nil
}

func (d *Deque[T]) IsEmpty() bool {
	return d.Size() == 0
}

func (d *Deque[T]) Size() uint {
	return uint(len(d.data))
}

func (d *Deque[T]) Clear() {
	d.data = d.data[:0]
}

func (a *Deque[T]) Eq(b *Deque[T]) bool {
	if a.Size() != b.Size() {
		return false
	}

	for i := range a.Size() {
		if a.data[i] != b.data[i] {
			return false
		}
	}
	return true
}

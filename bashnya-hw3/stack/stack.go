package stack

import "errors"

var ErrEmptyStack = errors.New("empty stack")

type Stack[T comparable] struct {
	data []T
}

func New[T comparable](elements ...T) *Stack[T] {
	return &Stack[T]{
		data: elements,
	}
}

func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrEmptyStack
	}

	last_i := s.Size() - 1
	elem := s.data[last_i]
	s.data = s.data[:last_i]

	return elem, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrEmptyStack
	}

	return s.data[s.Size()-1], nil
}

func (s *Stack[T]) Size() uint {
	return uint(len(s.data))
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}

func (a *Stack[T]) Eq(b *Stack[T]) bool {
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

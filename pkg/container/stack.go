package container

import "container/list"

type Stack[T any] interface {
	Len() int
	IsEmpty() bool
	Peek() T
	Push(v T)
	Pop() T
}

type stack[T any] struct {
	l *list.List
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{l: list.New()}
}

func (s *stack[T]) Len() int {
	return s.l.Len()
}

func (s *stack[T]) IsEmpty() bool {
	return s.l.Len() == 0
}

func (s *stack[T]) Peek() (t T) {
	if s.l.Back() != nil {
		return s.l.Back().Value.(T)
	}
	return
}

func (s *stack[T]) Push(v T) {
	s.l.PushBack(v)
}

func (s *stack[T]) Pop() (t T) {
	if s.l.Back() != nil {
		t = s.l.Remove(s.l.Back()).(T)
		return
	}
	return
}

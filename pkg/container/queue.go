package container

import (
	"container/list"
	"fmt"
	"strings"
)

type Queue[T any] interface {
	Len() int
	IsEmpty() bool
	Peek() T
	Push(v T)
	Pop() T
}

type queue[T any] struct {
	l *list.List
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{l: list.New()}
}

func (q *queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *queue[T]) Peek() (t T) {
	if q.l.Front() != nil {
		t = q.l.Front().Value.(T)
		return
	}
	return
}
func (q *queue[T]) Len() int {
	return q.l.Len()
}

func (q *queue[T]) Push(v T) {
	q.l.PushBack(v)
}

func (q *queue[T]) Pop() (t T) {
	if q.l.Front() != nil {
		t = q.l.Remove(q.l.Front()).(T)
		return
	}
	return
}

func (q *queue[T]) String() string {
	es := make([]string, 0, q.Len())
	e := q.l.Front()
	for ; e != nil; e = e.Next() {
		es = append(es, fmt.Sprintf("%s", e.Value))
	}
	return strings.Join(es, " ->")
}

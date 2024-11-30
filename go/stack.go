package main

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	lock sync.Mutex
	s    []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{lock: sync.Mutex{}, s: make([]T, 0)}
}

func (s *Stack[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, v)
}

func (s *Stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var zero T
	l := len(s.s)
	if l == 0 {
		return zero, errors.New("empty stack")
	}
	ele := s.s[l-1]
	s.s = s.s[:l-1]
	return ele, nil

}

func (s *Stack[T]) IsEmpty() bool {
	l := len(s.s)
	if l == 0 {
		return true
	}
	return false
}

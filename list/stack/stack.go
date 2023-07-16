package main

import (
	"container/list"
)

type Stack struct {
	data *list.List
}

func (s *Stack) Push(v interface{}) {
	s.data.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.data.Remove(s.data.Back())
}

func (s *Stack) Peek() interface{} {
	return s.data.Back()
}

func (s *Stack) GetSize() int {
	return s.data.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.GetSize() == 0
}

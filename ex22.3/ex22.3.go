package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{ list.New() }
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	pop := s.v.Back()
	if pop != nil {
		return s.v.Remove(pop)
	}
	return nil
}

func (s *Stack) isEmpty() bool {
	return s.v.Len() == 0
}

func main() {
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	for !stack.isEmpty() {
		fmt.Println(stack.Pop())
	}
}
package main

import "fmt"

type LinkedStack struct {
	top *StackNode
}

type StackNode struct {
	data interface{}
	last *StackNode
}

func NewArrayStack() *LinkedStack {
	return &LinkedStack{}
}

func (s *LinkedStack) Push(i interface{}) {
	cur := &StackNode{
		data: i,
	}

	if s.top == nil {
		s.top = cur
	} else {
		cur.last = s.top
		s.top = cur
	}
}

func (s *LinkedStack) Pop() interface{} {
	if s.top == nil {
		panic("stack is empty")
	}
	item := s.top.data
	s.top = s.top.last

	return item
}

func main() {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push("testHello")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

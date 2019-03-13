package main

import "fmt"

type ArrayStack struct {
	top      int
	data     []interface{}
	capacity int
}

func NewArrayStack(capacity int) *ArrayStack {
	data := make([]interface{}, capacity, capacity)
	return &ArrayStack{
		top:      -1,
		data:     data,
		capacity: capacity,
	}
}

func (s *ArrayStack) Push(i interface{}) {
	s.top++
	if s.top == s.capacity {
		panic("stack is full")
	}

	s.data[s.top] = i
}

func (s *ArrayStack) Pop() interface{} {
	if s.top < 0 {
		panic("stack is empty")
	}

	item := s.data[s.top]
	s.data[s.top] = nil
	s.top--
	return item
}

func main() {
	s := NewArrayStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

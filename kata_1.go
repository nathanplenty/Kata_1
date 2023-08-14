package main

import "fmt"

// struct -> table with only one column, each can use different types
type Stack struct {
	Value []interface{}
}

// function with a pointer return (allows to ceate more structs)
func NewStack() *Stack {
	stack := Stack{make([]interface{}, 0)}
	return &stack
}

// function adds a new input at the top of the stack (or bottom of the table) LIFO-rule
func (s *Stack) Push(v interface{}) {
	s.Value = append(s.Value, v)
	fmt.Println("[", v, "] pushed!")
}

// function returns the latest index's element and delets the latest index afterwards
func (s *Stack) Pop() (int, bool) {
	b := false
	i := 0
	if s.IsEmpty() == false {
		s.Peek()
		i = len(s.Value) - 1
		s.Value = s.Value[:i]
		fmt.Println("Last Index poped!")
		b = true
	}
	return i, b
}

// function returns the latest index's element
func (s *Stack) Peek() (int, bool) {
	b := false
	i := 0
	if s.IsEmpty() == false {
		i = len(s.Value) - 1
		v := s.Value[i]
		fmt.Println("Index:", i, ", Value:", v)
		b = true
	}
	return i, b
}

// function indicates whether the input is empty
func (s *Stack) IsEmpty() bool {
	b := false
	if len(s.Value) == 0 {
		b = true
	}
	return b
}

func main() {
	stack := NewStack()
	stack.Push(5)
	stack.Push(10)
	stack.Push("text")
	stack.Peek()
	stack.Pop()
	isEmpty := stack.IsEmpty()
	fmt.Println("Is Stack empty?", isEmpty)
}

package main

import "fmt"

// struct -> table with only one column, each can use different types
type Stack struct {
	Values []interface{}
}

// function with a pointer return (allows to ceate more structs)
func NewStack() *Stack {
	stack := Stack{}
	return &stack
}

// function adds a new input at the top of the stack (or bottom of the table) LIFO-rule
func (s *Stack) Push(v interface{}) {
	s.Values = append(s.Values, v)
}

// function returns the latest index's element and delets the latest index afterwards
func (s *Stack) Pop() int {
	i := 0
	if s.IsEmpty() == false {
		i = len(s.Values) - 1
		s.Values = s.Values[:i]
	}
	return i
}

// function returns the latest index's element
func (s *Stack) Peek() int {
	i := 0
	if s.IsEmpty() == false {
		i = len(s.Values) - 1
	}
	return i
}

// function indicates whether the input is empty
func (s *Stack) IsEmpty() bool {
	b := false
	if len(s.Values) == 0 {
		b = true
	}
	return b
}

func main() {
	stack := NewStack()
	stack.Push(5)
	stack.Push('2')
	stack.Push("text")
	stack.Peek()
	stack.Pop()
	isEmpty := stack.IsEmpty()
	fmt.Println("Is Stack empty?", isEmpty)
}

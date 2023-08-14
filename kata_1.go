package main

import "fmt"

// struct -> table with only one column, each can use different types
type Stack struct {
	Stack []int
}

// function with a pointer return (allows to ceate more structs)
func NewStack() *Stack {
	stack := Stack{make([]int, 0)}
	return &stack
}

// function adds a new input at the top of the stack (or bottom of the table) LIFO-rule
func (s *Stack) Push(val int) {
	s.Stack = append(s.Stack, val)
}

// function returns the latest index's element and delets the latest index afterwards
func (s *Stack) Pop() (int, bool) {
	b := false
	v := 0
	if s.IsEmpty() == false {
		i := len(s.Stack) - 1
		v = s.Stack[i]
		s.Stack = s.Stack[:i]
		b = true
	}
	return v, b
}

// function returns the latest index's element
func (s *Stack) Peek() (int, bool) {
	b := false
	v := 0
	if s.IsEmpty() == false {
		i := len(s.Stack) - 1
		v = s.Stack[i]
		b = true
	}
	return v, b
}

// function indicates whether the input is empty
func (s *Stack) IsEmpty() bool {
	b := false
	if len(s.Stack) == 0 {
		b = true
	}
	return b
}

func main() {
	stack := NewStack()
	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	val, _ := stack.Peek()
	fmt.Println(val)
	val, _ = stack.Pop()
	fmt.Println(val)
	isEmpty := stack.IsEmpty()
	fmt.Println(isEmpty)
}

package main

import "fmt"

// struct -> table with only one column, each can use different types
type Stack struct {
	Values []interface{}
}

// function with a pointer return (allows to ceate more structs)
func CreateNewStack() *Stack {
	stack := Stack{}
	return &stack
}

// function adds a new input at the top of the stack (or bottom of the table) LIFO-rule
func (s *Stack) PushNewValue(v interface{}) {
	s.Values = append(s.Values, v)
}

// function returns the latest index's element and delets the latest index afterwards
func (s *Stack) PopLatestValue() interface{} {
	if s.IsStackEmpty() == true {
		return "nil"
	}
	// <!> Full slice expressions
	popepValue := s.PeekLatestValue()
	s.Values = s.Values[:len(s.Values)-1]
	return popepValue
}

// function returns the latest index's element
// del index -> use interface
func (s *Stack) PeekLatestValue() interface{} {
	if s.IsStackEmpty() == true {
		return "nil"
	}
	return s.Values[len(s.Values)-1]
}

// function indicates whether the input is empty
func (s *Stack) IsStackEmpty() bool {
	// <!> If possible to One-Liner
	if len(s.Values) == 0 {
		return true
	}
	return false
}

func main() {
	stack := CreateNewStack()
	stack.PushNewValue(5)
	stack.PushNewValue('2')
	stack.PushNewValue("text")
	fmt.Println("Peek:", stack.PeekLatestValue())
	fmt.Println("Pop:", stack.PopLatestValue())
	isEmpty := stack.IsStackEmpty()
	fmt.Println("Is Stack empty?", isEmpty)
}

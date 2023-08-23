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
func (stack *Stack) PushNewValue(v interface{}) {
	stack.Values = append(stack.Values, v)
}

// function returns the latest index's element and delets the latest index afterwards
func (stack *Stack) PopLatestValue() interface{} {
	if stack.IsStackEmpty() == true {
		return "nil"
	}
	// <!> Full slice expressions
	popepValue := stack.PeekLatestValue()
	stack.Values = stack.Values[:len(stack.Values)-1]
	return popepValue
}

// function returns the latest index's element
// del index -> use interface
func (stack *Stack) PeekLatestValue() interface{} {
	if stack.IsStackEmpty() == true {
		return "nil"
	}
	return stack.Values[len(stack.Values)-1]
}

// function indicates whether the input is empty
func (stack *Stack) IsStackEmpty() bool {
	// <!> If possible to One-Liner
	if len(stack.Values) == 0 {
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
	fmt.Println("Is Stack empty?", stack.IsStackEmpty())
}

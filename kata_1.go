package main

import "fmt"

type Stack struct {
	Values []interface{}
}

func CreateNewStack() *Stack {
	stack := Stack{}
	return &stack
}

func (stack *Stack) PushNewValue(value interface{}) {
	stack.Values = append(stack.Values, value)
}

func (stack *Stack) PopLatestValue() interface{} {
	if stack.IsStackEmpty() == true {
		return nil
	}
	popepValue := stack.PeekLatestValue()
	lengthOfValue := len(stack.Values) - 1
	stack.Values = stack.Values[:lengthOfValue:lengthOfValue]
	return popepValue
}

func (stack *Stack) PeekLatestValue() interface{} {
	if stack.IsStackEmpty() == true {
		return nil
	}
	return stack.Values[len(stack.Values)-1]
}

func (stack *Stack) IsStackEmpty() bool {
	return stack.Values == nil
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

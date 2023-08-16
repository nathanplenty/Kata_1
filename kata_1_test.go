package main

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	b := false
	stack := NewStack()
	b = stack.IsEmpty()
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(stack)
}

func TestPushPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	i, b := stack.Peek()
	v := stack.Value[i]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(stack)
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	stack.Push(10)
	i, b := stack.Pop()
	v := stack.Value[i-1]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(stack)
}

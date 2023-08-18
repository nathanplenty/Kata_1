package main

import (
	// "fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	b := false
	stack := NewStack()
	b = stack.IsEmpty()
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
}

func TestPushPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	i := stack.Peek()
	v := stack.Values[i]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	stack.Push(10)
	i := stack.Pop()
	v := stack.Values[i-1]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}

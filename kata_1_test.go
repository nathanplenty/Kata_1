package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := CreateNewStack()
	b := stack.IsStackEmpty()
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
}

func TestPushPeek(t *testing.T) {
	stack := CreateNewStack()
	stack.PushNewValue(5)
	v := stack.PeekLatestValue()
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}

func TestPop(t *testing.T) {
	stack := CreateNewStack()
	stack.PushNewValue(5)
	stack.PushNewValue(10)
	v := stack.PopLatestValue()
	if v != 10 {
		t.Errorf("v = %v; want 5", v)
	}
}

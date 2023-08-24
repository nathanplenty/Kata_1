package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := CreateNewStack()
	boolean := stack.IsStackEmpty()
	if boolean != true {
		t.Errorf("boolean = %v; want true", boolean)
	}
}

func TestPushPeek(t *testing.T) {
	stack := CreateNewStack()
	stack.PushNewValue(5)
	value := stack.PeekLatestValue()
	if value != 5 {
		t.Errorf("value = %v; want 5", value)
	}
}

func TestPop(t *testing.T) {
	stack := CreateNewStack()
	stack.PushNewValue(5)
	stack.PushNewValue(10)
	value := stack.PopLatestValue()
	if value != 10 {
		t.Errorf("value = %v; want 5", value)
	}
}

func TestForNil(t *testing.T) {
	stack := CreateNewStack()
	value := stack.PeekLatestValue()
	if value != nil {
		t.Errorf("value = %v; want <nil>", value)
	}
}

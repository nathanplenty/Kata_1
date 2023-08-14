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

package main

import (
	"testing"
	"fmt"
)

func newStack() *Stack {
	var s *Stack
	s = s.New(10)
	return s
}

func TestStack_Push(t *testing.T) {
	s := newStack()
	s.Push(10)
	x := s.Pop()
	fmt.Print(x)
}

package main

type Stack struct {
	Size    int
	Top     int
	Element []int
	MaxSize int
}

func (s *Stack) New(max int) *Stack {
	if s == nil {
		s = new(Stack)
	}
	s.Size = 0
	s.MaxSize = max
	//s.Top=nil
	return s
}

func (s *Stack) Pop() int {
	if s.Size == 0 {
		return 0
	}
	item := s.Element[s.Size-1]
	s.Size -= 1
	s.Element = s.Element[:s.Size]
	if s.Size > 0 {
		s.Top = s.Element[s.Size-1]
	} else {
		s.Top = 0
	}
	return item
}

func (s *Stack) Push(x int) {
	s.Element = append(s.Element, x)
	s.Size++
	s.Top = x
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) IsFull() bool {
	return s.Size == s.MaxSize
}

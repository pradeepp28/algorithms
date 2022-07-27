package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(e int) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	l := len(s.elements)
	top := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return top, true
}

func main() {
	s := &Stack{}
	fmt.Println(s.Pop())
	s.Push(3)
	s.Push(2)
	s.Push(1)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

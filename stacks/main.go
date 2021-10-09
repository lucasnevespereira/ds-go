package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Elements []int
}

// pushs new element in top of stack
func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

// retrieves nelement from top of stack
func (s *Stack) Pop() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("stack is empty")
	}

	lastElementIdx := len(s.Elements) - 1
	lastElement := s.Elements[lastElementIdx]
	s.Elements = s.Elements[:lastElementIdx]

	return lastElement, nil
}

// returns top element of stack without updating it
func (s *Stack) Peek() (int, error) {
	if s.Elements[0] == 0 {
		return 0, errors.New("stack is empty")
	}
	lastIdx := len(s.Elements) - 1
	return s.Elements[lastIdx], nil

}

// returns size of stack
func (s *Stack) Length() int {
	return len(s.Elements)
}

// checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func main() {
	// A stack is similar to a queue but is only open in one end
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack) // {[1 2 3]}
	fmt.Println("------")

	peek1, _ := stack.Peek()
	fmt.Println(peek1) // 3

	fmt.Println(stack.Length()) // 3

	fmt.Println("------")
	elem1, _ := stack.Pop()
	fmt.Println(elem1) // 3
	elem2, _ := stack.Pop()
	fmt.Println(elem2) // 2
	elem3, _ := stack.Pop()
	fmt.Println(elem3) // 1

	fmt.Println("------")
	fmt.Println(stack.IsEmpty()) // true
}

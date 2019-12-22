package gods

import (
	"fmt"
)

// Stack using linked list
type Stack struct {
	Top *Node
	Length int
}

// Node that has link of previous node
type Node struct {
	Value interface{}
	Prev *Node
}

// NewStack : constructor for new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
	
}

// Add addes element on stacks
func (s *Stack)Add(items ...interface{}) {
	for _, i := range items {
		n := Node{Value: i, Prev: s.Top}
		s.Top = &n
		s.Length++
	}
}

// Pop removes top element from stack
func (s *Stack)Pop() interface{} {
	
	if s.Length == 0 {
		return nil
	}
	
	t :=  s.Top
	s.Top = t.Prev
	s.Length--
	return t.Value
}

// Peek return top element from Stack
func (s *Stack) Peek() interface{} {
	if s.Length == 0 {
		return nil
	}
	return s.Top.Value
}

// Len return lenght of stack
func (s *Stack) Len() int {
	return s.Length
}

// Push inserts a element from stack
func (s *Stack) Push(item interface{}) {
	n := &Node{Value: item, Prev: s.Top}
	s.Top = n
	s.Length++

}


// Print prints all the elements from stack
func (s *Stack) Print() {
	top := s.Top
	for top != nil {
		fmt.Println(top.Value)
		top = top.Prev
	}
}

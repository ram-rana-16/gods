package gods

import (
	"testing"
	"fmt"
)

func TestNewSatck(t *testing.T) {
	s := NewStack()
	s.Add(1, 2, 3, 4, 5, 6, 7)
	s.Print()
	fmt.Println("Adding a new element 9")
	s.Add(9)
	s.Print()
	fmt.Println("printing peek = ", s.Peek())
	fmt.Println("removing top")
	s.Pop()
	s.Pop()
	s.Print()
	fmt.Println("lenght of stakc", s.Len())
}
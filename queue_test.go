package gods

import (
	"testing"
	"fmt"
)

func TestNewQ(t *testing.T) {
	q := NewQ()
	fmt.Printf("Dequeue from empty Q %+v\n", q.Dequeue())
	q.Enqueue(1, 2, 3, 4, 5, 6)
	q.Print()
	fmt.Printf("Dequeue from Queue :%+v\n", q.Dequeue())
	fmt.Println("Enqueue to Queue 77")
	q.Enqueue(77)
	q.Print()
	fmt.Printf("Dequeue from Queue :%+v\n", q.Dequeue())
	q.Print()

	fmt.Println("length of the queueu", q.Len())
}
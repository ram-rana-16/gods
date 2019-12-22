package gods

import(
	"fmt"
)

// Queue : a Q datat structure 
type Queue struct {
	Items []interface{}
}

// NewQ Return a queue 
func NewQ() *Queue {
	c := Queue{}
	return &c
}


// Enqueue addes element into the Q
func (q *Queue) Enqueue(items ...interface{}) {
	for _, item := range items {
		q.Items = append(q.Items, item)
	}
}

// Dequeue removes element from Q
func (q *Queue) Dequeue() interface{} {
	if len(q.Items) > 0 {
		fmt.Println("insede")
		x := q.Items[0]
		q.Items = q.Items[1:]
		return x
	}
	return nil
}

// Len returns lenght of the Queue 
func (q *Queue) Len() int {
	return len(q.Items)
}

// Print : prints all the elements from queue 
func (q *Queue) Print() {

	for _, val := range q.Items {
		fmt.Println(val)
	}
}

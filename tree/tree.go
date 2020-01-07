package tree

import (
	"fmt"
)

// Tree is a binary tree with integer value
type Tree struct {
	Value int
	L *Tree
	R *Tree
}

// Print prints a tree in Inorder
func (t *Tree)Print() {
	if t == nil {
		return
	}
	t.L.Print()
	fmt.Println(t.Value)
	t.R.Print()

}
// Insert inserts new node to a tree 
func (t *Tree)Insert(v int) *Tree{
	if t == nil {
		t = &Tree{Value: v} 
	}else if t.Value < v {
		t.R = t.R.Insert(v)
	}else {
		t.L = t.L.Insert(v)
	}
	return t
}

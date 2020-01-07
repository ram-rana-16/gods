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

// Delete deletes node from tree 
func (t *Tree)Delete(v int) *Tree{
	if t == nil {
		return t
	}
	fmt.Println("node ", t)
	if t.Value > v {
		t.L = t.L.Delete(v)
	}else if t.Value < v {
		t.R = t.R.Delete(v)
	}else if t.L == nil && t.R == nil {
		fmt.Println("case 1")
		t = nil
	}else if t.L == nil {  // case 2: one child
		fmt.Println("case 2")
		t = t.R
	}else if t.R == nil {
		t = t.L
	}else { // case 3: 3 child 
		minNode := findMin(t.R)
		t.Value = minNode.Value
		t.R = t.R.Delete(minNode.Value)
	}
	fmt.Println("returned t", t)
	return t

}


func findMin(t *Tree) *Tree {
	for t.L != nil {
		t = t.L
	}
	return t
}

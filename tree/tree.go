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
func (t *Tree)Print(res []int) []int{
	if t == nil {
		return res
	}
	res = t.L.Print(res)
	fmt.Println(t.Value)
	res = append(res, t.Value)
	res = t.R.Print(res)
	return res
}
// Insert inserts new node to a tree 
func (t *Tree)Insert(v int) *Tree{
	if t == nil {
	t = &Tree{Value: v} 
	}else if t.Value <= v {
		t.R = t.R.Insert(v)
	}else {
		t.L = t.L.Insert(v)
	}
	return t
}

// Delete deletes node from tree 
func (tree *Tree)Delete(value int) *Tree{
	if tree == nil {
		return tree
	}
	if tree.Value < value {
		tree.R = tree.R.Delete(value)
	}else if tree.Value > value {
		tree.L = tree.L.Delete(value)
	}else if tree.L == nil && tree.R == nil {
		tree = nil
	}else if tree.L == nil {
		tree = tree.R
	}else if tree.R == nil {
		tree = tree.L
	}else {
		minNode := findMin(tree.R)
		tree.Value = minNode.Value
		tree.R = tree.R.Delete(minNode.Value)
	}
	return tree

}


func findMin(t *Tree) *Tree {
	for t.L != nil {
		t = t.L
	}
	return t
}





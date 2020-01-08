package tree

import(
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	tree := &Tree{Value:6}
	tree = tree.Insert(8).Insert(9).Insert(2)
	tree.Print(nil)
}

func TestDelete(t *testing.T) {
	tree := &Tree{Value:10}
	tree.Insert(5).Insert(7).Insert(2)
	x := tree.Print(nil)
	fmt.Println("x list", x)
	fmt.Println("after deletion")
	tree = tree.Delete(10)
	tree.Print(nil)
}


func NewBST(root int, values ...int) *Tree {
	tree := &Tree{Value: root}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}


func TestNewBST(t *testing.T) {
	var test1 = NewBST(10, 5, 15, 22, 17, 34, 7, 2, 5, 1, 35, 27, 16, 30)
	test1.Print(nil)
	test1.Delete(22).Delete(17)
	fmt.Println("after deleetion")
	res := test1.Print([]int{})
	fmt.Println("res", res)
}
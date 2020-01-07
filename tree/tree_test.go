package tree

import(
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	tree := &Tree{Value:6}
	tree = tree.Insert(8).Insert(9).Insert(2)
	tree.Print()
}

func TestDelete(t *testing.T) {
	tree := &Tree{Value:12}
	tree.Insert(5).Insert(17).Insert(3).Insert(7).Insert(13).
	Insert(1).Insert(9).Insert(14).Insert(20).Insert(8).Insert(18)
	tree.Print()
	fmt.Println("after deletion")
	tree.Delete(17)
	tree.Print()
}
package tree

import(
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &Tree{Value:6}
	tree = tree.Insert(8).Insert(9).Insert(2)
	tree.Print()
}
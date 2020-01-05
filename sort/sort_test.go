package sort

import(
	"testing"
	"fmt"
)


func TestSelection(t *testing.T)  {
	arr := []int{64, 25, 12, 22, 11}
	out := SelectionsSort(arr)
	fmt.Println(out)
}

func TestInsertion(t *testing.T)  {
	arr := []int{8, -6, 7, 10, 8, -1, 6, 2, 4, -5, 1, 10, 8, -10, -9, -10, 8, 9, -2, 7, -2, 4}
	out := Insertion(arr)
	fmt.Println(out)
}
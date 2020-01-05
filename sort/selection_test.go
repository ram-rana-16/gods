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
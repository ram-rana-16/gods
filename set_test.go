
package gods
import 
(
	"testing"
	"fmt"
)
func TestNewSet(t *testing.T) {
	s := NewSet(1, 3 ,2, 3, 4, 5)
	s.Print()
}

func TestAddNew(t *testing.T) {
	s := NewSet(1, 2, 3, 4)
	s.Print()
	s.Add(8)
	fmt.Println("after adding new element")
	s.Print()
	s.Remove(8)
	fmt.Println("after removing the element 8")
	s.Print()
	s.Remove(8)
	fmt.Println("after removing the element 8")
	s.Print()
}
package gods

import "fmt"

// Set is implemented using builtin map 
type Set struct {
	Items map[interface{}]struct{}
}

// Print prints all the elenets of the set
func (s *Set) Print() {
	for k := range s.Items {
		fmt.Println(k)
	}
}

// Add adds a single element into the set
func (s *Set)Add(item interface{})  {
s.Items[item] = struct{}{}
}

// Remove Removes an item from set if it is present in the set else nothing
func (s *Set)Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.Items, item)
	}
}

// AddItems add list of items into the set
func (s *Set) AddItems(items ...interface{}) {
	for _, i := range items {
		s.Items[i] = struct{}{}
	}
}

// NewSet : constructor for set
func NewSet(items ...interface{}) *Set {
	var s = Set{}
	s.Items = make(map[interface{}]struct{})
	for _, item := range items {
		s.Items[item] = struct{}{}
	}

	return &s
}
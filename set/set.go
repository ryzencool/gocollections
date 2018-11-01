package set

import (
	"log"
)

// Set set implementation by map
type Set struct {
	set map[interface{}]struct{}
}

// NewSet init Set
func NewSet() *Set {
	return &Set{
		set: make(map[interface{}]struct{}),
	}
}

// Add add element into set
func (s *Set) Add(element interface{}) {
	s.set[element] = struct{}{}
}

// AddAll add elements into set
func (s *Set) AddAll(elements ...interface{}) {
	for _, e := range elements {
		s.set[e] = struct{}{}
	}
}

// AddSlice add slice into set
func (s *Set) AddSlice(elements []interface{}) {
	for _, e := range elements {
		s.set[e] = struct{}{}
	}
}

// Elements get elements from set
func (s *Set) Elements() []interface{} {
	rset := make([]interface{}, 0)
	for k := range s.set {
		rset = append(rset, k)
	}
	return rset
}

// Remove remove some element from set
func (s *Set) Remove(element interface{}) {
	log.Printf("value is : %v", s)
	delete(s.set, element)
}

// Contains if set contains some value
func (s *Set) Contains(element interface{}) bool {
	_, ok := s.set[element]
	return ok
}

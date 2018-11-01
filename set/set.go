package set

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
	delete(s.set, element)
}

// RemoveAll remove some elements from set
func (s *Set) RemoveAll(elements ...interface{}) {
	for ele := range elements {
		delete(s.set, ele)
	}
}

// Clear all elements in set
func (s *Set) Clear() {
	for ele := range s.set {
		delete(s.set, ele)
	}
}

// Contains if set contains some value
func (s *Set) Contains(element interface{}) bool {
	_, ok := s.set[element]
	return ok
}

// RetainAll get Intersection
func (s *Set) RetainAll(s1 *Set) {
	for _, v := range s.Elements() {
		if !s1.Contains(v) {
			s.Remove(v)
		}
	}
}

// Len get length of set
func (s *Set) Len() int {
	return len(s.set)
}

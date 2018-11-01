package set

import (
	"sync"
)

// ConcurrentSet concurrent solution of set
type ConcurrentSet struct {
	set sync.Map
}

// NewConcurrentSet init set
func NewConcurrentSet() sync.Map {
	var s sync.Map
	return s
}

// Len get length of concurrent set
func (c *ConcurrentSet) Len() int {
	l := 0
	c.set.Range(func(key, value interface{}) bool {
		l++
		return true
	})
	return l
}

// Add add element
func (c *ConcurrentSet) Add(element interface{}) {
	c.set.Store(element, struct{}{})
}

// AddAll add some elements
func (c *ConcurrentSet) AddAll(elements ...interface{}) {
	for _, v := range elements {
		c.set.Store(v, struct{}{})
	}
}

// Remove remove element
func (c *ConcurrentSet) Remove(element interface{}) {
	c.set.Delete(element)
}

// RemoveAll remove some elements
func (c *ConcurrentSet) RemoveAll(elements ...interface{}) {
	for _, v := range elements {
		c.set.Delete(v)
	}
}

// Elements get elements of set
func (c *ConcurrentSet) Elements() []interface{} {
	var list []interface{}
	c.set.Range(func(key, value interface{}) bool {
		list = append(list, key)
		return true
	})
	return list
}

// Clear clear all
func (c *ConcurrentSet) Clear() {
	c.set.Range(func(k, v interface{}) bool {
		c.set.Delete(k)
		return true
	})
}

// Contains contain
func (c *ConcurrentSet) Contains(element interface{}) bool {
	_, ok := c.set.Load(element)
	return ok
}

// RetainAll retain all
func (c *ConcurrentSet) RetainAll(c1 *ConcurrentSet) {
	c.set.Range(func(k, v interface{}) bool {
		if !c1.Contains(k) {
			c.Remove(k)
		}
		return true
	})
}

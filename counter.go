package collections

import (
	"sort"
)

// Counter like python's Counter class
type Counter struct {
	counter map[interface{}]int
}

// NewCounter initialize Counter
func NewCounter() *Counter {
	return &Counter{
		counter: make(map[interface{}]int),
	}
}

// UpdateString insert elements into map, if string please convert it to rune or []byte
func (c *Counter) UpdateString(elems string) {
	for _, s := range elems {
		if _, ok := c.counter[s]; ok {
			c.counter[s]++
		} else {
			c.counter[s] = 1
		}
	}
}

// UpdateStuct update struct
func (c *Counter) UpdateStuct(elem interface{}) {
	if _, ok := c.counter[elem]; ok {
		c.counter[elem]++
	} else {
		c.counter[elem] = 1
	}
}

// MostCommon get most increase, n < len(c.counter)
func (c *Counter) MostCommon(n int) PairList {
	pl := make(PairList, len(c.counter))
	i := 0
	for k, v := range c.counter {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl[:n]
}

// Pair to present map
type Pair struct {
	Key   interface{}
	Value int
}

// PairList alias of Pair slice
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

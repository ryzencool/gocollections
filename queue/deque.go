package queue

import "log"

// Deque deque
type Deque struct {
	queue []interface{}
	head  int
	tail  int
}

// NewDeque new deque
func NewDeque() *Deque {
	return &Deque{
		queue: make([]interface{}, 8),
	}
}

func (d *Deque) calculateSize(num int) int {
	initCap := 8
	if initCap < num {
		initCap = num
		initCap |= initCap >> 1
		initCap |= initCap >> 2
		initCap |= initCap >> 4
		initCap |= initCap >> 8
		initCap |= initCap >> 16
		initCap++
	}
	if initCap < 0 {
		initCap >>= 1
	}
	return initCap
}

func (d *Deque) allocateElements(num int) {
	d.queue = make([]interface{}, d.calculateSize(num))
}

func (d *Deque) doubleCapacity() {
	n := len(d.queue)
	p := d.head
	tmp := make([]interface{}, n<<1)
	copy(tmp[len(tmp)-n+p:], d.queue[p:])

	copy(tmp[:p], d.queue[:p])
	for i, v := range tmp {
		log.Printf("tmp [%v] is [%v]", i, v)
	}
	d.queue = tmp
}

// AppendLeft append left
func (d *Deque) AppendLeft(element interface{}) {
	d.head = (d.head - 1) & (len(d.queue) - 1)
	d.queue[d.head] = element
	if d.head == d.tail {
		d.doubleCapacity()
	}
}

// AppendRight append right
func (d *Deque) AppendRight(element interface{}) {
	d.queue[d.tail] = element
	d.tail = (d.tail + 1) & (len(d.queue) - 1)
	if d.tail == d.head {
		d.doubleCapacity()
	}
}

// RemoveLeft remove left
func (d *Deque) RemoveLeft() interface{} {
	p := d.head
	res := d.queue[p]
	if d.queue[p] == nil {
		return nil
	}
	d.queue[p] = nil
	d.head = (d.head + 1) & (len(d.queue) - 1)
	return res
}

// RemoveLast remove right
func (d *Deque) RemoveLast() interface{} {
	p := d.tail
	res := d.queue[p]
	if d.queue[p] == nil {
		return nil
	}
	d.queue[p] = nil
	d.tail = (d.tail - 1) & (len(d.queue) - 1)
	return res
}


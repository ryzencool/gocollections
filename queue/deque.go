package queue

// Deque deque
type Deque struct {
	queue []interface{}
	head  int
	tail  int
}

// NewDeque new deque
func (d *Deque) NewDeque(cap int) *Deque {
	return &Deque{
		queue: make([]interface{}, cap),
	}
}

func (d *Deque) AppendLeft(element interface{}) {
	d.head = (d.head - 1) % (len(d.queue) - 1)
	d.queue[d.head] = element
}

func (d *Deque) AppendRight(element interface{}) {
	d.tail = (d.tail - 1) % (len(d.queue) - 1)
	d.queue[d.tail] = element
}

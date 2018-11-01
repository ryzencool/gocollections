package queue

// Deque deque
type Deque struct {
	queue []interface{}
	head  int
	tail  int
}

// NewDeque new deque
func NewDeque(cap int) *Deque {
	return &Deque{
		queue: make([]interface{}, cap),
	}
}

func (d *Deque) calculateSize() {
	
}

func (d *Deque) doubleCapacity() {

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

package queue

import (
	"log"
	"testing"
	
)

func TestQueue(t *testing.T) {
	q := NewDeque()

	q.AppendLeft(1)
	q.AppendLeft(1)
	q.AppendLeft(1)
	q.AppendRight(2)
	log.Printf("value: %v", q.Len())
	for _, v := range q.Elements() {
		log.Printf("value:%v", v)
	}
}

func BenchmarkQueue(t *testing.B) {
	t.StopTimer()
	q := NewDeque()
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		q.AppendLeft(1)
	}
}

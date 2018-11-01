package queue

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewDeque(16)
	q.AppendLeft("java")
	q.AppendRight("python")
	log.Printf("value:%v", q)
}

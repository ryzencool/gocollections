package collections

import (
	"log"
	"testing"
	"time"
)

func TestMongoCon(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- opt()
	}()
	time.Sleep(2 * time.Second)
	done := make(chan interface{})
	defer close(done)
	select {
	case <-done:
		return
	case <-ch:
		log.Println(<-ch)
	}
	log.Println("结束")
}

func opt() string {
	log.Println("进入请求")
	time.Sleep(3 * time.Second)
	log.Println("走出请求")
	return "name"
}

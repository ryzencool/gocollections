package counter

import (
	"log"
	"testing"
)

// "log"
// "testing"

type User struct {
	name string
}

func TestString(t *testing.T) {
	c := NewCounter()
	c.UpdateString("abbcccddddeeeeeee")
	cc := c.MostCommon(2)
	for _, v := range cc {
		log.Println(string(rune(v.Key.(int32))), v.Value)
	}
	for _, v := range c.Elements() {
		log.Println(string(rune(v.Key.(int32))), v.Value)
	}
}

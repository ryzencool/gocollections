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

// func TestBase(t *testing.T) {
// 	c := NewCounter()
// 	// u1 := User{name: "string"}
// 	// c.Update(u1)
// 	// u2 := User{name: "name"}
// 	// c.Update(u2)
// 	// u3 := User{name: "string"}
// 	// c.Update(u3)
// 	bb := []User{User{name: "string"}, User{name: "name"}, User{name: "string"}}
// 	c.InitElems(bb)
// 	for k, v := range c.counter {
// 		log.Println(k, v)
// 	}

// }
func TestString(t *testing.T) {
	c := NewCounter()
	c.UpdateString("abbcccddddeeeeeee")
	cc := c.MostCommon(10)
	for _, v := range cc {
		log.Println(string(rune(v.Key.(int32))))
		log.Println(v.Value)
	}
}

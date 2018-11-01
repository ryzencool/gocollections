package set

import (
	"log"
	"testing"
)

func data() (Interface, Interface) {
	s := NewSet()
	s.Add("java")
	s.Add("php")
	s.AddAll("c++", "golang")
	s1 := NewSet()
	s1.Add("java")
	return s, s1

}

func TestRetainAll(t *testing.T) {
	s1, s2 := data()
	s1.RetainAll(s2)
	for _, v := range s1.Elements() {
		log.Printf("value:%v", v)
	}

}

func TestMapSet(t *testing.T) {
	m := make(map[interface{}]string)
	del := func(v interface{}) {
		delete(m, v)
	}
	m["java"] = "name"
	m["php"] = "spring"
	del("java")
	for k := range m {
		log.Printf("the m is:%v", k)
	}
}

func TestStruct(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	u1 := user{name: "zmy", age: 11}
	u2 := user{name: "zmy", age: 11}
	s := NewSet()
	s.AddAll(u1, u2)
	log.Printf("value is : %v", s)
}

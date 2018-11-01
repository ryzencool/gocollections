package set

import (
	"log"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet()
	s.Add("java")
	s.Add("php")
	s.AddAll("c++", "golang")
	e := s.Elements()
	s.Remove("java")
	s.Remove("golang")
	for _, v := range e {
		log.Printf("the elements of set is : %v", v)
	}

	ok := s.Contains("java")
	if ok {
		log.Printf("set contains java")
	} else {
		log.Printf("set not contains java")
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

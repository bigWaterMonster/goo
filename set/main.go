package main

import (
	"fmt"
	"unsafe"
)

var Exists = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

func New(items ...interface{}) *Set {
	// 获取Set的地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}



func main()  {
	fmt.Println(unsafe.Sizeof(struct {

	}{}))
}

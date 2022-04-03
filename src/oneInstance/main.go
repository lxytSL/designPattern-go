package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	dowork()
}

type singleton struct{}

func (s *singleton) dowork() {
	fmt.Println("dowork")
}

var (
	one    sync.Once
	sigton *singleton
)

// use sync.One achieve
func getInstance() Singleton {
	one.Do(
		func() {
			sigton = &singleton{}
		},
	)
	return sigton
}

func getInstance2() Singleton {
	if sigton == nil {
		sigton = &singleton{}
	}
	return sigton
}

func main() {
	s1 := getInstance2()
	fmt.Printf("%p\n", s1)
	s2 := getInstance2()
	fmt.Printf("%p\n", s2)
}

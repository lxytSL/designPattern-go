package main

import (
	"fmt"
)

type Factory interface {
	doWork_1()
	doWork_2()
}

type Tools_1 struct{}

func (t *Tools_1) doWork_1() { fmt.Println("tools_1 work1") }
func (t *Tools_1) doWork_2() { fmt.Println("tools_1 work2") }

type Tools_2 struct{}

func (t *Tools_2) doWork_1() { fmt.Println("tools_2 work1") }
func (t *Tools_2) doWork_2() { fmt.Println("tools_2 work2") }

func newFactory(tool string) Factory {
	switch tool {
	case "Tools_1":
		return &Tools_1{}
	case "Tools_2":
		return &Tools_2{}
	default:
		return nil
	}
}

func main() {
	tools_1 := newFactory("Tools_1")
	tools_1.doWork_1()
	tools_1.doWork_2()
	tools_2 := newFactory("Tools_2")
	tools_2.doWork_1()
	tools_2.doWork_2()
}

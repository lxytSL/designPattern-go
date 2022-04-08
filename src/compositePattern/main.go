package main

import "fmt"

type Component interface {
	Add(Component)
	Remove(Component)
	DisPlay(int)
}

type leaf struct {
	name string
}

func (leaf) Add(c Component) {
	panic("lead not component add compacity")
}

func (leaf) Remove(c Component) {
	panic("lead not component remove compacity")
}

func (l leaf) DisPlay(deepth int) {
	for i := 0; i < deepth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("leaf %s execute!\n", l.name)
}

type Composite struct {
	name string
	c    []Component
}

func (cps *Composite) Add(cpn Component) {
	cps.c = append(cps.c, cpn)
}
func (cps *Composite) Remove(cpn Component) {
	for i := range cps.c {
		if cps.c[i] == cpn {
			cps.c[i] = nil
		}
	}
}
func (cps *Composite) DisPlay(deepth int) {
	for i := 0; i < deepth; i++ {
		fmt.Printf("-")
	}
	fmt.Println(cps.name)
	for _, c := range cps.c {
		c.DisPlay(deepth + 1)
	}
}

func main() {
	head := Composite{name: "head"}
	leafA := &leaf{"leafA"}
	leafB := &leaf{"leafB"}
	leafC := &leaf{"leafC"}
	leafD := &leaf{"leafD"}
	head.Add(leafA)
	head.Add(leafB)
	cps := Composite{name: "Composite"}
	head.Add(&cps)
	for _, c := range head.c {
		if _, ok := c.(*Composite); ok {
			c.Add(leafC)
			c.Add(leafD)
		}
	}
	head.DisPlay(1)
}

package main

import "fmt"

type Builder interface {
	dowork()
}

type concreteProductA struct{}

func (cpa *concreteProductA) dowork() {
	fmt.Println("concreteProductA")
}

type concreteProductB struct{}

func (cpb *concreteProductB) dowork() {
	fmt.Println("concreteProductB")
}

type Drictor struct {
	Builder
}

func newConstruct(b Builder) Drictor {
	return Drictor{b}
}

func (this *Drictor) Construct() {
	this.dowork()
}

func main() {
	cpa := concreteProductA{}
	d := newConstruct(&cpa)
	d.Construct()
	cpb := concreteProductB{}
	d = newConstruct(&cpb)
	d.Construct()
}

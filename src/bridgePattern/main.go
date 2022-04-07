package main

import "fmt"

type Implementor interface {
	run()
}

type ConcretImplementorA struct{}

func (cpa *ConcretImplementorA) run() {
	fmt.Println("ConcretImplementorA run")
}

type ConcretImplementorB struct{}

func (cpb *ConcretImplementorB) run() {
	fmt.Println("ConcretImplementorB run")
}

type Abstrction interface {
	run(*Implementor)
}

type RefinedAbstrction struct{}

func (rfa *RefinedAbstrction) run(ipl Implementor) {
	ipl.run()
}

func main() {
	cpa := ConcretImplementorA{}
	cpb := ConcretImplementorB{}
	fra := RefinedAbstrction{}
	fra.run(&cpa)
	fra.run(&cpb)

}

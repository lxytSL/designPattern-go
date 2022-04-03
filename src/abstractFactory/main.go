package main

import "fmt"

type productA interface {
	doworkProductA()
}

type concreteProductAA struct{}

func (p *concreteProductAA) doworkProductA() {
	fmt.Println("concreteProductAA")
}

type concreteProductAB struct{}

func (p *concreteProductAB) doworkProductA() {
	fmt.Println("concreteProductAB")
}

type productB interface {
	doworkProductB()
}
type concreteProductBA struct{}

func (p *concreteProductBA) doworkProductB() {
	fmt.Println("concreteProductBA")
}

type concreteProductBB struct{}

func (p *concreteProductBB) doworkProductB() {
	fmt.Println("concreteProductBB")
}

type abstractFactory interface {
	getProductA(productName string) productA
	getProductB(productName string) productB
}
type FactoryA struct{}

func (f *FactoryA) getProductA(productName string) productA {
	switch productName {
	case "AA":
		return &concreteProductAA{}
	case "AB":
		return &concreteProductAB{}
	}
	return nil
}
func (f *FactoryA) getProductB(productName string) productB {
	return nil
}

type FactoryB struct{}

func (f *FactoryB) getProductA(productName string) productA {
	return nil
}
func (f *FactoryB) getProductB(productName string) productB {
	switch productName {
	case "BA":
		return &concreteProductBA{}
	case "BB":
		return &concreteProductBB{}
	}
	return nil
}

type factoryproduct struct{}

func (fp factoryproduct) GetFactory(factoryNmae string) abstractFactory {
	switch factoryNmae {
	case "factoryA":
		return &FactoryA{}
	case "factoryB":
		return &FactoryB{}
	}
	return nil
}

func newFactoryProduct() *factoryproduct {
	return &factoryproduct{}
}

func main() {
	fp := newFactoryProduct()
	factoryA := fp.GetFactory("factoryA")
	factoryB := fp.GetFactory("factoryB")

	productAA := factoryA.getProductA("AA")
	productAB := factoryA.getProductA("AB")

	productBA := factoryB.getProductB("BA")
	productBB := factoryB.getProductB("BB")

	productAA.doworkProductA()
	productAB.doworkProductA()
	productBA.doworkProductB()
	productBB.doworkProductB()
}

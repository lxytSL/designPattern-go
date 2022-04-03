package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type ConcretePrototype1 struct {
	Prototype1 string
}
type ConcretePrototype2 struct {
	Prototype2 string
}

type Prototype struct {
	ConcretePrototype1
	ConcretePrototype2
}

// low copy
func (p *Prototype) Clone() *Prototype {
	pc := *p
	return &pc
}

// deepCopy
func (p *Prototype) BackUp() *Prototype {
	pc := new(Prototype)
	if err := deepCopy(pc, p); err != nil {
		fmt.Printf("deepCopy err:%v\n", err)
		panic("deepCopy")
	}
	return pc
}

func deepCopy(dst, src interface{}) error {
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buff.Bytes())).Decode(dst)
}

func main() {
	p := Prototype{ConcretePrototype1{"ConcretePrototype1"}, ConcretePrototype2{"ConcretePrototype2"}}
	fmt.Printf("Prototy %p\n", &p)
	pc := p.Clone()
	fmt.Printf("clone %p\n", &pc)
	deeppc := p.BackUp()
	fmt.Printf("deep clone %p\n", &deeppc)
}

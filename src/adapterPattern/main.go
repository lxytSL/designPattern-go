package main

import "fmt"

type Target interface {
	requests()
}

type adaptee interface {
	specificrequest()
}

type adpteeImpl struct{}

func (m1 *adpteeImpl) specificrequest() {
	fmt.Println("adaptee")
}

type adapter struct {
	*adpteeImpl
}

func (adp *adapter) requests() {
	adp.adpteeImpl.specificrequest()
}

func NewAdaptee() *adpteeImpl {
	return &adpteeImpl{}
}

func NewAdapter(m *adpteeImpl) *adapter {
	return &adapter{
		adpteeImpl: m,
	}
}

func main() {
	m2 := NewAdaptee()
	tagret := NewAdapter(m2)
	tagret.requests()
}

package mychain

import (
	"fmt"
)

type IMyChain interface {
	WithName(name string) *myChain
	WithAge(age int) *myChain
	Err() error
	PrintInfo()
}

func NewMyChain() IMyChain {
	return &myChain{}
}

type myChain struct {
	name string
	age  int
	err  error
}

func (c *myChain) WithName(name string) *myChain {
	c.name = name
	return c
}

func (c *myChain) WithAge(age int) *myChain {
	c.age = age
	return c
}

func (c *myChain) Err() error {
	return c.err
}

func (c *myChain) PrintInfo() {
	fmt.Printf("my name is %s and i'm %d years old ", c.name, c.age)
}

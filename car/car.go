package car

import (
	"errors"
	"time"
)

type Car struct {
	Name  string
	Price float32
}

func (c *Car) SetName(name string) string {
	if name != "" {
		c.Name = name
	}

	time.Sleep(10 * time.Second)

	return c.Name
}

func New(name string, price float32) (*Car, error) {
	if name == "" {
		return nil, errors.New("missing name")
	}

	return &Car{
		Name:  name,
		Price: price,
	}, nil
}

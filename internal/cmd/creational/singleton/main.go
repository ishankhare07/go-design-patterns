package main

import (
	"fmt"
	"patterns/internal/pkg/creational/singleton"
)

func main() {
	c1 := singleton.GetInstance()
	fmt.Println("counter from 1", c1.AddOne())

	c2 := singleton.GetInstance()
	fmt.Println("counter from 2", c2.AddOne())
}

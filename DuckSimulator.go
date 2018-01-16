package main

import (
	"fmt"
)

func main() {

	var m mallard
	var r rubberduck

	doyourstuff(m)
	doyourstuff(r)
}

func doyourstuff(d duck) {
	fmt.Println(d.fly())
	fmt.Println(d.quack())
	fmt.Println(d.swim())
}

type duck interface {
	fly() string
	quack() string
	swim() string
}

type mallard struct {
}

func (m mallard) fly() string {
	return "flying"
}

func (m mallard) quack() string {
	return "quack quack"
}

func (m mallard) swim() string {
	return "swimming"
}

type rubberduck struct {
}

func (r rubberduck) fly() string {
	return "cant fly"
}

func (r rubberduck) quack() string {
	return "squeak squeak"
}

func (r rubberduck) swim() string {
	return "floating"
}

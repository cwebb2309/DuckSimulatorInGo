package main

import "testing"

func TestMallardFly(t *testing.T) {

	var m mallard
	var s string

	s = m.fly()
	if s != "flying" {
		t.Errorf("Expected flying got %s", s)
	}
}

func TestMallardSwim(t *testing.T) {

	var m mallard
	var s string

	s = m.swim()
	if s != "swimming" {
		t.Errorf("Expected swimming got %s", s)
	}
}

func TestMallardQuack(t *testing.T) {

	var m mallard
	var s string

	s = m.quack()
	if s != "quack quack" {
		t.Errorf("Expected quack quack got %s", s)
	}
}

func TestRubberduckFly(t *testing.T) {

	var r rubberduck
	var s string

	s = r.fly()
	if s != "cant fly" {
		t.Errorf("Expected cant fly got %s", s)
	}
}

func TestRubberduckSwim(t *testing.T) {

	var r rubberduck
	var s string

	s = r.swim()
	if s != "floating" {
		t.Errorf("Expected floating got %s", s)
	}
}

func TestRubberduckQuack(t *testing.T) {

	var r rubberduck
	var s string

	s = r.quack()
	if s != "squeak squeak" {
		t.Errorf("Expected squeak squeak got %s", s)
	}
}

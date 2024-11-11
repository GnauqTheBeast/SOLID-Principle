package main

import "fmt"

// Definition: Clients should not be forced to depend on interfaces they do not use.

// Bad code
/*
type Worker interface {
	Work()
	Eat()
}

type Human struct{}

func (h *Human) Work() {
	fmt.Println("Human working")
}

func (h *Human) Eat() {
	fmt.Println("Human eating")
}

type Robot struct{}

func (r *Robot) Work() {
	fmt.Println("Robot working")
}

// Problem: The Robot is forced to implement Eat() even though it doesn't eat.
func (r *Robot) Eat() {
	// Robots don't eat
	panic("Robot can't eat")
}
*/

// Solution: The interfaces are split into Worker and Eater, so clients only depend on what they need.
type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Human struct{}

func (h *Human) Work() {
	fmt.Println("Human working")
}

func (h *Human) Eat() {
	fmt.Println("Human eating")
}

type Robot struct{}

func (r *Robot) Work() {
	fmt.Println("Robot working")
}

func main() {}

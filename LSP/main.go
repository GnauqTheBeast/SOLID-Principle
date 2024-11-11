package main

import "fmt"

// Definition: Subtypes must be substitutable for their base types without altering the correctness of the program.

// Bad code
/*
type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s *Sparrow) Fly() {
	fmt.Println("Sparrow flying")
}

type Ostrich struct{}

func (o *Ostrich) Fly() {
	// Ostriches can't fly
	panic("Ostrich can't fly")
}
*/

// Solution: The Move() method is more appropriate since it encompasses different forms of movement.
type Bird interface {
	Move()
}

type Sparrow struct{}

func (s *Sparrow) Move() {
	fmt.Println("Sparrow flying")
}

type Ostrich struct{}

func (o *Ostrich) Move() {
	fmt.Println("Ostrich running")
}

func main() {}

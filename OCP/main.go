package main

import "fmt"

// Definition: A class should be open for extension but closed for modification.

// Bad code
/*
type Payment struct {
	Method string
}

func (p *Payment) Process(amount float64) {
	if p.Method == "credit_card" {
		fmt.Println("Processing credit card payment:", amount)
	} else if p.Method == "paypal" {
		fmt.Println("Processing PayPal payment:", amount)
	}
}
*/

// Solution: We can add new payment methods by implementing the PaymentProcessor interface without modifying existing code.
type PaymentProcessor interface {
	Process(amount float64)
}

type CreditCard struct{}

func (c *CreditCard) Process(amount float64) {
	fmt.Println("Processing credit card payment:", amount)
}

type PayPal struct{}

func (p *PayPal) Process(amount float64) {
	fmt.Println("Processing PayPal payment:", amount)
}

func ProcessPayment(p PaymentProcessor, amount float64) {
	p.Process(amount)
}

func main() {}

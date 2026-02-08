package main

import "fmt"

// Factory Method Pattern
//
// 5-Year-Old Explanation:
// Imagine a magical Toy Factory. You don't need to know how to build a car or a doll.
// You just go to the factory window and say "I want a car!" or "I want a doll!".
// The factory hands you the toy, ready to play!
//
// Real World Scenario:
// A payment processing system. You have a processPayment method, but based on the user's choice
// (Credit Card, PayPal, Bitcoin), you need to create a specific PaymentProcessor object
// (like CreditCardProcessor or PayPalProcessor) to handle the transaction.

// Toy is something we can play with.
type Toy interface {
	Play()
}

// Car is a type of Toy.
type Car struct{}

func (c *Car) Play() {
	fmt.Println("Car says: Vroom Vroom! 🚗")
}

// Doll is another type of Toy.
type Doll struct{}

func (d *Doll) Play() {
	fmt.Println("Doll says: Hello, let's have tea! 🧸")
}

// ToyFactory is where we ask for new toys.
func MakeToy(toyType string) Toy {
	switch toyType {
	case "car":
		return &Car{}
	case "doll":
		return &Doll{}
	default:
		return nil
	}
}

func main() {
	fmt.Println("--- Factory Method Pattern: The Magic Toy Factory ---")

	// I want a car!
	fmt.Println("Kid: Can I have a car?")
	myToy1 := MakeToy("car")
	myToy1.Play()

	// I want a doll!
	fmt.Println("Kid: Can I have a doll?")
	myToy2 := MakeToy("doll")
	myToy2.Play()
}

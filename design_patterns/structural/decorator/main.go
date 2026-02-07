package main

import "fmt"

// Decorator Pattern
//
// 5-Year-Old Explanation:
// Imagine you have a plain vanilla Ice Cream.
// You can eat it like that.
// BUT, you can also "decorate" it!
// You can add sprinkles.
// You can add chocolate sauce.
// You can add a cherry.
// It's still an Ice Cream, but now it's YUMMIER!

type IceCream interface {
	GetCost() int
	GetDescription() string
}

// BasicIceCream is the plain vanilla scoop.
type BasicIceCream struct{}

func (b *BasicIceCream) GetCost() int {
	return 10 // Cost is $10
}

func (b *BasicIceCream) GetDescription() string {
	return "Vanilla Ice Cream"
}

// -- The Decorators --

// ChocolateSauce adds chocolate to the ice cream.
type ChocolateSauce struct {
	iceCream IceCream // Contains an Ice Cream inside it
}

func (c *ChocolateSauce) GetCost() int {
	return c.iceCream.GetCost() + 5 // Adds $5
}

func (c *ChocolateSauce) GetDescription() string {
	return c.iceCream.GetDescription() + " + Chocolate Sauce"
}

// Sprinkles adds sprinkles.
type Sprinkles struct {
	iceCream IceCream
}

func (s *Sprinkles) GetCost() int {
	return s.iceCream.GetCost() + 2 // Adds $2
}

func (s *Sprinkles) GetDescription() string {
	return s.iceCream.GetDescription() + " + Sprinkles"
}

func main() {
	fmt.Println("--- Decorator Pattern: Making Ice Cream Yummy ---")

	// 1. Plain Ice Cream
	var myIceCream IceCream = &BasicIceCream{}
	fmt.Printf("Order 1: %s. Cost: $%d\n", myIceCream.GetDescription(), myIceCream.GetCost())

	// 2. Add Chocolate Sauce
	myIceCream = &ChocolateSauce{iceCream: myIceCream}
	fmt.Printf("Order 2: %s. Cost: $%d\n", myIceCream.GetDescription(), myIceCream.GetCost())

	// 3. Add Sprinkles
	myIceCream = &Sprinkles{iceCream: myIceCream}
	fmt.Printf("Order 3: %s. Cost: $%d\n", myIceCream.GetDescription(), myIceCream.GetCost())
}

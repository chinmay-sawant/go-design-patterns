package main

import "fmt"

// Functional Options Pattern
//
// 5-Year-Old Explanation:
// Imagine you are ordering a Pizza.
// Every pizza starts with just dough and tomato sauce (The Base).
// But you can choose to ADD extra things if you want!
// "Add Cheese please!"
// "Add Pepperoni please!"
// You don't HAVE to add them, but you CAN if you want to.
//
// Real World Scenario:
// Configuring a Server (like an HTTP server). You start with a default server configuration.
// Then, you can use options like WithPort(8080), WithTimeout(30s), or WithLogger(myLogger)
// to customize it. This is cleaner than having a constructor with 10 different arguments.

// Pizza is what we are making.
type Pizza struct {
	Dough    string
	Sauce    string
	Cheese   bool
	Toppings []string
}

// PizzaOption is a function that changes the Pizza.
// It's like a special instruction card you give to the chef.
type PizzaOption func(*Pizza)

// NewPizza makes a basic pizza, then applies any extra options you asked for.
func NewPizza(opts ...PizzaOption) *Pizza {
	// Start with the default "Basic" pizza
	p := &Pizza{
		Dough:    "Regular",
		Sauce:    "Tomato",
		Cheese:   false, // Default is no extra cheese
		Toppings: []string{},
	}

	// Apply all the options (instructions) one by one
	for _, opt := range opts {
		opt(p)
	}

	return p
}

// WithExtraCheese is an option to add cheese.
func WithExtraCheese() PizzaOption {
	return func(p *Pizza) {
		p.Cheese = true
	}
}

// WithTopping is an option to add a specific topping.
func WithTopping(topping string) PizzaOption {
	return func(p *Pizza) {
		p.Toppings = append(p.Toppings, topping)
	}
}

func main() {
	fmt.Println("--- Functional Options Pattern: Ordering Pizza ---")

	// 1. A Plain Pizza (No options)
	fmt.Println("\nOrder 1: Plain Pizza")
	pizza1 := NewPizza()
	fmt.Printf("Pizza 1: %+v\n", pizza1)

	// 2. A Cheese Pizza with Pepperoni
	fmt.Println("\nOrder 2: Cheese + Pepperoni")
	pizza2 := NewPizza(
		WithExtraCheese(),
		WithTopping("Pepperoni"),
	)
	fmt.Printf("Pizza 2: %+v\n", pizza2)

	// 3. The "Everything" Pizza
	fmt.Println("\nOrder 3: The Works!")
	pizza3 := NewPizza(
		WithExtraCheese(),
		WithTopping("Mushrooms"),
		WithTopping("Olives"),
		WithTopping("Onions"),
	)
	fmt.Printf("Pizza 3: %+v\n", pizza3)
}

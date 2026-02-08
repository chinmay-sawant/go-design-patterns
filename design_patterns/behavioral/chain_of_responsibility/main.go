package main

import "fmt"

// Chain of Responsibility Pattern
//
// 5-Year-Old Explanation:
// Imagine you want to buy a toy that costs $100.
// 1. You ask your Big Brother. He has $10. He says "I can't afford it, ask Dad."
// 2. You ask Dad. He has $50. He says "I can't afford it, ask Mom."
// 3. You ask Mom. She has $200! She says "Ok, I'll buy it."
// The request gets passed up the chain until someone can handle it!
//
// Real World Scenario:
// IT Support Desk. You call Level 1 support. If they can't fix it, they escalate to Level 2.
// If Level 2 can't fix it, they escalate to Level 3 (The Engineers).
// The request moves up the chain until it finds someone who knows the answer.

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(cost int)
}

// BaseHandler helps us link the chain
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) Handler {
	b.next = next
	return next
}

func (b *BaseHandler) HandleNext(cost int) {
	if b.next != nil {
		b.next.Handle(cost)
	} else {
		fmt.Println("End of chain: No one can afford this!")
	}
}

// -- Concrete Handlers --

// Brother
type Brother struct {
	BaseHandler
}

func (b *Brother) Handle(cost int) {
	if cost <= 10 {
		fmt.Printf("Brother: I have $10. I can buy this toy for $%d!\n", cost)
	} else {
		fmt.Printf("Brother: I only have $10. Too expensive ($%d). Asking Dad...\n", cost)
		b.HandleNext(cost)
	}
}

// Dad
type Dad struct {
	BaseHandler
}

func (d *Dad) Handle(cost int) {
	if cost <= 50 {
		fmt.Printf("Dad: I have $50. I can buy this toy for $%d!\n", cost)
	} else {
		fmt.Printf("Dad: I only have $50. Too expensive ($%d). Asking Mom...\n", cost)
		d.HandleNext(cost)
	}
}

// Mom
type Mom struct {
	BaseHandler
}

func (m *Mom) Handle(cost int) {
	if cost <= 200 {
		fmt.Printf("Mom: I have $200. I can buy this toy for $%d!\n", cost)
	} else {
		fmt.Printf("Mom: I only have $200. Too expensive ($%d). We can't buy it.\n", cost)
		m.HandleNext(cost)
	}
}

func main() {
	fmt.Println("--- Chain of Responsibility: Buying a Toy ---")

	// Create the chain components
	brother := &Brother{}
	dad := &Dad{}
	mom := &Mom{}

	// Link them: Brother -> Dad -> Mom
	brother.SetNext(dad).SetNext(mom)

	// Test 1: Cheap toy
	fmt.Println("\nRequest: Toy costs $5")
	brother.Handle(5)

	// Test 2: Medium toy
	fmt.Println("\nRequest: Toy costs $40")
	brother.Handle(40)

	// Test 3: Expensive toy
	fmt.Println("\nRequest: Toy costs $150")
	brother.Handle(150)

	// Test 4: Super expensive toy
	fmt.Println("\nRequest: Toy costs $500")
	brother.Handle(500)
}

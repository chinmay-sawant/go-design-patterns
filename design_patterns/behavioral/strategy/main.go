package main

import "fmt"

// Strategy Pattern
//
// 5-Year-Old Explanation:
// Imagine you want to go to the park.
// You can choose DIFFERENT ways to get there:
// 1. Walk (Slow but free)
// 2. Bike (Faster)
// 3. Car (Super fast but needs gas)
// The "Strategy" is just "How do we get there?".
// You can pick a different strategy depending on if you are in a rush!
//
// Real World Scenario:
// Saving a file. You might have a "SaveStrategy".
// - SaveToLocalDisk (for offline use)
// - SaveToS3 (cloud storage)
// - SaveToDatabase (binary data)
// The application just says "Save(file)", and the specific Strategy handles WHERE it goes.

// TravelStrategy is the interface for our travel method.
type TravelStrategy interface {
	Travel(destination string)
}

// -- Concrete Strategies --

// WalkStrategy
type WalkStrategy struct{}

func (w *WalkStrategy) Travel(destination string) {
	fmt.Printf("Walking to %s. It will take a long time, but it's healthy!\n", destination)
}

// CarStrategy
type CarStrategy struct{}

func (c *CarStrategy) Travel(destination string) {
	fmt.Printf("Driving to %s. Vroom! We'll be there fast!\n", destination)
}

// BusStrategy
type BusStrategy struct{}

func (b *BusStrategy) Travel(destination string) {
	fmt.Printf("Taking the bus to %s. We sit with other people!\n", destination)
}

// -- The Context (The Traveler) --
type Traveler struct {
	strategy TravelStrategy
}

func (t *Traveler) SetStrategy(s TravelStrategy) {
	t.strategy = s
}
func (t *Traveler) GoTo(destination string) {
	t.strategy.Travel(destination)
}

func main() {
	fmt.Println("--- Strategy Pattern: Choosing How to Travel ---")

	me := &Traveler{}

	// 1. I have time, let's walk
	fmt.Println("\nSituation 1: Sunny day, plenty of time.")
	me.SetStrategy(&WalkStrategy{})
	me.GoTo("The Park")

	// 2. I'm late! Take the car!
	fmt.Println("\nSituation 2: I'm LATE!")
	me.SetStrategy(&CarStrategy{})
	me.GoTo("The Airport")

	// 3. Car broke down, take the bus
	fmt.Println("\nSituation 3: Car won't start.")
	me.SetStrategy(&BusStrategy{})
	me.GoTo("School")
}

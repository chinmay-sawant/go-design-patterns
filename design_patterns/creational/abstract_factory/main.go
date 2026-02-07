package main

import "fmt"

// Abstract Factory Pattern
//
// 5-Year-Old Explanation:
// Imagine two furniture stores: one sells "Modern" stuff (cool and new),
// and the other sells "Old-Fashioned" stuff (fancy and classic).
// If you go to the "Modern Store", EVERY piece of furniture you buy (chair, sofa) matches the modern style.
// If you go to the "Old-Fashioned Store", EVERYTHING matches the old style.
// You don't mix and match by accident!

// Chair interface
type Chair interface {
	Sit()
}

// Sofa interface
type Sofa interface {
	LieDown()
}

// --- Modern Furniture ---
type ModernChair struct{}
func (m *ModernChair) Sit() { fmt.Println("Sitting on a cool Modern Chair.") }

type ModernSofa struct{}
func (m *ModernSofa) LieDown() { fmt.Println("Sleeping on a cool Modern Sofa.") }

// --- Old-Fashioned (Victorian) Furniture ---
type VictorianChair struct{}
func (v *VictorianChair) Sit() { fmt.Println("Sitting on a fancy Victorian Chair.") }

type VictorianSofa struct{}
func (v *VictorianSofa) LieDown() { fmt.Println("Sleeping on a fancy Victorian Sofa.") }

// --- The Factory Interface ---
// This tells us what any furniture factory MUST be able to make.
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}

// --- Concrete Factories ---

// ModernFactory only makes Modern stuff
type ModernFactory struct{}
func (m *ModernFactory) CreateChair() Chair { return &ModernChair{} }
func (m *ModernFactory) CreateSofa() Sofa   { return &ModernSofa{} }

// VictorianFactory only makes Victorian stuff
type VictorianFactory struct{}
func (v *VictorianFactory) CreateChair() Chair { return &VictorianChair{} }
func (v *VictorianFactory) CreateSofa() Sofa   { return &VictorianSofa{} }

func main() {
	fmt.Println("--- Abstract Factory Pattern: Matching Furniture Sets ---")

	// Let's buy a Modern set!
	fmt.Println("\nVisiting the Modern Store...")
	var modernFactory FurnitureFactory = &ModernFactory{}
	chair1 := modernFactory.CreateChair()
	sofa1 := modernFactory.CreateSofa()
	chair1.Sit()
	sofa1.LieDown()

	// Let's buy a Victorian set!
	fmt.Println("\nVisiting the Victorian Store...")
	var oldFactory FurnitureFactory = &VictorianFactory{}
	chair2 := oldFactory.CreateChair()
	sofa2 := oldFactory.CreateSofa()
	chair2.Sit()
	sofa2.LieDown()
}

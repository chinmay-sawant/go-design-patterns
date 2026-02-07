package main

import "fmt"

// Builder Pattern
//
// 5-Year-Old Explanation:
// Imagine you are building a LEGO castle.
// You don't just "poof" a whole castle into existence.
// You build it step-by-step:
// 1. First, you build the walls.
// 2. Then, you add a roof.
// 3. Maybe you add a moat (water) around it.
// The Builder lets you choose exactly what parts you want!

// Castle is the complex object we are building.
type Castle struct {
	Walls string
	Roof  string
	Moat  string
}

// Show tells us what our castle looks like.
func (c *Castle) Show() {
	fmt.Printf("Castle details: Walls=[%s], Roof=[%s], Moat=[%s]\n", c.Walls, c.Roof, c.Moat)
}

// CastleBuilder helps us build the castle step-by-step.
type CastleBuilder struct {
	castle Castle
}

func NewCastleBuilder() *CastleBuilder {
	return &CastleBuilder{}
}

// BuildWalls adds walls to our castle.
func (b *CastleBuilder) BuildWalls(style string) *CastleBuilder {
	b.castle.Walls = style
	return b
}

// BuildRoof adds a roof.
func (b *CastleBuilder) BuildRoof(style string) *CastleBuilder {
	b.castle.Roof = style
	return b
}

// BuildMoat adds a moat.
func (b *CastleBuilder) BuildMoat(typeOfMoat string) *CastleBuilder {
	b.castle.Moat = typeOfMoat
	return b
}

// GetResult returns the finished castle.
func (b *CastleBuilder) GetResult() Castle {
	return b.castle
}

func main() {
	fmt.Println("--- Builder Pattern: Building a Lego Castle ---")

	// Builder 1: A simple castle
	fmt.Println("\nBuilder 1 making a simple castle...")
	builder1 := NewCastleBuilder()
	castle1 := builder1.BuildWalls("Stone").BuildRoof("Wood").GetResult()
	castle1.Show()

	// Builder 2: A super fancy castle!
	fmt.Println("\nBuilder 2 making a SUPER fancy castle...")
	builder2 := NewCastleBuilder()
	castle2 := builder2.
		BuildWalls("Diamond").
		BuildRoof("Gold").
		BuildMoat("Lava").
		GetResult()
	castle2.Show()
}

package main

import "fmt"

// Adapter Pattern
//
// 5-Year-Old Explanation:
// Imagine you are traveling to a different country.
// Your phone charger has a square plug, but the wall socket only takes round plugs!
// You can't plug it in directly.
// You need an "Adapter" (a special connector) that changes your square plug into a round one.
// Now you can charge your phone!
//
// Real World Scenario:
// Integrating a legacy system with a modern one.
// The new system expects data in JSON format, but the old legacy system gives data in XML.
// You create an "XMLtoJSONAdapter" so the new system can use the old data without crashing.

// -- The "Client" (What we want to use) --
type Computer interface {
	InsertSquareUSB()
}

// Mac is a computer that accepts square USBs.
type Mac struct{}

func (m *Mac) InsertSquareUSB() {
	fmt.Println("Success: Square USB connected to Mac.")
}

// -- The "Service" (What we have, but it doesn't fit) --
type WindowsMachine struct{}

func (w *WindowsMachine) InsertRoundUSB() {
	fmt.Println("Success: Round USB connected to Windows Machine.")
}

// -- The Adapter --
// WindowsAdapter makes a WindowsMachine look like a Computer (Mac-style input).
type WindowsAdapter struct {
	windowMachine *WindowsMachine
}

func (w *WindowsAdapter) InsertSquareUSB() {
	fmt.Println("Adapter: Converting Square USB signal to Round USB...")
	w.windowMachine.InsertRoundUSB()
}

func main() {
	fmt.Println("--- Adapter Pattern: The Power Plug Converter ---")

	// 1. Using a Mac directly (Fits perfectly)
	fmt.Println("\nUser: I have a Mac.")
	mac := &Mac{}
	mac.InsertSquareUSB()

	// 2. Trying to use a Windows Machine with a Square USB
	fmt.Println("\nUser: I have a Windows Machine, but my cable is Square!")
	windowsMachine := &WindowsMachine{}

	// We need an adapter!
	adapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	// Now we can use the "Square USB" method on the adapter
	fmt.Println("User: Pugging Square cable into Adapter...")
	adapter.InsertSquareUSB()
}

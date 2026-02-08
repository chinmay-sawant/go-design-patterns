package main

import "fmt"

// Prototype Pattern
//
// 5-Year-Old Explanation:
// Imagine you drew a perfect picture of a robot.
// If your friend wants the SAME robot, you don't have to draw it all over again from scratch.
// You can just put your drawing in a photocopy machine and press "COPY"!
// Now you have two robots!
//
// Real World Scenario:
// Game development. Spawning hordes of enemies (e.g., "Goblins"). Instead of creating a new
// Goblin object from scratch (loading textures, sounds, stats) every single time, you create one
// "Master Goblin" and just "clone" it 100 times, maybe tweaking their position or HP slightly.

// Robot is our prototype interface.
type Robot interface {
	Clone() Robot
	SayVal() string
}

// BlueRobot is a specific type of robot.
type BlueRobot struct {
	Name  string
	Power int
}

// Clone creates a copy of the BlueRobot.
func (r *BlueRobot) Clone() Robot {
	// We create a NEW BlueRobot and copy the values over.
	return &BlueRobot{
		Name:  r.Name + "_Clone", // Adding "_Clone" just so we can see it's a copy
		Power: r.Power,
	}
}

func (r *BlueRobot) SayVal() string {
	return fmt.Sprintf("I am %s and my power is %d", r.Name, r.Power)
}

func main() {
	fmt.Println("--- Prototype Pattern: Cloning Robots ---")

	// 1. Create the original robot
	originalRobot := &BlueRobot{Name: "Robo-One", Power: 100}
	fmt.Println("Original:", originalRobot.SayVal())

	// 2. Clone it! (Hit the copy button)
	clonedRobot := originalRobot.Clone()
	fmt.Println("Clone:   ", clonedRobot.SayVal())

	// Prove they are different objects
	fmt.Printf("\n(Proof: Original is at address %p, Clone is at address %p)\n", originalRobot, clonedRobot)
}

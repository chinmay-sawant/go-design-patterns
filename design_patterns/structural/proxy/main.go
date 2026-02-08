package main

import "fmt"

// Proxy Pattern
//
// 5-Year-Old Explanation:
// Imagine a secret clubhouse.
// You can't just walk in.
// There is a Security Guard (The Proxy) at the door.
// The Guard asks: "What is the password?"
// If you know the password, the Guard lets you in (Accesses the Real Object).
// If you don't know it, the Guard says "STOP! Go away!"
//
// Real World Scenario:
// A Cache Proxy or Access Control.
// - Cache: Before asking the real database (which is slow), check if we already have the answer in memory (fast).
// - Access Control: Before letting a user delete a file, check if they are an Admin.

type Door interface {
	Open(password string)
}

// RealDoor is the actual door to the clubhouse.
type RealDoor struct{}

func (r *RealDoor) Open(password string) {
	fmt.Println("Door: Squeak... The door opens. Welcome to the Secret Clubhouse!")
}

// SecurityProxy is the guard protecting the door.
type SecurityProxy struct {
	door *RealDoor
}

func (s *SecurityProxy) Open(password string) {
	if password == "secret123" {
		fmt.Println("Proxy: Password correct! Opening the door.")
		if s.door == nil {
			s.door = &RealDoor{}
		}
		s.door.Open(password)
	} else {
		fmt.Println("Proxy: WRONG PASSWORD! You cannot enter!")
	}
}

func main() {
	fmt.Println("--- Proxy Pattern: The Security Guard ---")

	var myDoor Door = &SecurityProxy{}

	// 1. Try with wrong password
	fmt.Println("\nKid: Can I come in? (Password: pizza)")
	myDoor.Open("pizza")

	// 2. Try with correct password
	fmt.Println("\nKid: Can I come in? (Password: secret123)")
	myDoor.Open("secret123")
}

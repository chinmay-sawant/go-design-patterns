package main

import (
	"fmt"
	"sync"
)

// Singleton Pattern
//
// 5-Year-Old Explanation:
// Imagine a country can only have ONE President.
// No matter how many people ask "Who is the President?", they should all get the SAME answer.
// You can't have two Presidents at the same time!

// president is our private struct. No one can make a new one directly.
type president struct {
	name string
}

// instance is the ONE and ONLY instance of president we will ever have.
var instance *president

// once is a special tool ("magic lock") that makes sure we only create the president ONCE.
var once sync.Once

// GetPresident is the way to reach the President.
func GetPresident() *president {
	// once.Do makes sure that the code inside it runs only ONE time,
	// even if a million people call GetPresident at the same time.
	once.Do(func() {
		fmt.Println("Creating the President for the very first time!")
		instance = &president{name: "Mr. Gopher"}
	})
	return instance
}

func main() {
	fmt.Println("--- Singleton Pattern: The One and Only President ---")

	// First person asks for the President
	fmt.Println("Person 1: calls GetPresident()")
	p1 := GetPresident()
	fmt.Printf("Person 1 sees President: %s\n", p1.name)

	// Second person asks for the President
	fmt.Println("Person 2: calls GetPresident()")
	p2 := GetPresident()
	fmt.Printf("Person 2 sees President: %s\n", p2.name)

	// Let's check if they are the SAME president
	if p1 == p2 {
		fmt.Println("Result: They are the SAME President! (Works as expected)")
	} else {
		fmt.Println("Result: They are DIFFERENT! (Something is wrong)")
	}
}

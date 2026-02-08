package main

import "fmt"

// Pipeline Pattern
//
// 5-Year-Old Explanation:
// Imagine a factory line making Cakes.
// It happens in steps:
// Step 1: Someone puts dough in a pan (Generator).
// Step 2: Someone bakes it (Stage 1).
// Step 3: Someone puts icing on it (Stage 2).
// Step 4: Someone puts a cherry on top (Stage 3).
// The cake moves from person to person.
// While the Baker is baking Cake #1, the Generator is already preparing Cake #2!
//
// Real World Scenario:
// Data Processing ETL (Extract, Transform, Load).
// Stage 1: Read lines from a huge log file (Extract).
// Stage 2: Parse the JSON and filter errors (Transform).
// Stage 3: Write the errors to a database (Load).
// Each stage runs in parallel, passing data to the next stage instantly.

// 1. Generator: Converts a list of numbers into a channel
func Generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 2. Square Stage: Takes a number, squares it, passes it on
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			result := n * n
			fmt.Printf("Square Stage: %d * %d = %d\n", n, n, result)
			out <- result
		}
		close(out)
	}()
	return out
}

// 3. AddOne Stage: Takes a number, adds 1, passes it on
func AddOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			result := n + 1
			fmt.Printf("AddOne Stage: %d + 1 = %d\n", n, result)
			out <- result
		}
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("--- Pipeline Pattern: The Number Factory ---")

	// Set up the assembly line
	// Numbers -> [Generator] -> [Square] -> [AddOne] -> Output

	// Step 1: Input
	input := Generator(1, 2, 3, 4)

	// Step 2: Square it
	stage1 := Square(input)

	// Step 3: Add one
	stage2 := AddOne(stage1)

	// Consume the final output
	fmt.Println("\nCollecting Final Results:")
	for result := range stage2 {
		fmt.Printf("Final Result: %d\n", result)
	}
}

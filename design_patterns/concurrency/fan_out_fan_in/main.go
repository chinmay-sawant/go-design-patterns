package main

import (
	"fmt"
	"sync"
)

// Fan-Out / Fan-In Pattern
//
// 5-Year-Old Explanation:
// Fan-Out (Splitting): Imagine you have a bag of candy to wrap.
// You give a handful to Friend A, a handful to Friend B, and a handful to Friend C.
// They all wrap candy at the same time (Fan-Out).
//
// Fan-In (Collecting): When they are done, they all put the wrapped candy back into ONE big bowl.
// You collect everything from everyone into one place (Fan-In).

// Worker: Wraps the candy (Multiplies number by 10)
func CandyWrapper(id int, input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			fmt.Printf("Friend %d wrapping candy %d...\n", id, n)
			out <- n * 10 // "Wrapped"
		}
		close(out)
	}()
	return out
}

// Fan-In: Collects from multiple channels into one
func Merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Function to copy from one channel to the output
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Wait for everyone to finish, then close the output
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("--- Fan-Out/Fan-In: Candy Wrapping Party ---")

	// 1. The input (Bag of candy)
	candyBag := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			candyBag <- i
		}
		close(candyBag)
	}()

	// 2. Fan-Out: Give work to 3 friends
	friend1 := CandyWrapper(1, candyBag)
	friend2 := CandyWrapper(2, candyBag)
	friend3 := CandyWrapper(3, candyBag)

	// 3. Fan-In: Collect all wrapped candy into one bowl
	finalBowl := Merge(friend1, friend2, friend3)

	// 4. Look at the result
	fmt.Println("\nChecking the bowl:")
	count := 0
	for c := range finalBowl {
		fmt.Printf("Got wrapped candy: %d\n", c)
		count++
	}
	fmt.Printf("Total candies wrapped: %d\n", count)
}

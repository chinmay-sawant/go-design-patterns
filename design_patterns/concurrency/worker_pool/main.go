package main

import (
	"fmt"
	"time"
)

// Worker Pool Pattern
//
// 5-Year-Old Explanation:
// Imagine you have a BIG pile of bricks to move.
// You could do it all by yourself, but it would take forever!
// Instead, you hire a TEAM of 3 workers.
// You stand there and point: "You move this brick! You move that brick!"
// The workers grab bricks from the pile as fast as they can until the pile is gone.
// 3 people working is much faster than 1!
//
// Real World Scenario:
// Handling HTTP Requests. A web server receives 10,000 requests per second.
// Creating 10,000 threads would kill the server (Out of Memory).
// Instead, you have a "Worker Pool" of 50 workers.
// The requests go into a queue, and the 50 workers pick them up one by one.

// Worker is a person who does the job.
func Worker(id int, jobs <-chan int, results chan<- int) {
	// The worker keeps watching the "jobs" pile.
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Millisecond * 500) // Pretend work takes time
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 // Send the result back
	}
}

func main() {
	fmt.Println("--- Worker Pool: The Construction Team ---")

	const numJobs = 5

	// 1. Make a channel for jobs (The pile of bricks)
	jobs := make(chan int, numJobs)

	// 2. Make a channel for results (The finished wall)
	results := make(chan int, numJobs)

	// 3. Hire 3 workers
	// We start them immediately, but they will wait until there are jobs.
	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	// 4. Assign jobs (Put bricks in the pile)
	for j := 1; j <= numJobs; j++ {
		fmt.Println("Manager: Adding job", j)
		jobs <- j
	}
	close(jobs) // "That's all the work for today!"

	// 5. Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	fmt.Println("All jobs finished!")
}

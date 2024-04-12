package main

import (
	"fmt"
	"runtime"
	"time"
)

const JOB_COUNT = 30
const WORKER_COUNT = 30

func main() {
	// create channels for our jobs and results
	jobs := make(chan int, JOB_COUNT)
	results := make(chan int, JOB_COUNT)
	done := make(chan bool, JOB_COUNT)

	// create a worker pool
	for range WORKER_COUNT {
		go worker(jobs, results, done)
	}

	fmt.Printf("Total number of Go Routine: %d\n", runtime.NumGoroutine())

	// populate our queue with jobs and then close the channel
	for n := range JOB_COUNT {
		jobs <- n
	}
	fmt.Printf("%d jobs enqueued\n\n", len(jobs))
	close(jobs)

	start := time.Now() // start timer
	// listen for the done signal

	for range done {
	}
	close(results)

	// print the results
	for result := range results {
		fmt.Println(result)
	}
	elapsed := time.Since(start) // end timer
	fmt.Printf("Processing took %s\n", elapsed)

}

func worker(jobs <-chan int, results chan<- int, done chan<- bool) {
	for n := range jobs {
		results <- fib(n)
	}
	done <- true // signal that we're done
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

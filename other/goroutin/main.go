package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to calculate prime numbers up to a certain number
func calculatePrimes(limit int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			// Uncomment this line to see output, but it might slow down the program
			// fmt.Println(i)
		}
	}
}

func main() {
	// Number of goroutines to use
	numGoroutines := 10

	// Define the upper limit for prime number calculation
	limit := 100000

	// Create WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start time for profiling
	start := time.Now()

	// Launch multiple goroutines for calculating prime numbers
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go calculatePrimes(limit, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the duration of execution
	fmt.Println("Execution Time:", time.Since(start))
}

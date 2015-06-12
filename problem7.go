// Problem 7: What is the 10001st prime number?
// This is not the best problem to try for concurrency on (since we care about the ordering of these things)
// but it was still fun!
package main

import "fmt"

func main() {
	var c chan int = make(chan int)
	go primeGenerator(c)
	primeReceiver(c)
}

func primeGenerator(c chan int) {
	for i := 3; ; i++ {
		if isPrime(i) {
			c <- i
		}
	}
}

func primeReceiver(c chan int) {
	primes := []int{2}
	for len(primes) <= 10000 {
		primes = append(primes, <-c)
	}
	fmt.Println(primes[10000])

}

func isPrime(value int) bool {
	if value <= 3 {
		return value > 1
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}

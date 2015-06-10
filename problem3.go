package main

import (
	"fmt"
	"math"
)

/* Okay, some notes. Go is entirely 'pass-by-value', which means that
variables passed to another function are copied. So this code makes sense:
func mutate(arg int) {
  arg = 0
}
func main() {
  x := 5
  mutate(x)
  fmt.Println(x) // x is still 5
}

This is all awesome because it's safer. However, Go has pointers, which means
you can pass a variable pointer and access/change the variable itself. Really
handy if performance/memory is at a premium. This works like below:
func zero(xPtr *int) {
  *xPtr = 0
}
func main() {
  x := 5
  zero(&x) //This is passing a pointer to x. Yeah, I know...
  fmt.Println(x) // x is zero
}
*/

// Problem 3: What is the largest prime factor of the number 600851475143?
func main() {
	x := 600851475143
	// Since we are starting from the highest possible prime, we don't have
	// to do the whole 'find the max in this list' thing
	for i := sqrt(x); i >= 1; i-- {
		if x%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}

func isPrime(candidate int) bool {
	if candidate <= 1 {
		return false
	}

	// This is really smart. Different algorithm starts from the highest possible
	// prime and counts down - prevents a lot of time being wasted.
	for i := sqrt(candidate); i >= 1; i-- {
		if i == 1 {
			return true
		}

		if candidate%i == 0 {
			return false
		}
	}
	return true
}

func sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

package main

import "fmt"

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
	fmt.Println(getMax(getPrimes(factorize(x))))
}

func getPrimes(candidates []int) []int {
	primes := []int{}
	for _, val := range candidates {
		if len(factorize(val)) == 2 {
			primes = append(primes, val)
		}
	}
	return primes
}

func factorize(number int) []int {
	factors := []int{1, number}
	iterator := number/2 + 1
	for i := 2; i <= iterator; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func getMax(list []int) int {
	max := list[0]
	for _, val := range list {
		if val > max {
			max = val
		}
	}
	return max
}

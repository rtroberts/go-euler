// Problem 5: What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main

import "fmt"

func main() {
	for i := 20; true; i++ {
		if isDivisible(i) {
			fmt.Println(i)
			break
		}
	}
}

func isDivisible(p int) bool {
	for i := 1; i < 21; i++ {
		if p%i != 0 {
			return false
		}
	}
	return true
}

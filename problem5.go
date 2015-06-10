// Problem 5: What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main

import "fmt"

func main() {
	cont := true
	for i := 20; cont; i++ {
		// I tried adding 'go' here and the program took far longer. I'm wondering if
		// maybe I need channels to do this right?
		isDivisible(i, &cont)
	}
}

func isDivisible(num int, cont *bool) bool {
	for i := 1; i < 21; i++ {
		if num%i != 0 {
			return false
		}
	}
	fmt.Println(num)
	*cont = false
	return true
}

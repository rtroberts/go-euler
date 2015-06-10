// Problem 6: Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package main

import "fmt"

func main() {
	diff := (sumOfSquares(100) - squareOfSum(100))
	fmt.Println(diff)
}

func squareOfSum(num int) int {
	sumnum := 0
	for i := 1; i <= num; i++ {
		sumnum = sumnum + i
	}
	return sumnum * sumnum
}

func sumOfSquares(num int) int {
	sumnum := 0
	for i := 1; i <= num; i++ {
		sumnum = sumnum + i*i
	}
	return sumnum
}

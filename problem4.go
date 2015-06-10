// Problem 4: Find the largest palindrome made from the product of two 3-digit numbers.
package main

import "fmt"
import "strconv"

func main() {
	for i, j := 999, 999; i > 99; i, j = i-1, j-1 {
		if isPalindrome(i * j) {
			fmt.Println(i * j)
			break
		}
	}
}

func isPalindrome(p int) bool {
	str := strconv.Itoa(p)
	// Runes are the golang 'character', basically
	// Convert the string to a slice of runes
	runes := []rune(str)
	// i starts at 0, j starts at length. We swap the runes in-place.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	if string(runes) == str {
		return true
	}
	return false
}

package main

import "fmt"

func main() {
	// Make an empty list
	list := []int{}
	// Iterate over the numbers
	for i := 0; i < 1001; i++ {
		if i%3 == 0 || i%5 == 0 {
			// add them to the list to be summed
			list = append(list, i)
		}
	}
	fmt.Println(summation(list))
}

func summation(ilist []int) int {
	total := 0
	for _, val := range ilist {
		total = total + val
	}
	return total
}

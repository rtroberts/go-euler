package main

import "fmt"

func main() {
	list := []int{1, 2}
	sumlist := []int{}
	for list[len(list)-1] < 4000000 {
		list = fibonacci(list)
	}

	for _, val := range list {
		if val%2 == 0 {
			sumlist = append(sumlist, val)
		}
	}
	fmt.Println(summation(sumlist))
}

func fibonacci(list []int) []int {
	next := list[len(list)-2] + list[len(list)-1]
	list = append(list, next)
	return list
}

func summation(ilist []int) int {
	total := 0
	for _, val := range ilist {
		total = total + val
	}
	return total
}

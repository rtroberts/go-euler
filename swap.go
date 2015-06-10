package main

import "fmt"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

package main

import "fmt"
import "os"
import "strconv"

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fibonacci(input))
}

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

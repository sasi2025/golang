package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	//  Here we Start with first two numbers
	seq := []int{0, 1}
	for i := 2; i < n; i++ {
		next := seq[i-1] + seq[i-2]
		seq = append(seq, next)
	}
	return seq
}

func main() {
	var length int
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)

	result := fibonacci(length)

	// Printing  in required format
	for i, v := range result {
		if i == len(result)-1 {
			fmt.Print(v)
		} else {
			fmt.Print(v, ",")
		}
	}
}

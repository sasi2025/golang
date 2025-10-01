package main

import "fmt"

// function to remove duplicates
func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool) //map collection of key value pairs [integer from arrayc]
	result := []int{}

	for _, val := range arr {
		if !seen[val] { // if not already present
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 2, 5, 4, 1, 4, 5}
	fmt.Println("Original Array:", arr)

	uniqueArr := removeDuplicates(arr)
	fmt.Println("Array after removing duplicates:", uniqueArr)
}

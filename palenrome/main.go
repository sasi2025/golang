package main

import "fmt"

// Function to expand from the center and return the longest palindrome
func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// When loop ends, left and right have moved one step extra
	return s[left+1 : right]
}

// Main function to find the longest palindrome
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	longest := ""
	for i := 0; i < len(s); i++ {
		// Odd length palindrome (center at i)
		p1 := expandAroundCenter(s, i, i)
		if len(p1) > len(longest) {
			longest = p1
		}

		// Even length palindrome (center between i and i+1)
		p2 := expandAroundCenter(s, i, i+1)
		if len(p2) > len(longest) {
			longest = p2
		}
	}
	return longest
}

func main() {
	s := "babad"
	result := longestPalindrome(s)
	fmt.Println("Input:", s)
	fmt.Println("Longest Palindromic Substring:", result)
}

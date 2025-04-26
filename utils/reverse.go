package utils

import "strings"

// ReverseString reverses the input string, properly handling UTF-8 characters
func ReverseString(input string) string {
	// Convert string to runes to properly handle UTF-8 characters
	runes := []rune(input)
	
	// Create output slice of runes
	reversed := make([]rune, len(runes))
	
	// Fill reversed slice from end to start
	for i, r := range runes {
		reversed[len(runes)-1-i] = r
	}
	
	// Trim spaces and return result
	return strings.TrimSpace(string(reversed))
} 
package main

import (
	"slices"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	sb := []byte(s)

	slices.Reverse(sb)

	s2 := string(sb)

	return s == s2
}

func main() {}

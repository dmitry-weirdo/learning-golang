package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	su := strings.ToUpper(s) // don't care about uppercasing anymore

	noDashes := strings.ReplaceAll(su, "-", "")

	//fmt.Printf("No dashes: %v \n", noDashes)

	// if we need the first group to be k
	mod := len(noDashes) % k

	i := 0

	var sb strings.Builder

	if mod != 0 {
		sb.WriteString(noDashes[:mod])
		sb.WriteString("-")
		i = mod // we start from a string that divides by k
	}

	for i < len(noDashes) {
		sb.WriteString(noDashes[i : i+k])
		sb.WriteString("-")
		i += k
	}

	withLastDash := sb.String()

	if len(withLastDash) < 1 { // corner case of just "-" character in the input
		return ""
	}

	return withLastDash[:len(withLastDash)-1]
}

func test(s string, k int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Initial license key: %v \n", s)
	fmt.Printf("K (split section size, except possibly the first section): %v \n", k)

	result := licenseKeyFormatting(s, k)

	fmt.Printf("Formatted license key: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "5F3Z-2e-9-w"
	k := 4
	expected := "5F3Z-2E9W"

	test(s, k, expected)
}

func test2() {
	s := "2-5g-3-J"
	k := 2
	expected := "2-5G-3J"

	test(s, k, expected)
}

func main() {
	// 482. License Key Formatting
	test1()
	test2()
}

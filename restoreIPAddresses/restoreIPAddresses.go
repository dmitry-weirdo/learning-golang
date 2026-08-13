package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	// todo: same code can be rewritten recursively
	return restoreIpAddresses_backtracking(s)
}

func restoreIpAddresses_backtracking(s string) []string {
	n := len(s)

	// In the IP address, we need at least 4 digits: "0.0.0.0"
	// and max 12 digits: "123.123.123.123"
	// Since we cannot add or remove digits, we need to use all digits in the string
	// -> if there are not enough or too many digits, the result is empty.
	if n < 4 || n > 12 {
		return []string{}
	}

	result := make([]string, 0)

	// todo: we can use expression as global slice as better memory (then we have to backtrack after each character attempt. But current solution is already good on memory
	// actually, not strings.Builder since it does not support removing characters from the end, just the complete reset to 0 characters.
	var dfs func(start int, dotsRemaining int, expression string)

	dfs = func(start int, dotsRemaining int, expression string) {
		//fmt.Println()
		//fmt.Printf("Start: %v, dotsRemaining: %v, expression: \"%v\" \n", start, dotsRemaining, expression)

		if (start >= n) && (dotsRemaining == -1) { // reached the end of the string and all dots used -> success, add to the result
			//fmt.Printf("Success! Adding \"%v\" to the result. \n", expression)
			result = append(result, expression)
			return
		}

		if start >= n {
			// reached the end but not all dots used -> fail
			//fmt.Printf("Reached the end of the string, but there are still %v dots remaining. Fail. \n", dotsRemaining)
			return
		}

		// fail if there are too many digits for the current remaining dots
		remainingDigits := n - start
		maxValidRemainingDigits := getMaximumValidRemainingDigits(dotsRemaining)

		if remainingDigits > maxValidRemainingDigits { // too many digits for the remaining dots
			//fmt.Printf("%v > %v digits remaining, and %v dots remaining. Fail. \n", maxValidRemainingDigits, remainingDigits, dotsRemaining)
			return
		}

		dot := "."
		if dotsRemaining == 0 {
			dot = "" // do not append after the last octet, we're just adding digits
		}

		// try to add 1 digit
		if isValidOctet(s, start, start) {
			dfs(start+1, dotsRemaining-1, expression+s[start:start+1]+dot)
		}

		// try to add 2 digits
		if isValidOctet(s, start, start+1) {
			dfs(start+2, dotsRemaining-1, expression+s[start:start+2]+dot)
		}

		// try to add 3 digits
		if isValidOctet(s, start, start+2) {
			dfs(start+3, dotsRemaining-1, expression+s[start:start+3]+dot)
		}
	}

	dfs(0, 3, "")

	return result
}

func getMaximumValidRemainingDigits(dotsRemaining int) int {
	switch dotsRemaining {

	// corner case -> all dots used, there must be 0 remaining digits
	case -1:
		return 0

	case 0:
		return 3

	case 1:
		return 6

	case 2:
		return 9

	case 3:
		return 12

	default:
		panic(fmt.Sprintf("Too many dots remaining %v .", dotsRemaining))
	}
}

func isValidOctet(s string, start, end int) bool { // [start:end] inclusive
	if end >= len(s) { // going over the end of the string is non-valid
		return false
	}

	substring := s[start : end+1]
	n := len(substring)

	if (n < 1) || (n > 4) { // octet must be 1 to 3 chars
		return false
	}

	if (substring[0] == '0') && (n > 1) { // 0, allowed, leading 0 not allowed
		return false
	}

	octetValue, err := strconv.Atoi(substring)

	if err != nil {
		fmt.Println("invalid number:", err) // todo: we can throw here
		return false
	}

	if octetValue > 255 { // values > 255 are not allowed
		return false
	}

	return true
}

func test(s string, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String to find all possible IPs: %v \n", s)

	result := restoreIpAddresses(s)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	s := "25525511135"
	expected := []string{"255.255.11.135", "255.255.111.35"}

	test(s, expected)
}

func test2() {
	s := "0000"
	expected := []string{"0.0.0.0"}

	test(s, expected)
}

func test3() {
	s := "101023"
	expected := []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}

	test(s, expected)
}

func test4() {
	s := "123" // not enough digits -> expect empty result
	expected := []string{}

	test(s, expected)
}

func test5() {
	s := "1234561234567" // too many digits -> expect empty result
	expected := []string{}

	test(s, expected)
}

func main() {
	// 93. Restore IP Addresses
	test1()
	test2()
	test3()
	test4()
	test5()
}

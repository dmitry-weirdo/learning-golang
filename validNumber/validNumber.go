package main

import "fmt"

func isNumber(s string) bool {
	// todo: implement the Deterministic Finite Automaton (DFA) solution

	return isNumber_bruteforce(s)
}

func isNumber_bruteforce(s string) bool {
	n := len(s)

	i := 0
	j := 0

	dotUsed := false
	exponentUsed := false

	// skip the leading sign
	if isSign(s[i]) {
		i++
	}

	if i >= n { // just single sign, nothing more -> fail
		return false
	}

	// decimal is allowed to start with dot .
	if isDot(s[i]) {
		dotUsed = true
		i++
	}

	j = skipDigits(s, i)
	if j <= i { // no digits skipped
		return false
	}

	i = j

	if i >= n { // reached the end of the string -> success!
		return true
	}

	// skip the dot after the digits
	if isDot(s[i]) {
		if dotUsed { // second dot -> fail
			return false
		}

		dotUsed = true
		i++
	}

	// skip digits
	if dotUsed {
		j = skipDigits(s, i)

		i = j

		if i >= n { // 123.45 -> ok
			return true
		}
	}

	// now what is left is just the exponent
	if isExponent(s[i]) {
		exponentUsed = true
		i++

		if i >= n { // string ends on e/E character -> fail
			return false
		}
	}

	if exponentUsed { // what remains is just the integer (sign and 1+ digits)
		// skip the optional leading sign
		if isSign(s[i]) {
			i++
		}

		// after the exponent, we must have >= 1 digits
		j = skipDigits(s, i)

		if j <= i { // no digits skipped -> fail
			return false
		}

		i = j

		if i >= n { // reached the end of the string
			return true
		}
	}

	return false
}

func skipDigits(s string, i int) int {
	for i < len(s) && isDigit(s[i]) {
		i++
	}

	return i
}

func isSign(b byte) bool {
	return b == '+' || b == '-'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isExponent(b byte) bool {
	return b == 'e' || b == 'E'
}

func isDot(b byte) bool {
	return b == '.'
}

func test(s string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := isNumber(s)

	fmt.Printf("String is a valid number: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	valid := []string{
		"0",
		"2",
		"0089",
		"-0.1",
		"+3.14",
		"4.",
		"-.9",
		"2e10",
		"-90E3",
		"3e+7",
		"+6e-1",
		"53.5e93",
		"-123.456e789",
	}

	for _, v := range valid {
		test(v, true)
	}
}

func test2() {
	invalid := []string{
		"+",
		"e",
		".",
		"abc",
		"1a",
		"1e",
		"e3",
		"99e2.5",
		"--6",
		"-+3",
		"95a54e53",
	}

	for _, v := range invalid {
		test(v, false)
	}
}

func main() {
	// 65. Valid Number
	test1()
	test2()
}

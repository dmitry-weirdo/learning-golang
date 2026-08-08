package main

import (
	"fmt"
	"strings"
)

func maskPII(s string) string {
	// is email when contains "@"
	if strings.Contains(s, "@") {
		return handleEmail(s)
	} else {
		return handlePhone(s)
	}
}

func handleEmail(s string) string {
	sl := strings.ToLower(s)

	atIndex := strings.Index(sl, "@")
	name := sl[:atIndex]
	domainWithAt := sl[atIndex:]

	// first letter + "*****" + lastLetter
	nameMasked := name[:1] + "*****" + name[len(name)-1:]

	return nameMasked + domainWithAt
}

func handlePhone(s string) string {
	// strings.Map
	numbersOnly := strings.Map(func(r rune) rune {
		if ('0' <= r) && (r <= '9') {
			return r
		}

		return -1
	}, s)

	lastFourDigits := numbersOnly[len(numbersOnly)-4:]

	countryCodeLen := len(numbersOnly) - 10

	switch countryCodeLen {
	case 0:
		return "***-***-" + lastFourDigits

	case 1:
		return "+*-***-***-" + lastFourDigits

	case 2:
		return "+**-***-***-" + lastFourDigits

	case 3:
		return "+***-***-***-" + lastFourDigits
	}

	panic(fmt.Sprintf("Illegal country code length: %v", countryCodeLen))
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := maskPII(s)

	fmt.Printf("String with hidden personal info: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "LeetCode@LeetCode.com"
	expected := "l*****e@leetcode.com"

	test(s, expected)
}

func test2() {
	s := "AB@qq.com"
	expected := "a*****b@qq.com"

	test(s, expected)
}

func test3() {
	s := "1(234)567-890"
	expected := "***-***-7890"

	test(s, expected)
}

func main() {
	// 831. Masking Personal Information
	test1()
	test2()
	test3()
}

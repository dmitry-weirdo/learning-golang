package main

import (
	"fmt"
	"strings"
)

func findKthBit(n int, k int) byte {
	return findKthBit_bruteforce(n, k)
}

func findKthBit_bruteforce(n int, k int) byte {
	s := "0" // this is s[1]

	// S[i+1] = S[i] + "1" + reverse(invert(S[i])) for i > 1

	var sb strings.Builder

	for i := 2; i <= n; i++ {
		sb.Reset()
		sb.WriteString(s)
		sb.WriteByte('1')

		invertedReversed := invertAndReverse(s)
		sb.WriteString(invertedReversed)

		s = sb.String()
	}

	//fmt.Printf("s[%v] = \"%v\" \n", n, s)

	return s[k-1] // k is i-based, omg
}

func invertAndReverse(s string) string {
	var sb strings.Builder

	for i := len(s) - 1; i >= 0; i-- { // we append reversed
		if s[i] == '0' {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}

	return sb.String()
}

func test(n, k int, expectedResult byte) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - to calculate S[n]: %v \n", n)
	fmt.Printf("K - return [K-1]-th character of S[%v]: %v \n", n, k)

	result := findKthBit(n, k)

	fmt.Printf("S[%v][%v]: %c \n", n, k-1, result)
	fmt.Printf("Expected result: %c \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %c, actual result = %c \n", expectedResult, result)
	}
}

func test1() {
	// s[3] = "0111001", k is 1-indexed
	test(3, 1, '0')
}

func test2() {
	// s[4] = "011100110110001", k is 1-indexed
	test(4, 11, '1')
}

func main() {
	// 1545. Find Kth Bit in Nth Binary String
	test1()
	test2()
}

package main

import (
	"fmt"
)

const ( // corner numbering in the response
	BOTTOM_RIGHT = 0
	TOP_RIGHT    = 1
	TOP_LEFT     = 2
)

func mirrorReflection(p int, q int) int {
	// Q <= P, but it doesn't matter

	// We need to count where m*p = n*q, i.e. the LCM(p, q)
	// This will be the first intersection with the corner.
	// We're searching just the heights,
	// i.e. when having an infinite height, when will n*q height (total height of the ray path)
	// hit a whole number of m*p.

	// Then we have 3 cases:

	// case 1:
	// m = odd, n = odd -> top-right
	// Example: 1 * p = 1 * q -> direct 1st hit of top-right

	// case 2:
	// m = odd, n = even -> top-left
	// Example: 1 * p = 2 * q -> 2nd hit hits top-left

	// case 3:
	// m = even, n = odd -> bottom-right
	// Example: 2 * p = 3 * q -> 4th hit hist the bottom right.

	// Explanation of case 3 (we assume that Q <= P):
	// 1st hit uses Q height
	// 2nd hit uses (P - Q) height and hits the top.
	// 3rd hit continues down, using Q - (P - Q) remaining count of the 2nd hit
	// i.e. 2 + 3 hits use Q height, 1 + 2 + 3 hits use 2 * Q height.
	// 4th hit uses Q and hits the bottom-right.
	// i.e. 1 + 2 + 3 + 4 hits used 3*Q height.
	// We went to the top of the room and back, i.e. total height used is 2 * P.

	// Note that m = even, q = even is impossible,
	// since then we can divide both by 2 without losing the divisibility and find the smaller LCM.
	// I.e. this won't be a LCM.
	pqLcm := lcm(p, q)

	m := pqLcm / p
	n := pqLcm / q

	mIsOdd := m%2 == 1
	nIsOdd := n%2 == 1

	if mIsOdd && nIsOdd {
		return TOP_RIGHT
	}

	if mIsOdd && !nIsOdd {
		return TOP_LEFT
	}

	if !mIsOdd && nIsOdd {
		return BOTTOM_RIGHT
	}

	// this must never happen
	errorMessage := fmt.Sprintf("p = %v, q = %v, LCM(p, q) = %v, m = %v, n = %v, mIsOdd = %v, nIsOdd = %v. This must never happen.",
		p, q, pqLcm, m, n, mIsOdd, nIsOdd)

	panic(errorMessage)
}

func lcm(a, b int) int {
	gcdOfAAndB := gcd(a, b)

	return (a / gcdOfAAndB) * b
}

func gcd(a, b int) int {
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func test(p, q int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("P - width of the mirror room: %v \n", p)
	fmt.Printf("Q - height of the first: %v \n", q)

	result := mirrorReflection(p, q)

	fmt.Printf("First corner where the ray will hit: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(2, 1, TOP_LEFT)
}

func test2() {
	test(3, 1, TOP_RIGHT)
}

func test3() {
	test(3, 2, BOTTOM_RIGHT)
}

func main() {
	// 858. Mirror Reflection
	test1()
	test2()
	test3()
}

package main

import "fmt"

func climbStairs(n int) int {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 1

	// since it's just a Fibonacci sequence, we can also solve it iteratively and use just 2 value instead of a map
	return rec(n, m)
}

func rec(n int, m map[int]int) int {
	r, ok := m[n]
	if ok { // already memoized
		return r
	}

	v := rec(n-2, m) + rec(n-1, m)
	m[n] = v
	return v
}

func recBad(n int) int { // will be 2^n, memory overflow
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	return recBad(n-1) + recBad(n-2)
}

func test(n int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("n: %v \n", n)

	result := climbStairs(n)

	fmt.Printf("Result: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 2
	expected := 2

	test(n, expected)
}

func test2() {
	n := 3
	expected := 3

	test(n, expected)
}

func main() {
	test1()
	test2()
}

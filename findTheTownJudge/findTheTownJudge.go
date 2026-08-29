package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	// numbers are from 1 to n -> we allocate (n + 1)
	in := make([]int, n+1)
	out := make([]int, n+1)

	for _, v := range trust {
		from := v[0]
		to := v[1]

		// we're guaranteed that all trust pairs are unique -> we count them without any uniqueness check
		out[from]++
		in[to]++
	}

	// We search for a node where:
	// - count(outbound) = 0
	// - count(inbound) = n - 1
	for i := 1; i <= n; i++ {
		if (out[i] == 0) && (in[i] == n-1) {
			return i
		}
	}

	return -1
}

func test(n int, m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - total people: %v \n", n)
	fmt.Printf("Trust edges: %v \n", m)

	result := findJudge(n, m)

	fmt.Printf("Judge index: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 2

	trust := [][]int{
		{1, 2},
	}

	expected := 2

	test(n, trust, expected)
}

func test2() {
	n := 3

	trust := [][]int{
		{1, 3},
		{2, 3},
	}

	expected := 3

	test(n, trust, expected)
}

func test3() {
	n := 3

	trust := [][]int{
		{1, 3},
		{2, 3},
		{3, 1},
	}

	expected := -1

	test(n, trust, expected)
}

func main() {
	// 997. Find the Town Judge
	test1()
	test2()
	test3()
}

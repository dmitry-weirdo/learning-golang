package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	var dfs func(leftCount, rightCount int, s string, n int)

	dfs = func(leftCount, rightCount int, s string, n int) {
		// Prune the further execution if:
		// leftCount > n
		// rightCount > n
		// leftCount < rightCount
		if leftCount > n || rightCount > n || leftCount < rightCount {
			return
		}

		if leftCount == n && rightCount == n {
			// reached the target -> add the current string to result and stop
			result = append(result, s)
			return
		}

		// add (
		dfs(leftCount+1, rightCount, s+"(", n)

		// add )
		dfs(leftCount, rightCount+1, s+")", n)
	}

	dfs(0, 0, "", n)

	return result
}

func test(n int, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Count of () pairs: %v \n", n)

	result := generateParenthesis(n)

	fmt.Printf("Possible pairs of %v pairs of () parentheses:\n%v\n", n, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

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
	n := 3

	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	test(n, expected)
}

func test2() {
	n := 1

	expected := []string{"()"}

	test(n, expected)
}

func main() {
	// 22. Generate Parentheses
	test1()
	test2()
}

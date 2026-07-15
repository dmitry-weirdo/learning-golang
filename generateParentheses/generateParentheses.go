package main

import (
	"fmt"
)

var resultGlobal []string

func generateParenthesis(n int) []string {
	resultGlobal = make([]string, 0)

	dfs(0, 0, "", n)

	return resultGlobal
}

func dfs(leftCount, rightCount int, s string, n int) {
	// Prune the further execution if:
	// leftCount > n
	// rightCount > n
	// leftCount < rightCount
	if leftCount > n || rightCount > n || leftCount < rightCount {
		return
	}

	if leftCount == n && rightCount == n {
		// reached the target -> add the current string to result and stop
		resultGlobal = append(resultGlobal, s)
		return
	}

	// add (
	dfs(leftCount+1, rightCount, s+"(", n)

	// add )
	dfs(leftCount, rightCount+1, s+")", n)
}

func test(n int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Count of () pairs: %v \n", n)

	result := generateParenthesis(n)

	fmt.Printf("Possible pairs of %v pairs of () parentheses:\n%v\n", n, result)
}

func main() {
	// 22. Generate Parentheses
	test(7)
}

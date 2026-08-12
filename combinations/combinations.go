package main

import "fmt"

func combine(n int, k int) [][]int {
	// iterative, filing digit-by-digit, always increasing
	// O(2^n) - at every position of N numbers we either select or not select it
	// works in around the same 30 ms range
	return combine_iterative(n, k)

	// backtracking, has a lot of recursion
	// O(2^n) - at every position of N numbers we either select or not select it
	// works in 29+ ms
	//return combine_backtracking(n, k)
}

func combine_iterative(n int, k int) [][]int {
	result := make([][]int, 0) // array of arrays of size K

	digits := make([]int, k) // current K digits, they will be filled always increasing

	i := 0 // index ocr current digit

	for i >= 0 { // until we finished all in the 0-th digit and went to -1
		// increase current digit (we're stating from digits[0] == 0)
		digits[i]++

		if digits[i] > n {
			// (overflow > N) on the current digit
			// -> we've iterated everything in ths is digit -> continue the previous digit
			i--
			continue
		}

		if i == k-1 {
			// we're iterating on the last digit, and this is the valid answer < N
			// -> add it to the result
			// !!! to not append selection itself, its elements will be changed
			result = append(result, copyArray(digits))

			continue
		}

		// current digit is ok, but not last -> continue with next digit = (currentDigit + 1), since we're building increasing
		digits[i+1] = digits[i]
		i++
	}

	return result
}

func combine_backtracking(n int, k int) [][]int {
	result := make([][]int, 0) // array of arrays of size K

	selection := make([]int, 0) // current selection for backtracking

	// We're calculating the result ascending.
	// For every number, we can either select it or not
	// If we selected K numbers, we add to the result and stop the current branch
	var dfs func(current int) // returns nothing , it will add to result itself

	dfs = func(current int) {
		//fmt.Printf("current: %d. Selection %d \n", current, selection)

		if len(selection) == k { // combination of K -> add to result, stop current branch iteration
			// !!! to not append selection itself, its elements will be changed
			result = append(result, copyArray(selection))
			return
		}

		if current > n { // we're over the last number from 1 to N -> this is a dead-end branch
			return
		}

		// case 1: select current number -> add the current number to the selection, continue with the next number
		selection = append(selection, current)
		dfs(current + 1)

		// backtrack - remove the added current number
		selection = selection[:len(selection)-1]

		// case 2: do NOT select the current number -> just continue with the next number, no changes
		dfs(current + 1)
	}

	dfs(1) // start with the smallest number

	// result will be collected by dfs execution
	return result
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func test(n, k int, expectedResult [][]int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - numbers from 1 to N: %v \n", n)
	fmt.Printf("K - count of numbers to select: %v \n", k)

	result := combine(n, k)

	fmt.Printf("All combinations of %v distinct values from numbers from 1 to %v: \n%v \n", k, n, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	n := 4
	k := 2

	expected := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}

	test(n, k, expected)
}

func test2() {
	n := 1
	k := 1

	expected := [][]int{
		{1},
	}

	test(n, k, expected)
}

func main() {
	// 77. Combinations
	test1()
	test2()
}

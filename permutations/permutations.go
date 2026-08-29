package main

import (
	"demo/matrixCommon"
	"fmt"
)

func permute(nums []int) [][]int {
	n := len(nums)

	m := make(map[int]bool) // used values, since all values are distinct // todo: we can use indices instead
	a := make([]int, n)

	result := make([][]int, 0)

	var dfs func(i int)

	dfs = func(i int) {
		if i >= n {
			// all array values used -> save the result
			result = append(result, copyArray(a))
			return
		}

		for _, v := range nums {
			if m[v] { // value already used -> do not
				continue
			}

			// select the value
			a[i] = v
			m[v] = true

			dfs(i + 1)

			// backtrack
			m[v] = false
		}
	}

	dfs(0)

	return result
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func test(arr []int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text

	result := permute(arr)

	fmt.Printf("All permutations: \n")
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Expected result: \n")
	matrixCommon.PrintIntMatrix(expectedResult)

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
	a := []int{1, 2, 3}

	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	test(a, expected)
}

func test2() {
	a := []int{0, 1}

	expected := [][]int{
		{0, 1},
		{1, 0},
	}

	test(a, expected)
}

func test3() {
	a := []int{1}

	expected := [][]int{
		{1},
	}

	test(a, expected)
}

func main() {
	// 46. Permutations
	test1()
	test2()
	test3()
}

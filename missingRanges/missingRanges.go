package main

import (
	"demo/matrixCommon"
	"fmt"
)

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	if len(nums) < 1 { // all values from [lower; upper] is missing
		return [][]int{{lower, upper}}
	}

	n := len(nums)

	result := make([][]int, 0)

	// add the start if required
	if lower < nums[0] {
		result = append(result, []int{lower, nums[0] - 1})
	}

	// add all in-between values
	for i := 0; i < n-1; i++ {
		if nums[i]+1 == nums[i+1] { // consecutive numbers -> no range in between them
			continue
		}

		result = append(result, []int{nums[i] + 1, nums[i+1] - 1})
	}

	// add the end if required
	if upper > nums[n-1] {
		result = append(result, []int{nums[n-1] + 1, upper})
	}

	return result
}

func test(arr []int, l, r int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Range: [%v; %v] \n", l, r)

	result := findMissingRanges(arr, l, r)

	fmt.Printf("Missing ranges: \n")
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
	a := []int{0, 1, 3, 50, 75}
	l := 0
	r := 99

	expected := [][]int{
		{2, 2},
		{4, 49},
		{51, 74},
		{76, 99},
	}

	test(a, l, r, expected)
}

func test2() {
	a := []int{-1}
	l := -1
	r := -1

	expected := [][]int{}

	test(a, l, r, expected)
}

func test3() {
	a := []int{2}
	l := 1
	r := 3

	expected := [][]int{
		{1, 1},
		{3, 3},
	}

	test(a, l, r, expected)
}

func test4() {
	a := []int{3}
	l := 1
	r := 5

	expected := [][]int{
		{1, 2},
		{4, 5},
	}

	test(a, l, r, expected)
}

func test5() {
	a := []int{1, 3}
	l := 1
	r := 3

	expected := [][]int{
		{2, 2},
	}

	test(a, l, r, expected)
}

func test6() {
	a := []int{}
	l := 1
	r := 3

	expected := [][]int{
		{1, 3},
	}

	test(a, l, r, expected)
}

func main() {
	// 163. Missing Ranges
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}

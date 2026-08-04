package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	return findDisappearedNumbers_optimized(nums)
}

func findDisappearedNumbers_optimized(nums []int) []int {
	// we're using the fact that all values are positive
	// array indexes are [0...n - 1]
	// expected values are [1...n]

	// If we find a value i, we negate the value [i - 1]
	// If it is already negative, then we know that i is the duplicate.
	// We leave a[i] as positive.
	// Therefore, all elements will be marked negative except the element of [j - 1], where j is the negative index
	for _, v := range nums {
		n := abs(v)

		if nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
		}
	}

	// a single [i] element that remained positive -> missing value is [i + 1]
	result := make([]int, 0)

	for i, v := range nums {
		if v > 0 { // we found an index that was not marked to negative value
			result = append(result, i+1)
		}
	}

	return result
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)

	result := findDisappearedNumbers(arr)

	fmt.Printf("Array after finding missing elements: %v \n", arr)
	fmt.Printf("Missing elements: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

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
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{5, 6}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1}
	expected := []int{2}

	test(arr, expected)
}

func main() {
	// 448. Find All Numbers Disappeared in an Array
	// This is the same solution as in "645. Set Mismatch", we're just returning all missing elements and not returning the duplicates
	test1()
	test2()
}

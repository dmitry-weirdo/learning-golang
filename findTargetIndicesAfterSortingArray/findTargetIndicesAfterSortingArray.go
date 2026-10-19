package main

import "fmt"

func targetIndices(nums []int, target int) []int {
	// count the values smaller than target
	// count the values equal to target
	// This will take O(n), no need to sort the array
	smallerCount := 0
	equalCount := 0

	for _, v := range nums {
		if v == target {
			equalCount++
		} else if v < target {
			smallerCount++
		}
	}

	if equalCount == 0 { // no target values -> return an empty array
		return []int{}
	}

	result := make([]int, equalCount)

	for i := range equalCount {
		result[i] = smallerCount + i
	}

	return result
}

func test(arr []int, target int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Value to search in the target array: %v \n", arr)

	result := targetIndices(arr, target)

	fmt.Printf("Indices of %v in the sorted array: %v \n", target, result)
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
	test(
		[]int{1, 2, 5, 2, 3},
		2,
		[]int{1, 2},
	)
}

func test2() {
	test(
		[]int{1, 2, 5, 2, 3},
		3,
		[]int{3},
	)
}

func test3() {
	test(
		[]int{1, 2, 5, 2, 3},
		5,
		[]int{4},
	)
}

func main() {
	// 2089. Find Target Indices After Sorting Array
	test1()
	test2()
	test3()
}

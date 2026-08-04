package main

import "fmt"

func findMissingElements(nums []int) []int {
	minValue := nums[0]
	maxValue := nums[0]

	m := make(map[int]bool)

	for _, v := range nums {
		minValue = min(minValue, v)
		maxValue = max(maxValue, v)
		m[v] = true
	}

	// we already know the size of the result
	resultSize := (maxValue - minValue + 1) - len(nums)
	result := make([]int, resultSize)

	j := 0 // write position in result

	for i := minValue; i <= maxValue; i++ {
		if _, ok := m[i]; !ok { // value is not present in the array -> add it to the result
			result[j] = i
			j++
		}
	}

	return result
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	n := len(arr) / 2

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("N (half of array length): %v \n", n)

	result := findMissingElements(arr)

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
	arr := []int{1, 4, 2, 5}
	expected := []int{3}

	test(arr, expected)
}

func test2() {
	arr := []int{7, 8, 6, 9}
	expected := []int{}

	test(arr, expected)
}

func test3() {
	arr := []int{5, 1}
	expected := []int{2, 3, 4}

	test(arr, expected)
}

func main() {
	// 3731. Find Missing Elements
	test1()
	test2()
	test3()
}

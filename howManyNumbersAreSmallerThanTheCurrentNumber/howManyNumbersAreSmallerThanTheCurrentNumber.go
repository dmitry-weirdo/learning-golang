package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	// since values are just 0...100, we can count the frequencies of every value
	freq := make([]int, 102) // +1 index for prefix sums

	for _, v := range nums {
		freq[v]++
	}

	// calculate the prefix sum of the frequencies
	// freq[i] = sum of count of values <= i
	for i := 1; i < len(freq); i++ {
		freq[i] += freq[i-1]
	}

	// todo: we can also replace in place in nums, but ok
	result := make([]int, len(nums))

	for i, v := range nums {
		if v == 0 {
			// we don't have a proper prefixSums[0], so handle this 0 separately
			result[i] = 0
		} else {
			// for value v, we need the count of values < v, i.e. values from 0 up to (v - 1) -> use freq[v-1]
			result[i] = freq[v-1]
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

	result := smallerNumbersThanCurrent(arr)

	fmt.Printf("Numbers count smaller than the current number: %v \n", result)
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
	arr := []int{8, 1, 2, 3, 3}
	expectedResult := []int{4, 0, 1, 1, 1}

	test(arr, expectedResult)
}

func test2() {
	arr := []int{6, 5, 4, 8}
	expectedResult := []int{2, 1, 0, 3}

	test(arr, expectedResult)
}

func test3() {
	arr := []int{7, 7, 7, 7}
	expectedResult := []int{0, 0, 0, 0}

	test(arr, expectedResult)
}

func test4() {
	arr := []int{0, 99, 100}
	expectedResult := []int{0, 1, 2}

	test(arr, expectedResult)
}

func main() {
	// 1365. How Many Numbers Are Smaller Than the Current Number
	test1()
	test2()
	test3()
	test4()
}

package main

import "fmt"

func productExceptSelf(nums []int) []int {
	// for a[i], we need
	// prefix product: a[0] * a[1] * ... * a[i - 1]
	// postfix product: a[i + 1] * a[i + 2] * ... * a[len - 1]

	// so we need to multiply prefix[i - 1] * postfix[i + 1]

	// actually, we don't need a prefix and postfix arrays, we can just multiply directly into the result array

	n := len(nums)
	result := make([]int, n)

	// go from 0 to n - 1, a[0] = 1, then multiply on a[i - 1]
	for i := range nums {
		if i == 0 {
			result[i] = 1
		} else {
			result[i] = result[i-1] * nums[i-1]
		}
	}

	// go down from n - 1 to 0, a[n - 1] = 1, then multiply on a[i + 1]
	// for postfix values, we just need 2 values - current and next
	nextPostfix := 1
	currentPostfix := 1

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			currentPostfix = 1
		} else {
			currentPostfix = nextPostfix * nums[i+1]
		}

		// multiply the prefixSum on the postfixSum
		result[i] = result[i] * currentPostfix

		nextPostfix = currentPostfix
	}

	return result
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)

	result := productExceptSelf(arr)

	fmt.Printf("Product array:   %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := 0; i < len(expectedResult); i++ {
		r := result[i]
		er := expectedResult[i]

		if r != er {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, er, i, r)
		}
	}
}

func test1() {
	arr := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	test(arr, expected)
}

func test2() {
	arr := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}

	test(arr, expected)
}

func main() {
	// 238. Product of Array Except Self
	test1()
	test2()
}

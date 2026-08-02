package main

import "fmt"

func maxPairStrength(nums []int) int64 {
	maxStrength := int64(1)
	var currentStrength int64
	var currentGcd int64

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			currentGcd = int64(gcd(nums[i], nums[j]))

			// strength = a * b / gcd(a, b)^2
			currentStrength = int64(nums[i]) * int64(nums[j]) / (currentGcd * currentGcd)

			maxStrength = max(maxStrength, currentStrength)
		}
	}

	return maxStrength
}

func gcd(a, b int) int {
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func test(arr []int, expectedResult int64) {
	fmt.Println()
	fmt.Printf("======================== \n")

	fmt.Printf("Array: %v \n", arr)

	result := maxPairStrength(arr)

	fmt.Printf("Best pair of array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{4, 6, 8}
	expected := int64(12)

	test(arr, expected)
}

func test2() {
	arr := []int{2, 3, 5}
	expected := int64(15)

	test(arr, expected)
}

func test3() {
	arr := []int{3, 3}
	expected := int64(1)

	test(arr, expected)
}

func testGcd(a, b, expectedResult int) {
	fmt.Println()
	fmt.Printf("======================== \n")

	fmt.Printf("A: %v \n", a)
	fmt.Printf("B: %v \n", b)

	result := gcd(a, b)

	fmt.Printf("GCD of %v and %v: %v \n", a, b, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func testGcd1() {
	testGcd(10, 15, 5)
}

func testGcd2() {
	testGcd(48, 18, 6)
}

func testGcd3() {
	testGcd(19, 29, 1)
}

func main() {
	// 4010. Maximize Pair Strength Using GCD
	testGcd1()
	testGcd2()
	testGcd3()

	test1()
	test2()
	test3()
}

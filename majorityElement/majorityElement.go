package main

import "fmt"

func majorityElement(nums []int) int {
	return majorityElement_BoyerMoore(nums)
}

func majorityElement_BoyerMoore(nums []int) int {
	// If we're guaranteed that there is an element with frequency > n / 2,
	// then we can use the "Boyer–Moore majority vote algorithm".
	// It is O(n) time and O(1) space.

	// todo: If the problem doesn't guarantee that a majority exists, use Boyer–Moore to get a candidate and then make a second pass to verify.

	// see https://www.youtube.com/watch?v=7pnhv842keE

	element := nums[0] // we're guaranteed that the array is non-empty
	count := 0

	for _, v := range nums {
		if count == 0 { // count reached 0 -> re-initialize with the current element and count = 1
			element = v
			count = 1
			continue
		}

		if v == element { // element stays the same -> increase count
			count++
		} else { // element is different -> decrease count
			count--
		}
	}

	return element
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := majorityElement(arr)

	fmt.Printf("Element that is a majority (appears more than N / 2 times): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{3, 2, 3}, 3)
}

func test2() {
	test([]int{2, 2, 1, 1, 1, 2, 2}, 2)
}

func test3() {
	test([]int{1}, 1)
}

func main() {
	// 169. Majority Element
	test1()
	test2()
	test3()
}

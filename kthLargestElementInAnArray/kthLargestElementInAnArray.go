package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	// todo: this implementations is failing with TLE on testcase 44/46 that is k = 50000, array = [1, 2, 3, 4, 5, 1 ..... 1, -5, -4, -3, -2, -1]

	// implementation with quick select algorithm
	// it is like QuickSort, but after partitioning, we do not handle both sides,
	// we're continuing only onto the part where the target index is

	// Average O(n) time if we assume that every partition clears about a half of the array:
	// n + n/2 + n/4 + ... + 1 = O(2n) = O(n)

	// In case of unlucky pivot (e.g. all elements in the array are equal),
	// it can degrade to O(n^2): n + (n - 1) + (n - 2) + ... + 1

	// k-th largest means index ( len(n) - k ) in the sorted array
	targetIndex := len(nums) - k

	return quickSelect(nums, targetIndex, 0, len(nums)-1)
}

func quickSelect(a []int, targetIndex int, left int, right int) int {
	// todo: we may select a random pivot instead
	pivotIndex := right
	pivot := a[pivotIndex]

	writePos := left

	for i := left; i < right; i++ { // excluding right - it's pivot itself
		// we're moving elements that are <= pivot to the left
		// it means that elements > pivot remain at the right
		if a[i] <= pivot {
			// value belongs to left -> write it to the position at the left
			// it can be swapped with itself
			a[i], a[writePos] = a[writePos], a[i]

			writePos++
		}
	}

	// to the left of writePos, all elements are now <= pivot
	// starting with writePos, all elements up to [right - 1] are > pivot
	// swap pivot with the first element that is > pivot
	a[writePos], a[pivotIndex] = a[pivotIndex], a[writePos]

	if targetIndex < writePos { // target value is to the left of the pivot
		return quickSelect(a, targetIndex, left, writePos-1)
	} else if targetIndex > writePos { // target value is to the right of the pivot
		return quickSelect(a, targetIndex, writePos+1, right)
	} else { // pivot is the target -> return it!
		return pivot
	}
}

func test(a []int, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", a)
	fmt.Printf("K-th largest to search, K = %v \n", k)

	result := findKthLargest(a, k)

	fmt.Printf("%v-th largest element: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expected := 5

	test(arr, k, expected)
}

func test2() {
	arr := []int{2, 2, 2, 2, 2, 2}
	k := 4
	expected := 2

	test(arr, k, expected)
}

func test3() {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	expected := 4

	test(arr, k, expected)
}

func main() {
	// 215. Kth Largest Element in an Array
	test1()
	test2()
	test3()
}

package main

import "fmt"

func reversePairs(nums []int) int {
	return reversePairs_mergeSort(nums) // O(n * log n)
	//return reversePairs_bruteForce(nums) // O(n^2) - fails TLE on n = 50000
}

func reversePairs_mergeSort(nums []int) int {
	// We're counting on each merge sort when we have 2 sorted arrays,
	// but we haven't yet merged them.

	// Genius of this solution that BEFORE MERGING,
	// all elements in left still have all their original indexes < all elements in right.

	// Since both arrays are sorted, if a[leftIndex] > 2 * a[rightIndex],
	// then all further elements in left (i > leftIndex) also satisfy this condition.
	// I.e. when we found this condition, we add [leftEnd - leftIndex + 1] to result,
	// and then continue with leftIndex++ and rightIndex++, without iterating from [rightIndex: rightEnd].

	// So all the compare operations should be O(n) in sum, it's the same O(n) than used for the merging.
	// (same iteration of len(left) + len(right) on every level of (log n) levels).
	// Basically, we increased from O(n * log n) to O(2 * n * log n), but is it the same complexity.

	count := 0

	// temp array for merging in the fixed memory area
	tempArray := make([]int, len(nums))

	// todo: we can also return count from the function
	var merge func(left, right int) // we're doing in place replacement

	merge = func(left, right int) {
		// base case -> 1 or fewer elements in the array -> do nothing
		if left >= right {
			return
		}

		// Will be more on left in case of even length.
		// len = 5, left = 0, right = 4 -> mid = 2, left = [0; 2], right = [3; 4]
		// len = 4, left = 0, right = 3 -> mid = 1, left = [0; 1], right = [2; 3]
		mid := (left + right) / 2

		// left array is [left : mid]
		// right array is [mid + 1 : right]
		merge(left, mid)
		merge(mid+1, right)

		// count target pairs
		i := left    // index in the left array
		j := mid + 1 // index in the right array

		// !!! we know that both left and right arrays are already sorted
		for (i <= mid) && (j <= right) {
			// if we reached the end of right (before ending left), we already added the sums for all the values on right
			// if we reached the end of left (before ending right), then all the elements on the left are smaller than 2*nums[j]

			if nums[i] <= 2*nums[j] {
				// left element is smaller than the current right
				// -> it will be also smaller than all the further elements on right (since the right array is sorted)
				// -> go to the next element on the left
				i++
			} else {
				// target found -> here is the magic trick

				// we're counting for all elements on the left in [i; mid] (since the left array is sorted)
				// !!! Thus, we have handled right[j], and we proceed to the next element of right
				count += mid - i + 1

				j++
			}
		}

		// merge arrays
		tempArrayIndex := 0
		i = left
		j = mid + 1

		for (i <= mid) && (j <= right) { // both arrays not empty
			if nums[i] < nums[j] { // take from left
				tempArray[tempArrayIndex] = nums[i]
				i++
				tempArrayIndex++
			} else { // take from right
				tempArray[tempArrayIndex] = nums[j]
				j++
				tempArrayIndex++
			}
		}

		// append the rest from the left array (if any)
		for i <= mid {
			tempArray[tempArrayIndex] = nums[i]
			i++
			tempArrayIndex++
		}

		for j <= right {
			tempArray[tempArrayIndex] = nums[j]
			j++
			tempArrayIndex++
		}

		//fmt.Printf("Left array %v and right array %v merged into temp array %v \n", nums[left:mid+1], nums[mid+1:right+1], tempArray[:tempArrayIndex])

		// copy from tempArray to original array
		copy(nums[left:], tempArray[:tempArrayIndex])
	}

	merge(0, len(nums)-1)

	return count
}

func reversePairs_bruteForce(nums []int) int {
	// brute-force - O(n^2) - TLE on test-case 36/140 (array of len 50000)
	n := len(nums)

	fmt.Printf("n: %d \n", n)

	count := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > 2*nums[j] {
				count++
			}
		}
	}

	return count
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := reversePairs(arr)

	fmt.Printf("Array after merge sort: %v \n", arr)

	fmt.Printf("Count of pairs i < j where ( a[i] > 2 * a[j] ): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 3, 2, 3, 1}
	expected := 2 // [3, 1] and [3, 1]

	test(arr, expected)
}

func test2() {
	arr := []int{2, 4, 3, 5, 1}
	expected := 3 // [4, 1], [3, 1], [5, 1]

	test(arr, expected)
}

func main() {
	// 493. Reverse Pairs
	test1()
	test2()
}

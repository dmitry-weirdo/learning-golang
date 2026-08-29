package main

import "fmt"

func findDuplicate(nums []int) int {
	// O(n) time
	// O(n) space - to store the array
	// passes in 0-3 ms -> faster than using a hash map
	return findDuplicate_hashMapAsArray(nums)

	// O(n) time
	// O(n) space - to store the hash map
	// Passes in 25-26 ms - Go's hash map is slow.
	//return findDuplicate_hashMap(nums)

	// O(2 * n) time
	// O(1) space
	// So should be slower than the HashMap solution that is purely O(n)
	// Passes in 8-9 ms
	//return findDuplicate_floydAlgorithm(nums)
}

func findDuplicate_hashMapAsArray(nums []int) int {
	m := make([]bool, len(nums))

	for _, v := range nums {
		if m[v] {
			return v
		}

		m[v] = true
	}

	panic("This must never happen")
}

func findDuplicate_hashMap(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		if m[v] {
			return v
		}

		m[v] = true
	}

	panic("This must never happen")
}

func findDuplicate_floydAlgorithm(nums []int) int {
	// we're using a[i] value as the Next index in the linked list
	// So if we went into the same index (a[i] == a[j]), this is a cycle.
	// Therefore, we can just execute the Floyd algorithm to find the cycle point.

	// It's the same as "142. Linked List Cycle II", just iterating over an array.

	slow := nums[nums[0]]       // slow.Next
	fast := nums[nums[nums[0]]] // fast.Next.Next

	for slow != fast {
		slow = nums[slow]       // slow.Next
		fast = nums[nums[fast]] // fast.Next.Next
	}

	// cycle detected -> now we have to find the point where the cycle starts
	slow = nums[0]

	for slow != fast { // now we're moving +1 on both slow and next
		slow = nums[slow] // slow.Next
		fast = nums[fast] // fast.Next
	}

	return slow
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := findDuplicate(arr)

	fmt.Printf("Duplicate element: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 3, 4, 2, 2}, 2)
}

func test2() {
	test([]int{3, 1, 3, 4, 2}, 3)
}

func test3() {
	test([]int{3, 3, 3, 3, 3}, 3)
}

func test4() {
	test([]int{1, 1}, 1)
}

func main() {
	// 287. Find the Duplicate Number
	test1()
	test2()
	test3()
	test4()
}

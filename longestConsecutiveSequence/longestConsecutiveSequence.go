package main

import "fmt"

func longestConsecutive(nums []int) int {
	// todo: implement with Union-Find, will it be faster?	//return longestConsecutive_unionFind(nums)

	// using just a single map
	// O(n) but much faster
	// passes in around 35-45 ms
	return longestConsecutive_optimized(nums)

	// O(n) since we're just run once via array and do potentially 2 merges and 1 add to maps.
	// All map operations are O(1).
	// Passes in 90+ ms
	//return longestConsecutive_intervalsMerge(nums)
}

type Interval struct {
	left  int
	right int
}

func (interval *Interval) Size() int {
	return interval.right - interval.left + 1
}

func longestConsecutive_optimized(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	// collect all values to a map for quick access
	// empty struct{} is smaller than boolean
	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		m[v] = struct{}{}
	}

	maxConsecutiveCount := 1

	// iterate all unique values
	for v := range m {
		// if the value is not the first in the sequence -> skip it
		if _, vMinusOneExists := m[v-1]; vMinusOneExists {
			continue
		}

		// value is the smallest in the sequence -> count the sequence starting from this value
		length := 1

		var exists bool
		_, exists = m[v+length]
		for exists {
			length++
			_, exists = m[v+length]
		}

		maxConsecutiveCount = max(maxConsecutiveCount, length)
	}

	return maxConsecutiveCount
}

func longestConsecutive_intervalsMerge(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	//n := len(nums)

	left := make(map[int]*Interval)  // interval.left to interval
	right := make(map[int]*Interval) // interval.right to interval

	// do not handle duplicate values
	m := make(map[int]bool)

	maxIntervalSize := 1

	for _, v := range nums {
		if m[v] { // value already handled -> don't try to merge it again
			continue
		}

		m[v] = true // mark current element as handled

		fmt.Println()
		fmt.Printf("Handling value %v \n", v)
		//fmt.Printf("Left map: %v \n", left)
		//fmt.Printf("Right map: %v \n", right)

		current := &Interval{v, v}

		mergedFromLeft := false
		mergedFromRight := false

		// value + 1, we search left(v + 1) to append from left
		potentialLeft := v + 1
		if interval, ok := left[potentialLeft]; ok {
			// append from left -> change left
			interval.left = current.left // join complete current interval at left
			maxIntervalSize = max(maxIntervalSize, interval.Size())

			fmt.Printf("Updated interval [%v; %v] to [%v; %v]. \n", potentialLeft, interval.right, interval.left, interval.right)

			current = interval

			// change the merged interval in left map
			delete(left, potentialLeft)
			left[interval.left] = interval // this will update the same interval also in the right map since it's a pointer

			mergedFromLeft = true
		}

		// value - 1, we search right(v - 1) to append from left
		potentialRight := v - 1
		if interval, ok := right[potentialRight]; ok {
			// append from right -> change right

			interval.right = current.right // join complete current interval at right
			maxIntervalSize = max(maxIntervalSize, interval.Size())

			fmt.Printf("Updated interval [%v; %v] to [%v; %v]. \n", interval.left, potentialRight, interval.left, interval.right)

			current = interval

			// change the merged interval in right map
			delete(right, potentialRight)
			right[interval.right] = interval // this will update the same interval also in the left map since it's a pointer

			mergedFromRight = true
		}

		if !mergedFromLeft && !mergedFromRight { // not merged -> add interval just of the current number
			if _, ok := left[v]; !ok {
				left[current.left] = current
				right[current.right] = current

				fmt.Printf("Added new interval [%v; %v] to the maps. \n", current.left, current.right)
			}
		}
	}

	return maxIntervalSize
}

func longestConsecutive_unionFind(nums []int) int {
	n := len(nums)

	// Into union-find, we put the indexes in the array
	// Union-find controls the transitive merges.

	// Example: 0, 1, 3, 2.
	// [0, 1] and [3] separate.
	// When we join 2 with 1 ( = 2 - 1)
	// Then we join 2 with 3  ( = 2 + 1)
	// Therefore, 0, 1 and 4 are also joint.
	// ! Especially note the 0 join that is not a neighbour of 2.

	// todo: implement
	return n
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := longestConsecutive(arr)

	fmt.Printf("Length of longest consecutive sequence of any elements: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{100, 4, 200, 1, 3, 2},
		4, // 1, 2, 3, 4
	)
}

func test2() {
	test(
		[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		9, // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	)
}

func test3() {
	test(
		[]int{1, 0, 1, 2},
		3, // 0, 1, 2
	)
}

func test4() {
	// -9,-9,-8,-8,-5,-5,-4,-4,-4,-3,-3,-2,-1,-1,-1,0,0,1,1,2,2,2,4,4,4,5,5,6,6
	test(
		[]int{4, 2, 2, -4, 0, -2, 4, -3, -4, -4, -5, 1, 4, -9, 5, 0, 6, -8, -1, -3, 6, 5, -8, -1, -5, -1, 2, -9, 1},
		8, // -5, -4, -3, -2, -1, 0, 1, 2
	)
}

func main() {
	// 128. Longest Consecutive Sequence
	test1()
	test2()
	test3()
	test4()
}

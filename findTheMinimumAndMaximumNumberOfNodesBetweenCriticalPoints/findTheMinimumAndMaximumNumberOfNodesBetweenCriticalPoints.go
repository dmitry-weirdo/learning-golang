package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func nodesBetweenCriticalPoints(head *ListNode) []int {
	prev := head
	curr := prev.Next
	next := curr.Next

	index := 1
	prevMinMaxIndex := -1

	firstIndex := -1
	lastIndex := -1

	const TOO_BIG_DIFF = 1_000_000_000 // there are <= 2 * 10^5 nodes in the list
	minDiff := TOO_BIG_DIFF

	for next != nil {
		isLocalMin := (curr.Val < prev.Val) && (curr.Val < next.Val)
		isLocalMax := (curr.Val > prev.Val) && (curr.Val > next.Val)

		if isLocalMin || isLocalMax {
			if firstIndex == -1 {
				firstIndex = index
			}

			if prevMinMaxIndex != -1 {
				minDiff = min(minDiff, index-prevMinMaxIndex)
			}

			prevMinMaxIndex = index
			lastIndex = index
		}

		// move to the next node
		prev = prev.Next
		curr = curr.Next
		next = next.Next
		index++
	}

	if firstIndex == lastIndex { // 0 or 1 min/max nodes
		return []int{-1, -1}
	}

	return []int{minDiff, lastIndex - firstIndex}
}

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := nodesBetweenCriticalPoints(list)

	fmt.Printf("Min and max distances between local minimums/maximums: %v \n", result)
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
		[]int{3, 1},
		[]int{-1, -1},
	)
}

func test2() {
	test(
		[]int{5, 3, 1, 2, 5, 1, 2},
		[]int{1, 3}, // 5-1, 1-1
	)
}

func test3() {
	test(
		[]int{1, 3, 2, 2, 3, 2, 2, 2, 7},
		[]int{3, 3}, // 3-3
	)
}

func main() {
	// 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points
	test1()
	test2()
	test3()
}

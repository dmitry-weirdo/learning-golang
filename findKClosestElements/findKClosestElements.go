package main

import (
	"fmt"
)

func findClosestElements(arr []int, k int, x int) []int {
	// edge-case -> len(arr) = k
	// -> just return the array
	if k >= len(arr) {
		return arr
	}

	// We need to find a consecutive window of size k
	left := 0
	right := len(arr) - k // we should have k elements at the end

	//bestWindowStart := -1

	// we're searching for the leftmost window since the minimum values are preferred in case of same distance
	for left < right { // todo: why <=, not <?
		// <= to handle left = 0, right = 2 -> mid = 1 -> right = 0 -> next iteration: left = 0, right = 0
		// left <= right if we're using right = mid - 1
		// left < right if we're using right = mid

		// we still need left <= right to evaluate the end of the array
		// left = 3, right = 5, mid = 4, move right -> left = mid + 1 = 5
		// we still need to handle left = 5, right = 5 -> it is a valid array

		mid := (left + right) / 2
		fmt.Println()
		fmt.Printf("left: %v, right %v, mid: %v \n", left, right, mid)

		// Now we check whether window [mid; mid + k - 1]
		// is a better window than a window 1 element to the right [mid + 1; mid + k].
		// The only different elements between these 2 windows are a[mid] and a[mid + k].
		// So if a[mid + k] is strictly better than a[mid],
		// then [mid; mid + k - 1] is not the best window, and we're moving right.
		// Else we're moving left, storing the current window as the first result that is better
		// than the windows to the right
		firstIndexAfterWindow := mid + k

		moveRight := false

		if firstIndexAfterWindow < len(arr) {
			// only check moving to right if there are still elements there
			startWindowElement := arr[mid]
			startWindowElementDistance := dist(startWindowElement, x)

			// magic heuristic that will move right in case of values at start and after the window are the same, and x is greater than this value
			//startWindowElementDistance := x - startWindowElement

			firstElementAfterWindow := arr[firstIndexAfterWindow]
			firstElementAfterWindowDistance := dist(firstElementAfterWindow, x)

			// magic heuristic that will move right in case of values at start and after the window are the same, and x is greater than this value
			//firstElementAfterWindowDistance := firstElementAfterWindow - x

			fmt.Printf("Start window distance (from a[%v] = %v to %v): %v \n", mid, startWindowElement, x, startWindowElementDistance)
			fmt.Printf("First after window distance (from a[%v] = %v to %v): %v \n", firstIndexAfterWindow, firstElementAfterWindow, x, firstElementAfterWindowDistance)

			moveRight = (startWindowElementDistance > firstElementAfterWindowDistance) || (startWindowElement == firstElementAfterWindow && startWindowElement < x)
		}

		if moveRight {
			// window has to be moved right -> it's definitely not the best window
			left = mid + 1

			fmt.Printf("Moving left to mid + 1 = %v. \n", left)
		} else {
			//bestWindowStart = mid // no moving to the right, keep this as current result

			// move left
			right = mid

			fmt.Printf("Moving right to mid - 1 = %v. \n", right)
		}
	}

	return arr[left : left+k]
}

func dist(val int, target int) int {
	return abs(val - target)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func test(arr []int, k, x int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("k (length of the subarray): %v \n", k)
	fmt.Printf("x (target value): %v \n", x)

	result := findClosestElements(arr, k, x)

	fmt.Printf("Closest %v elements to value %v: %v \n", k, x, result)
	//fmt.Printf("Expected LCA of elements p = %v and q = %v: %v \n", p, q, expectedResult)
	//
	//if result != expectedResult {
	//	fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result.Val)
	//}
}

func test1() {
	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3

	test(arr, k, x)
}

func test2() {
	arr := []int{1, 1, 2, 3, 4, 5}
	k := 4
	x := -1

	test(arr, k, x)
}

func test3() {
	arr := []int{1}
	k := 1
	x := 1

	test(arr, k, x)
}

func test4() {
	// this test-case is failing
	// if we use "while left < right + right = mid"
	// instead of "while left <= right + right = mid - 1
	arr := []int{0, 1, 2, 2, 2, 3, 6, 8, 8, 9}
	k := 5
	x := 9

	test(arr, k, x)
}

func test5() {
	arr := []int{1, 1, 2, 2, 2, 2, 2, 3, 3}
	k := 3
	x := 3

	test(arr, k, x)
}

func main() {
	// 658. Find K Closest Elements
	//test1()
	//test2()
	//test3()
	test4()
	//test5()
}

package main

import "fmt"

func leftMostBinarySearch(a []int, v int) int { // returns index
	left := 0
	right := len(a) // (len -1) fails on finding 0-th of array of 1 element

	for left < right {
		mid := (left + right) / 2

		if a[mid] < v {
			left = mid + 1
		} else { // if a[i] == mid, we will move left -> to the leftmost value
			right = mid
		}
	}

	return left
}

func rightMostBinarySearch(a []int, v int) int { // returns index
	left := 0
	right := len(a) // (len -1) fails on finding 0-th of array of 1 element

	for left < right {
		mid := (left + right) / 2

		//fmt.Printf("left: %v, right: %v, mid: %v, a[mid] = %v \n", left, right, mid, a[mid])

		if a[mid] <= v {
			left = mid + 1 // this will jump to the next value after v
		} else {
			right = mid
		}
	}

	// we're now at the left-most of the NEXT value
	// if the previous element is the target -> return it
	previousIndex := left - 1

	//fmt.Printf("left: %v, previous index: %v \n", left, previousIndex)

	if (previousIndex >= 0) && (a[previousIndex] == v) {
		return previousIndex
	}

	return left
}

func rightMostBinarySearch2(a []int, v int) int { // returns index
	left := 0
	right := len(a) // todo: is it? here we MUST set len and NOT (len - 1), else the array of 1 element will not work

	result := -1

	for left < right {
		mid := (left + right) / 2

		//fmt.Printf("left: %v, right: %v, mid: %v, a[mid] = %v \n", left, right, mid, a[mid])
		if a[mid] == v {
			result = mid // remember the latest position, it will be the right-most of the target value

			left = mid + 1 // this will jump to the next value after v
		} else if a[mid] < v {
			left = mid + 1 // this will jump to the next value after v
		} else {
			right = mid
		}
	}

	return result
}

func test(a []int, v int, expectedLeftMost int, expectedRightMost int) {
	fmt.Println()
	fmt.Println("==========================")

	fmt.Printf("Array: \n%v \n", a)

	leftOne := leftMostBinarySearch(a, v)
	rightOne := rightMostBinarySearch(a, v)
	rightOne2 := rightMostBinarySearch2(a, v)

	fmt.Printf("Left-most position of %v = [%v] \n", v, leftOne)
	fmt.Printf("(Alg. 1) Right-most position of %v = [%v] \n", v, rightOne)
	fmt.Printf("(Alg. 2) Right-most position of %v = [%v] \n", v, rightOne2)
	fmt.Printf("Expected left-most position of %v = [%v] \n", v, expectedLeftMost)
	fmt.Printf("Expected right-most position of %v = [%v] \n", v, expectedRightMost)

	// write failures if any

	// check left-most alg
	if expectedLeftMost == -1 {
		if (leftOne != -1) && (leftOne != len(a)) { // can return both -1 and len(arr)
			fmt.Printf("FAILURE: expected left-most = %v, actual left-most = %v \n", expectedLeftMost, leftOne)
		}
	} else {
		if leftOne != expectedLeftMost {
			fmt.Printf("FAILURE: expected left-most = %v, actual left-most = %v \n", expectedLeftMost, leftOne)
		}
	}

	// check right-most alg 1
	if expectedRightMost == -1 {
		if (rightOne != -1) && (rightOne != len(a)) { // can return both -1 and len(arr)
			fmt.Printf("FAILURE: (right-most Alg. 1) expected right-most = %v, actual right-most = %v \n", expectedRightMost, rightOne)
		}
	} else {
		if rightOne != expectedRightMost { // can return both -1 and len(1)
			fmt.Printf("FAILURE: (right-most Alg. 1) expected right-most = %v, actual right-most = %v \n", expectedRightMost, rightOne)
		}
	}

	// check right-most alg 2
	if expectedRightMost == -1 {
		if (rightOne2 != -1) && (rightOne2 != len(a)) { // can return both -1 and len(arr)
			fmt.Printf("FAILURE: (right-most Alg. 2) expected right-most = %v, actual right-most = %v \n", expectedRightMost, rightOne2)
		}
	} else {
		if rightOne2 != expectedRightMost { // can return both -1 and len(1)
			fmt.Printf("FAILURE: (right-most Alg. 2) expected right-most = %v, actual right-most = %v \n", expectedRightMost, rightOne2)
		}
	}
}

func search(nums []int, target int) int {
	// 704. Binary Search
	// this is the default solution of finding some of the indexes (not leftmost)
	// works even for arrays of size 1 (since `left <= right` clause).

	// If we do the leftmost option, array won't execute `left < right` for array size 1,
	// , and we have to check after the cycle whether arr[left] == target.

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func test1() {
	arr := []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4}

	test(arr, 1, 0, 3)
	test(arr, 2, 4, 6)
	test(arr, 3, 7, 8)
	test(arr, 4, 9, 9)
	test(arr, 5, -1, -1)
}

func test2() {
	arr := []int{1}

	test(arr, 1, 0, 0)
	test(arr, 2, -1, -1)
}

func test3() {
	arr := []int{}

	test(arr, 1, -1, -1)
	test(arr, 2, -1, -1)
}

func main() {
	test1()
	test2()
	test3()
}

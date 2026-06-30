package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // arr[right] is inclusive

	// arr[left:right] - right is NOT inclusive -> we need to do [left:right + 1]

	//for (left <= right) && (left < len(nums)) {
	for left <= right { // since we're starting with (right = len - 1), we cannot go over the size of the array
		mid := (left + right) / 2

		fmt.Println()
		fmt.Printf("left: %v, right: %v, mid: %v \n", left, right, mid)
		fmt.Printf("a[left:right] = a[%v:%v] = %v \n", left, right, nums[left:right+1])
		fmt.Printf("a[mid] = a[%v] = %v \n", mid, nums[mid])

		if nums[mid] == target { // found the target -> return it
			fmt.Printf("a[mid] = a[%v] = targetValue = %v. Returning mid = %d.\n", mid, target, mid)

			return mid
		}

		leftPartIsSorted := nums[left] <= nums[mid] // <= - equal included to work for the size 1 (if left == mid)

		if leftPartIsSorted {
			fmt.Printf("Left part of the array a[%v:%v] = %v is sorted. \n", left, mid, nums[left:mid+1])

			// we can check for target presence in the left part
			if (nums[left] <= target) && (target < nums[mid]) {
				right = mid - 1

				fmt.Printf("Target value %v is in the left sorted part of the array. Moving right to %v. \n", target, right)
			} else {
				left = mid + 1

				fmt.Printf("Target value %v is NOT in the left sorted part of the array. Moving left to %v. \n", target, left)

			}
		} else {
			fmt.Printf("Right part of the array a[%v:%v] = %v is sorted. \n", mid, right, nums[mid:right+1])

			// we can check for target presence in the right part
			if (nums[mid] < target) && (target <= nums[right]) {
				left = mid + 1

				fmt.Printf("Target value %v is in the right sorted part of the array. Moving left to %v. \n", target, left)
			} else {
				right = mid - 1

				fmt.Printf("Target value %v is NOT in the right sorted part of the array. Moving right to %v. \n", target, right)
			}
		}
	}

	return -1
}

func test(arr []int, target int, expected int) {
	fmt.Println()
	fmt.Println("=============================")

	fmt.Printf("Possibly rotated array: %v\n", arr)
	fmt.Printf("Target value: %v\n", target)

	result := search(arr, target)

	fmt.Printf("Result of searching for %v: %v \n", target, result)
	fmt.Printf("Expected result: %v \n", expected)

	if result != expected {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expected, result)
	}
}

func test1() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4

	test(arr, target, expected)
}

func test2() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	expected := -1

	test(arr, target, expected)
}

func test3() {
	arr := []int{1}
	target := 0
	expected := -1

	test(arr, target, expected)
}

func test4() {
	arr := []int{3, 1}
	target := 0
	expected := -1

	test(arr, target, expected)
}

func test5() {
	arr := []int{3, 1}
	target := 2
	expected := -1

	test(arr, target, expected)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}

package main

import (
	"fmt"
	"slices"
)

type Interval struct {
	left  int
	right int
	y     int
}

func visibleMountains(peaks [][]int) int {
	n := len(peaks)

	intervals := make([]Interval, n)

	// calculate intervals
	for i, peak := range peaks {
		intervals[i] = Interval{
			peak[0] - peak[1], // left = x - y
			peak[0] + peak[1], // right = x + y
			peak[1],
		}
	}

	// sort intervals by left asc, right desc
	// therefore, for the same left values, earlier interval will be the parent for merging
	// (we will merge right into left if i < j, left[i] <= left[j], right[i] >= right[j]).

	// special case: left[i] = left[j], right[i] = right[j] -> we remove both of these intervals

	slices.SortFunc(intervals, func(i, j Interval) int {
		if i.left == j.left { // for same left, sort by right desc
			return j.right - i.right
		}

		// for different left, sort by left desc
		return i.left - j.left
	})

	fmt.Printf("Sorted intervals: %v \n", intervals)

	// now we're sure that left values are non-decreasing,
	// i.e. i < j -> left[i] <= left[j]

	// So we're comparing only rights:
	// if right[i] > right[j] -> we skip mountain[j], it's hidden by mountain[j]
	// if right[i] < right[j] -> we switch currentRight to right[j] and count++
	// if right[i] == right[j] -> we skip both mountain[i] and mountain[j]

	currentRight := intervals[0].left // make sure first mountain will be counted

	visibleCount := 0

	for i, v := range intervals {
		if v.right <= currentRight {
			// i-th mountain is hidden by the current mountain -> skip it
			continue
		}

		// now we know that (v.right > currentRight) -> it's the new interval increased to the right
		currentRight = v.right

		isSameAsNext := (i < n-1) &&
			(v.left == intervals[i+1].left) &&
			(v.right == intervals[i+1].right)

		if !isSameAsNext { // do not count the new interval if it's equal to the next mountain
			visibleCount++
		}
	}

	return visibleCount
}

func test(arr [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Peek coordinates: %v \n", arr)

	result := visibleMountains(arr)

	fmt.Printf("Indexes to next greater elements: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := [][]int{
		{2, 2},
		{6, 3},
		{5, 4},
	}

	expected := 2

	test(arr, expected)
}

func test2() {
	arr := [][]int{
		{1, 3},
		{1, 3},
	}

	expected := 0 // equal peaks hide each other

	test(arr, expected)
}

func main() {
	// 2345. Finding the Number of Visible Mountains
	test1()
	test2()
}

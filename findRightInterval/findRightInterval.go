package main

import (
	"fmt"
	"slices"
)

type Interval struct {
	start int
	end   int
	index int // index in the original array
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)

	intervalsWithIndexes := make([]Interval, n)

	// build intervals
	for i, v := range intervals {
		interval := Interval{v[0], v[1], i}
		intervalsWithIndexes[i] = interval
	}

	// we sort by start, end, index. Actually, this must be by start since it's given that all start values are unique
	slices.SortFunc(intervalsWithIndexes, func(a, b Interval) int {
		if a.start != b.start { // this should be true in this problem (start values must be unique)
			return a.start - b.start
		}

		if a.end != b.end {
			return a.end - b.end
		}

		return a.index - b.index
	})

	//fmt.Printf("Intervals array sorted by start: %v \n", intervalsWithIndexes)

	lastIntervalStart := intervalsWithIndexes[n-1].start

	result := make([]int, n)

	for i, v := range intervals {
		endI := v[1]

		if lastIntervalStart < endI { // small optimization - if there is no such interval, no need for a binary search
			result[i] = -1
			continue
		}

		// search for first intervals[j] with start[j] >= end[i]
		j := binarySearchFirstWithStartGreaterThan(intervalsWithIndexes, endI)

		if j == -1 { // no interval found
			result[i] = -1
		} else { // interval found -> add its index in the original array to result
			result[i] = intervalsWithIndexes[j].index
		}
	}

	return result
}

func binarySearchFirstWithStartGreaterThan(arr []Interval, target int) int { // returns index in intervals
	left := 0
	right := len(arr) // it can be after the end of the array // !!! this will fail the test if we set len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if arr[mid].start >= target { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	if left >= len(arr) { // if it is after the array, interval is not found
		return -1
	}

	return left
}

func test(arr [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Intervals: %v \n", arr)

	result := findRightInterval(arr)

	fmt.Printf("Right interval indexes: %v \n", result)
	fmt.Printf("Expected result:        %v \n", expectedResult)

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
	intervals := [][]int{
		{1, 2},
	}

	expected := []int{-1}

	test(intervals, expected)
}

func test2() {
	intervals := [][]int{
		{3, 4},
		{2, 3},
		{1, 2},
	}

	expected := []int{-1, 0, 1}

	test(intervals, expected)
}

func test3() {
	intervals := [][]int{
		{1, 4},
		{2, 3},
		{3, 4},
	}

	expected := []int{-1, 2, -1}

	test(intervals, expected)
}

func main() {
	// 436. Find Right Interval
	test1()
	test2()
	test3()
}

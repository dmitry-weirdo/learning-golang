package main

import (
	"cmp"
	"fmt"
	"slices"
)

func maximumWeight(intervals [][]int) []int {
	n := len(intervals)

	// add indexes from original intervals array
	intervalsWithIndexes := createIntMatrix(n, 4)

	for i, v := range intervals {
		intervalsWithIndexes[i][0] = v[0] // left
		intervalsWithIndexes[i][1] = v[1] // right
		intervalsWithIndexes[i][2] = v[2] // score / weight
		intervalsWithIndexes[i][3] = i    // interval index in the original array
	}

	// sort by interval right end
	slices.SortFunc(intervalsWithIndexes, func(a, b []int) int {
		if a[1] != b[1] { // sort by right
			return cmp.Compare(a[1], b[1])
		}

		// right is equal -> sort by left
		return cmp.Compare(a[0], b[0])
	})

	//fmt.Printf("Intervals with indexes sorted by right, left: \n%v\n", intervalsWithIndexes)

	// DP matrix
	// i - using intervals from 0 to i
	// j - how many intervals are used (from 0 to 4)
	dp := createIntMatrix(n+1, 5) // todo: why N+1, not N?

	// for every dp[i][j], save the SORTED original interval indexes used
	indices := make([][][]int, n+1) // todo: why N+1, not N?
	for i := range n + 1 {
		indices[i] = make([][]int, 5)

		for j := range 5 { // fill with empty index lists // todo: do we need?
			indices[i][j] = make([]int, 0)
		}
	}

	for i, v := range intervalsWithIndexes {
		left := v[0]
		//right := v[1]
		weight := v[2]
		index := v[3]

		//fmt.Println()
		//fmt.Printf("Interval[%v]: left: %v, right: %v, weight: %v, original index: %v \n", i, left, right, weight, index)

		// Binary search gets the count of intervals whose intervals[k].right < intervals[i].left
		// search the previous interval[k] with intervals[k].right < intervals[i].left
		// !!! We're searching just to the left of [i], since the intervals are sorted
		// If there is no previous interval, we're using [i] itself
		k := searchRightmostLessThanTarget(intervalsWithIndexes, i, left)
		k++ // todo: why we're using the count of intervals, and NOT the k-th interval??? So in case of -1 it will be 0

		//fmt.Printf("Rightmost interval[%v] with right < interval[%v].left = %v: [%v; %v] \n", k, i, left, intervalsWithIndexes[k][0], intervalsWithIndexes[k][1])

		for j := 1; j < 5; j++ { // all possible intervals used up to index [i]
			s1 := dp[i][j]            // do not select the current interval (will be filled in from the previous step)
			s2 := dp[k][j-1] + weight // select [j - 1] intervals from the last interval that is before interval[i], add weight of interval[i]

			if s1 > s2 { // Adding an interval[i] from interval[k] provides NO improvement from the current dp[i][j]
				// propagate to [i + 1], so it will be s1 for the next [i]
				dp[i+1][j] = dp[i][j]
				indices[i+1][j] = copyArray(indices[i][j])
				continue
			}

			// handle the indexes
			newIndices := copyArray(indices[k][j-1]) // copy the indexes used for intervals[k]
			newIndices = append(newIndices, index)   // add the index of intervals[i]
			slices.Sort(newIndices)

			// compareLists compares lexicographically if the scores are equal
			if (s1 == s2) && compareListsLexicographically(indices[i][j], newIndices) < 0 {
				// if old result for [i][j] was the same result, and it was lexicographically smaller -> use the old result
				newIndices = copyArray(indices[i][j])
			}

			dp[i+1][j] = s2 // interval [i+1][j] is not selected?
			indices[i+1][j] = newIndices
		}
	}

	// collect the result from "after checking all N intervals", use 4 intervals
	return indices[n][4] // !!! notably it's [n], NOT [n - 1]
}

func searchRightmostLessThanTarget(arr [][]int, right int, target int) int {
	condition := func(x []int) bool {
		return x[1] >= target // leftmost with (right >= target)
	}

	index := binarySearchGeneric(
		arr,
		0,
		right, // insert position can be after the end of the array -> we're using the interval[i]
		condition,
	)

	// there is no previous element -> no result
	if index <= 0 {
		return -1
	}

	// go one element left -> this will be the last element < target
	index--

	if arr[index][1] >= target {
		return -1
	}

	return index
}

func binarySearchGeneric(
	arr [][]int, // todo: we can also generalize the type in the array
	left int, // usually it starts with 0, if we search in the complete array
	right int, // set len(arr) - 1 if you want to be within array. Set len(arr) if index after the array can be returned.
	condition func([]int) bool, // we will find the leftmost index satisfying this condition within [left; right] range
) int {
	// todo: this method can return an incorrect value for the empty array

	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		if condition(arr[mid]) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func compareListsLexicographically(a, b []int) int { // returns result as comparator
	minLength := min(len(a), len(b))

	for i := range minLength {
		if a[i] != b[i] { // different elements -> result is defined by comparing these elements
			return cmp.Compare(a[i], b[i])
		}
	}

	// if there are no differences in the elements, return the shorter array
	return cmp.Compare(len(a), len(b))
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func test(m [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Intervals: %v \n", m)

	result := maximumWeight(m)

	fmt.Printf("Indices of up to 4 intervals with the maximum total score: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

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
		[][]int{
			{1, 3, 2},
			{4, 5, 2},
			{1, 5, 5},
			{6, 9, 3},
			{6, 7, 1},
			{8, 9, 1},
		},
		[]int{2, 3},
	)
}

func test2() {
	test(
		[][]int{
			{5, 8, 1},
			{6, 7, 7},
			{4, 7, 3},
			{9, 10, 6},
			{7, 8, 2},
			{11, 14, 3},
			{3, 5, 5},
		},
		[]int{1, 3, 5, 6},
	)
}

func main() {
	// 3414. Maximum Score of Non-overlapping Intervals

	// This is a variation of "1235. Maximum Profit in Job Scheduling" with additional constraints:
	// - We can only select <= 4 intervals.
	// - Intervals with L[next] = R[prev] are considered intersecting.
	// - We need to return the leftmost indices if several interval combination have the same optimal score.
	test1()
	test2()
}

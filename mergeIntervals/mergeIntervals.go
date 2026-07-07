package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] { // if left is the same, sort by right
			return a[1] - b[1]
		}

		// sort by left
		return a[0] - b[0]
	})

	fmt.Printf("Sorted intervals: %v \n", intervals)

	// borders of the current merged interval
	result := make([][]int, 0)

	currLeft := intervals[0][0]
	currRight := intervals[0][1]

	for _, interval := range intervals {
		left := interval[0]
		right := interval[1]

		if left > currRight { // non-intersecting interval -> add existing merged interval to result
			fmt.Printf(
				"Interval [%v; %v] does NOT intersect with the current merged interval [%v; %v]. Adding the current merged interval [%v; %v] to the result. \n",
				left, right, currLeft, currRight, currLeft, currRight,
			)

			mergedInterval := []int{currLeft, currRight}
			result = append(result, mergedInterval)

			// set "current merged interval" to the first non-intersecting interval
			currLeft = left
			currRight = right

			continue
		}

		// new interval is intersecting -> merge with the current interval
		// basically, we just need to extend the right, since currLeft was already taken from the left-most or the merged interval

		// since the new interval can start later but end earlier, we need to select the max of the current right and new interval right
		// basically, we only need to extend if right > currRight

		//newRight := max(right, currRight)
		if right > currRight {
			fmt.Printf(
				"Interval [%v; %v] intersects with the current merged interval [%v; %v] and extends the right. Expanding the current merged interval from [%v; %v] to [%v; %v]. \n",
				left, right, currLeft, currRight, currLeft, currRight, currLeft, right,
			)

			currRight = right
		} else {
			fmt.Printf(
				"Interval [%v; %v] intersects with the current merged interval [%v; %v] but does not extend the right. Nothing to expand. \n",
				left, right, currLeft, currRight,
			)
		}
	}

	// add the last interval to the result
	fmt.Printf("Adding the last merged interval [%v; %v] to the result. \n", currLeft, currRight)

	mergedInterval := []int{currLeft, currRight}
	result = append(result, mergedInterval)

	return result
}

func test(intervals [][]int) {
	fmt.Println()
	fmt.Println("======================")

	fmt.Printf("Intervals: %v \n", intervals)

	mergedIntervals := merge(intervals)

	fmt.Printf("Merged intervals: %v \n", mergedIntervals)
}

func test1() {
	intervals := [][]int{
		{4, 7},
		{1, 4},
	}

	test(intervals)
}

func test2() {
	intervals := [][]int{
		{4, 7},
		{1, 4},
		{1, 3},
	}

	test(intervals)
}

func test3() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	test(intervals)
}

func test4() {
	intervals := [][]int{
		{1, 4},
		{2, 3},
	}

	test(intervals)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}

package main

import (
	"demo/matrixCommon"
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	// min value in matrix
	n := len(matrix)

	// left and right are values
	left := matrix[0][0]
	right := matrix[n-1][n-1]

	var mid int
	var count int

	for left < right {
		// mid is the binary search border

		// floor ((left + right) / 2)
		// so for negative values, -7 / 2 = -4, and NOT -3 as it would be for just the integer division
		// we can also use (left + right) >> 1
		//mid = int(math.Floor((float64(left) + float64(right)) / 2))
		mid = (left + right) >> 1

		fmt.Printf("left: %v, right: %v, mid %v \n", left, right, mid)

		// calculate the "from bottom-left" ladder matrix where the elements are <= mid
		count = countValuesLessThanOrEqual(matrix, n, mid)

		// we're searching for the leftmost value where count >= k,
		// i.e. the first element who has (k values) that are <= val (mid)
		if count >= k {
			// (there are enough elements that are <= mid) -> try to decrease mid to find element < mid that also satisfies this condition
			// for mid limit, we have more or same elements than k -> we go to the left of the binary search
			right = mid

			fmt.Printf("(Count = %v) >= (k = %v). Updated right to %v. Left:Right: [%v, %v] \n", count, k, right, left, right)
		} else {
			// (not enough elements are <= mid) -> result > mid
			// for mid limit, we have fewer elements than k -> we go to the right of the binary search
			left = mid + 1

			fmt.Printf("(Count = %v) < (k = %v). Updated left to %v. Left:Right: [%v, %v] \n", count, k, left, left, right)
		}
	}

	return left
}

func countValuesLessThanOrEqual(m [][]int, n int, mid int) int {
	// we start from bottom left
	i := n - 1
	j := 0

	count := 0

	// for values > mid, go up
	fmt.Printf("Searching for value %v \n", mid)
	fmt.Printf("i, j: [%v, %v] \n", i, j)

	for j < n {
		for (m[i][j] > mid) && (i > 0) {
			// in the current column j, go up until we find a row where mid can be
			// or until we reach the 0-th row
			// I.e. we're skipping the values from the bottom of the current column that exceed the target value
			i--

			fmt.Printf("i, j: [%v, %v] \n", i, j)
		}

		// after going up, we're sure that the lower rows can be excluded,
		// since it's guaranteed that every column is sorted

		// all values above in the column are added
		if m[i][j] <= mid { // do not add values of 0th row if they're above the value
			count += i + 1 // Index i is 0-based, we're adding values including the i-th row in the current column
		}

		// go one right
		j++
		fmt.Printf("i, j: [%v, %v] \n", i, j)
	}

	fmt.Printf("Returning count = %v \n", count)
	fmt.Println()
	return count
}

func test(mat [][]int, k int, expected int) {
	// todo: rewrite as common test with checking the result

	fmt.Println()
	fmt.Printf("======================== \n")
	matrixCommon.PrintIntMatrix(mat)
	kthSmallest := kthSmallest(mat, k)
	fmt.Printf("%v-th smallest element: %v \n", k, kthSmallest)

	if kthSmallest != expected {
		panic(fmt.Sprintf("Expected: %v, actual: %v", expected, kthSmallest))
	}
}

func test1() {
	// 1 5 9
	// 10 11 13
	// 12 14 15
	row0 := []int{1, 5, 9}
	row1 := []int{10, 11, 13}
	row2 := []int{12, 14, 15}

	mat := [][]int{
		row0,
		row1,
		row2,
	}

	k := 4
	expected := 10

	test(mat, k, expected)
}

func test2() { // with repeating values
	// 1 2
	// 1 3
	row0 := []int{1, 2}
	row1 := []int{1, 3}

	mat := [][]int{
		row0,
		row1,
	}

	k := 3
	expected := 2

	test(mat, k, expected)
}

func test3() {
	// 1 5 9
	// 10 11 13
	// 12 13 15
	// k = 8

	row0 := []int{1, 5, 9}
	row1 := []int{10, 11, 13}
	row2 := []int{12, 13, 15}

	mat := [][]int{
		row0,
		row1,
		row2,
	}

	k := 8
	expected := 13

	test(mat, k, expected)
}

func test4() {
	// -5
	// k = 1

	row0 := []int{-5}

	mat := [][]int{
		row0,
	}

	k := 1
	expected := -5

	test(mat, k, expected)
}

func test5() { // with repeating values
	// 1 2
	// 1 3
	// k = 4
	row0 := []int{1, 2}
	row1 := []int{1, 3}

	mat := [][]int{
		row0,
		row1,
	}

	k := 4
	expected := 3

	test(mat, k, expected)
}

func test6() { // with repeating values
	// -5 -4
	// -5 -4
	// k = 2
	row0 := []int{-5, -4}
	row1 := []int{-5, -4}

	mat := [][]int{
		row0,
		row1,
	}

	k := 2
	expected := -5 // -5, -5, -4, -4

	test(mat, k, expected)
}

func main() {
	// 378. Kth Smallest Element in a Sorted Matrix
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()

	fmt.Printf("-7 / 2: %v \n", -7/2)   // -3
	fmt.Printf("-7 >> 1: %v \n", -7>>1) // -4
}

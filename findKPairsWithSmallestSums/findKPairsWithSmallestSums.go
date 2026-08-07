package main

import (
	"fmt"
	"math"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	i := 0
	j := 0

	result := make([][]int, 0)
	result = append(result, []int{nums1[0], nums2[0]})

	for len(result) < k {
		nextLeft := math.MinInt32
		if i < len(nums1)-1 {
			nextLeft = nums1[i+1] + nums2[j]
		}

		nextRight := math.MinInt32
		if j < len(nums2)-1 {
			nextRight = nums1[i] + nums2[j+1]
		}

		if nextLeft < nextRight {
			result = append(result, []int{nums1[i+1], nums2[j]})
			i++
		} else {
			result = append(result, []int{nums1[i], nums2[j+1]})
			j++
		}
	}

	return result
}

func test(a1, a2 []int, k int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array 1: %v \n", a1)
	fmt.Printf("Array 2: %v \n", a1)
	fmt.Printf("K (count of smallest pairs): %v \n", k) // will always be 0

	result := kSmallestPairs(a1, a2, k)

	fmt.Printf("%v smallest pairs: \n%v \n", k, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := range expectedResult {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] {
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}

func test1() {
	a1 := []int{1, 7, 11}
	a2 := []int{2, 4, 6}
	k := 3

	expected := [][]int{
		{1, 2},
		{1, 4},
		{1, 6},
	}

	test(a1, a2, k, expected)
}

func test2() {
	a1 := []int{1, 1, 2}
	a2 := []int{1, 2, 3}
	k := 2

	expected := [][]int{
		{1, 1},
		{1, 1},
	}

	test(a1, a2, k, expected)
}

func test3() {
	a1 := []int{1, 2, 4, 5, 6}
	a2 := []int{3, 5, 7, 9}
	k := 3

	expected := [][]int{
		{1, 3},
		{2, 3},
		{1, 5},
	}

	test(a1, a2, k, expected)
}

func main() {
	// 373. Find K Pairs with Smallest Sums
	test1()
	test2()
	test3()
}

package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArraysNaive(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	totalLength := l1 + l2

	arr := make([]int, l1+l2)

	copy(arr, nums1)      // copy from 0-th element
	copy(arr[l1:], nums2) // you can copy from given element of dest!

	fmt.Printf("a1: %v \n", nums1)
	fmt.Printf("a2: %v \n", nums2)

	fmt.Printf("Merged array: %v \n", arr)
	sort.Ints(arr)

	fmt.Printf("Sorted merge array: %v \n", arr)

	if totalLength%2 == 0 { // even count of elements
		index1 := (totalLength / 2) - 1
		index2 := (totalLength / 2)

		sum := arr[index1] + arr[index2]

		return float64(sum) / float64(2)
	} else { // odd count of elements -> return middle elements
		value := arr[totalLength/2]
		return float64(value)
	}
}

func test(a1 []int, a2 []int) {
	fmt.Println()
	fmt.Println("===================")

	fmt.Printf("array 1: %v \n", a1)
	fmt.Printf("array 2: %v \n", a2)

	medianNaive := findMedianSortedArraysNaive(a1, a2)

	fmt.Printf("Naive median: %v \n", medianNaive)
}

func test1() {
	a1 := []int{1, 3}
	a2 := []int{2}

	test(a1, a2)
}

func test2() {
	a1 := []int{1, 3}
	a2 := []int{2, 4}

	test(a1, a2)
}

func main() {
	test1()
	test2()
}

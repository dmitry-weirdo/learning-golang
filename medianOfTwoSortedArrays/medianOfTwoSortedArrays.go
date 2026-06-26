package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	fmt.Printf("nums1: %v has length %v \n", nums1, m)
	fmt.Printf("nums2: %v has length %v \n", nums2, n)

	mergedLength := m + n
	fmt.Printf("Merged array will have length %v \n", mergedLength)

	if mergedLength%2 == 1 {
		// if mergedLength is odd, we need to find the middle element [mergedLength / 2]
		// len = 5 -> we need to find the merged[2], i.e. the 3rd biggest

		fmt.Printf("Merged array length %v is odd. We need to find element [%v] in the merged array. \n", mergedLength, mergedLength/2)
	} else {
		// if mergedLength is even, we need to find the average between elements [mergedLength / 2 - 1] and [mergedLength / 2]
		// len = 6 -> we need to find (merged[2] + merged[3]) / 2, i.e. between 3rd and 4th biggest

		fmt.Printf("Merged array length %v is even. We need to find average between elements [%v] and [%v] in the merged array. \n", mergedLength, mergedLength/2-1, mergedLength/2)
	}

	// Heuristic: we will find the average between
	// (m+n+1)/2 th biggest element and
	// (m+n+2)/2 th biggest element

	// Example even: len = 6 -> between 7 / 2 and 8 / 2 -> between 3-th and 4-th biggest elements -> expected median behaviour
	// Example odd: len = 5 -> between 6 / 2 and 7 / 2 -> between 3-th and 3-th biggest elements -> we'll return the same middle element
	k1 := (mergedLength + 1) / 2
	k2 := (mergedLength + 2) / 2

	fmt.Printf("Merged array length = %v. We will find an average between [%v]-th and [%v]-th biggest elements in the merged array. \n", mergedLength, k1, k2)

	k1Element := findKthBiggestElement(nums1, nums2, 0, 0, k1)
	k2Element := findKthBiggestElement(nums1, nums2, 0, 0, k2)

	// For odd total length: both positions point to the same middle element
	// For even total length: positions point to the two middle elements

	return float64(k1Element+k2Element) / 2.0
}

// i - starting index in array a
// j - starting index in array b
// k - we search for the k-th smallest element value. It is NOT 0-based
func findKthBiggestElement(a []int, b []int, i int, j int, k int) int {

	// end case -> both arrays not empty
	if k == 1 {
		return min(a[i], b[j])
	}

	return 0
}

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

	//medianNaive := findMedianSortedArraysNaive(a1, a2)
	median := findMedianSortedArrays(a1, a2)

	//fmt.Printf("Naive median: %v \n", medianNaive)
	fmt.Printf("O( log(m + n) ) median: %v \n", median)
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

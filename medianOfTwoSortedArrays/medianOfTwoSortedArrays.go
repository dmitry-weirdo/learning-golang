package main

import (
	"fmt"
	"math"
	"sort"
)

const INFINITY = math.MaxInt

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	fmt.Printf("nums1: %v has length %v \n", nums1, m)
	fmt.Printf("nums2: %v has length %v \n", nums2, n)

	mergedLength := m + n
	fmt.Printf("Merged array will have length %v \n", mergedLength)

	oddLength := mergedLength%2 == 1
	if oddLength {
		// if mergedLength is odd, we need to find the middle element [mergedLength / 2]
		// len = 5 -> we need to find the merged[2], i.e. the 3rd biggest

		fmt.Printf("Merged array length %v is odd. We need to find element [%v] in the merged array. \n", mergedLength, mergedLength/2)
	} else {
		// if mergedLength is even, we need to find the average between elements [mergedLength / 2 - 1] and [mergedLength / 2]
		// len = 6 -> we need to find (merged[2] + merged[3]) / 2, i.e. between 3rd and 4th biggest

		fmt.Printf("Merged array length %v is even. We need to find average between elements [%v] and [%v] in the merged array. \n", mergedLength, mergedLength/2-1, mergedLength/2)
	}

	// Without heuristic -> separate and obvious handling of odd and even cases
	if oddLength {
		// len = 5 -> we need the 3-rd biggest element
		k := (mergedLength + 1) / 2

		fmt.Printf("Merged array length = %v is odd. We will find %v-th biggest element in the merged array. \n", mergedLength, k)

		kElement := findKthBiggestElement(nums1, nums2, 0, 0, k)

		fmt.Printf("Returning %v-th biggest element = %v. \n", k, kElement)

		return float64(kElement)
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

	fmt.Printf("================================================\n")
	fmt.Printf("================================================\n")

	k2Element := findKthBiggestElement(nums1, nums2, 0, 0, k2)

	// For odd total length: both positions point to the same middle element
	// For even total length: positions point to the two middle elements

	fmt.Printf("Returning average between %v-th biggest element = %v and %v-th biggest element = %v. \n", k1, k1Element, k2, k2Element)
	return float64(k1Element+k2Element) / 2.0
}

// i - starting index in array a
// j - starting index in array b
// k - we search for the k-th smallest element value. Value K is NOT 0-based.
func findKthBiggestElement(a []int, b []int, i int, j int, k int) int {
	// we're using the arrays starting with a[i] and b[j]

	fmt.Println()
	fmt.Printf("============================\n")
	fmt.Printf("a start index: [%v], b start index [%v], k = %d \n", i, j, k)

	totalElementsAlreadyRemoved := i + j
	fmt.Printf("We already removed (%v elements from A) and (%v elements from B). Total of %v elements removed. \n", i, j, totalElementsAlreadyRemoved)

	// end case -> first array is empty -> just 1 array b left -> return k-th element starting with b[j]
	if i >= len(a) {
		bIndexForKthElement := j + k - 1 // array index starts with 0
		result := b[bIndexForKthElement]
		fmt.Printf("Array a exhausted. Returning %v-th biggest element starting with index %v. Returning b[%v] = %v. \n", k, i, bIndexForKthElement, result)

		return result
	}

	// end case -> second array is empty -> just 1 array a left -> return k-th element starting with a[i]
	if j >= len(b) {
		aIndexForKthElement := i + k - 1 // array index starts with 0
		result := a[aIndexForKthElement]
		fmt.Printf("Array b exhausted. Returning %v-th biggest element starting with index %v. Returning b[%v] = %v. \n", k, j, aIndexForKthElement, result)

		return result
	}

	// end case -> k = 1 and both arrays not empty -> return min of just 1 element from 2 arrays
	if k == 1 {
		result := min(a[i], b[j])
		fmt.Printf("k = 1 reached. Returning minimum of single elements a[%v] = %v and b[%v] = %v. Returning %v. \n", i, a[i], j, b[j], result)

		return result
	}

	elementsToRemove := k / 2
	elementsForNextIteration := k - elementsToRemove
	fmt.Printf("Current k: %v. Elements to remove: %v. Elements for the next run: %v \n", k, elementsToRemove, elementsForNextIteration)

	aMidIndex := i + elementsToRemove - 1
	bMidIndex := j + elementsToRemove - 1

	fmt.Printf("[i + k / 2] index in array a: %v \n", aMidIndex)
	fmt.Printf("[j + k / 2] index in array b: %v \n", bMidIndex)

	var removeFromA bool

	var aValue int
	if aMidIndex < len(a) {
		aValue = a[aMidIndex]
	} else {
		removeFromA = false
		aValue = INFINITY

		fmt.Printf("Index [%v] gets over array A of length %v. We will skip values from array B. \n", aMidIndex, len(a))
	}

	var bValue int
	if bMidIndex < len(b) {
		bValue = b[bMidIndex]
	} else {
		removeFromA = true
		bValue = INFINITY

		fmt.Printf("Index [%v] gets over array B of length %v. We will skip values from array A. \n", bMidIndex, len(b))
	}

	// todo: exhausting both arrays should NOT happen

	removeFromA = aValue < bValue

	var aNewIndex int
	var bNewIndex int

	if removeFromA {
		aNewIndex = i + elementsToRemove
		bNewIndex = j

		fmt.Printf("a[%v] = %v < b[%v] = %v \n", aMidIndex, aValue, bMidIndex, bValue)

		// aNewIndex can jump over len(a) - 1, therefore not logging a[aNewIndex]
		fmt.Printf("Jumping %v values ahead in array A. From a[%v] = %v to a[%v]. \n", elementsToRemove, i, a[i], aNewIndex)
		//fmt.Printf("Jumping %v values ahead in array A. From a[%v] = %v to a[%v] = %v. \n", elementsToRemove, i, a[i], aNewIndex, a[aNewIndex])
	} else {
		aNewIndex = i
		bNewIndex = j + elementsToRemove

		fmt.Printf("a[%v] = %v > b[%v] = %v \n", aMidIndex, aValue, bMidIndex, bValue)

		// bNewIndex can jump over len(b) - 1, therefore not logging b[bNewIndex]
		fmt.Printf("Jumping %v values ahead in array B. From b[%v] = %v to b[%v]. \n", elementsToRemove, j, b[j], bNewIndex)
		//fmt.Printf("Jumping %v values ahead in array B. From b[%v] = %v to b[%v] = %v. \n", elementsToRemove, j, b[j], bNewIndex, b[bNewIndex])
	}

	return findKthBiggestElement(a, b, aNewIndex, bNewIndex, elementsForNextIteration)
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

	medianNaive := findMedianSortedArraysNaive(a1, a2)
	median := findMedianSortedArrays(a1, a2)

	fmt.Printf("Naive median: %v \n", medianNaive)
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

func test3() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}

	test(a1, a2)
}

func test4() {
	a1 := []int{3}
	a2 := []int{1, 2, 4}

	test(a1, a2)
}

func main() {
	test1()
	//test2()
	//test3()
	//test4()
}

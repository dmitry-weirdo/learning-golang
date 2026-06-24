package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

// left and right inclusive
func quickSort(arr []int, leftIndex int, rightIndex int) {
	//fmt.Println()
	//fmt.Printf("quickSort the range [%v; %v] \n", leftIndex, rightIndex)

	if (rightIndex - leftIndex) <= 0 {
		//fmt.Printf(
		//	"There are just %v elements in range [%v; %v]. Nothing do sort. \n",
		//	rightIndex-leftIndex+1,
		//	leftIndex,
		//	rightIndex,
		//)

		return
	}

	pivotIndex := getRandomInt(leftIndex, rightIndex)
	//pivotIndex := 3 // for debugging
	//fmt.Printf("Selected pivot index: %v \n", pivotIndex)

	// swap pivot with the last value in the current range
	swap(arr, pivotIndex, rightIndex)
	//pivotValue := arr[rightIndex]

	//fmt.Printf("Selected pivot (swapped from [%v] to [%v]): %v \n", pivotIndex, rightIndex, pivotValue)
	//fmt.Printf("After swapping pivot:\n%v \n", arr)

	newPivotIndex := partition(arr, leftIndex, rightIndex)

	//fmt.Printf("[%v; %v]: executed partition around pivot %d \n", leftIndex, rightIndex, pivotValue)
	//printArray(arr, leftIndex, rightIndex, newPivotIndex)

	if (newPivotIndex > 0) && (leftIndex < newPivotIndex-1) {
		quickSort(arr, leftIndex, newPivotIndex-1)
	}

	if ((newPivotIndex + 1) < (len(arr) - 1)) && (newPivotIndex+1 < rightIndex) {
		quickSort(arr, newPivotIndex+1, rightIndex)
	}
}

func partition(arr []int, leftIndex int, rightIndex int) int { // returns index of the pivot
	left := leftIndex
	//right := rightIndex - 1 // pivot is at the right
	right := rightIndex // pivot is at the right, and yes, it must be included in the [left; right] scan

	pivotIndex := rightIndex
	pivot := arr[pivotIndex]

	//fmt.Printf("Doing partition on [%v; %v] around pivot %v \n", leftIndex, rightIndex, pivot)

	for left < right {
		// search >= pivot at left
		for (left < right) && (arr[left] <= pivot) {
			left++
		}

		// search <= pivot at right
		for (left < right) && (arr[right] >= pivot) {
			right--
		}

		// swap value for left with value for right
		swap(arr, left, right)
		//fmt.Printf("Swapped [%v] and [%v]:\n%v \n", left, right, arr)
	}

	// swap pivot to the left = right
	swap(arr, left, pivotIndex)

	return left
}

func getRandomInt(min int, max int) int {
	randomRange := max - min + 1 // from [0; n) - non-inclusive
	return min + rand.Intn(randomRange)
}

func swap(arr []int, i1 int, i2 int) {
	temp := arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = temp
}

func printArray(arr []int, left int, right int, pivotIndex int) {
	for i, v := range arr {
		if i == left {
			fmt.Printf("[")
		}

		if i == pivotIndex {
			fmt.Printf("|")
		}

		fmt.Printf("%v", v)

		if i == pivotIndex {
			fmt.Printf("|")
		}

		if i == right {
			fmt.Printf("]")
		}

		if i < (len(arr) - 1) { // no comma at the end
			fmt.Printf(",")
		}
	}

	fmt.Println()
}

func testRandomFunction() {
	for i := 1; i < 100; i++ {
		minValue := 3
		maxValue := 5
		randomInt := getRandomInt(minValue, maxValue)

		fmt.Printf("Random int in range [%v; %v]: %v \n", minValue, maxValue, randomInt)
	}
}

func generateRandomArray(size int, minValue int, maxValue int) []int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		randomInt := getRandomInt(minValue, maxValue)

		arr[i] = randomInt
	}

	return arr
}

func validateThatArrayIsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ { // i goes to pre-last element
		if arr[i] > arr[i+1] {
			fmt.Printf("Array is NOT sorted: a[%v] = %v > a[%v] = %v \n", i, arr[i], i+1, arr[i+1])

			return false
		}
	}

	return true
}

const (
	ONE_MILLION         = 1_000_000
	TEN_MILLION         = 10_000_000
	ONE_HUNDRED_MILLION = 100_000_000
	ONE_BILLION         = 1_000_000_000
)

func main() {
	/*	// partition test
		a := []int{0, 0, 1}
		partition(a, 0, 2)
		printArray(a, 0, 2, 2)

		return
	*/
	//testRandomFunction()

	size := ONE_BILLION
	minValue := -1000000
	maxValue := 1000000

	fmt.Printf("%v - Generating a random unsorted array of size: %d \n", time.Now(), size)
	arr := generateRandomArray(size, minValue, maxValue)

	// ONE_MILLION with values in range [-1000000; 1000000]
	// 2026-06-24 22:38:55.480415 +0200 CEST m=+0.000000001 - Generating a random unsorted array of size: 1000000
	// 2026-06-24 22:38:55.4974529 +0200 CEST m=+0.017037901 - Sorting a random unsorted array of size: 1000000
	// 2026-06-24 22:38:55.6328097 +0200 CEST m=+0.152394701 - Finished sorting an array of size: 1000000
	// Validating that the array is sorted...
	// 2026-06-24 22:38:55.6338105 +0200 CEST m=+0.153395501 - Finished validating an array of size: 1000000. Sorted: true

	// TEN_MILLION with values in range [-1000000; 1000000]
	//2026-06-24 22:40:18.4938417 +0200 CEST m=+0.000000001 - Generating a random unsorted array of size: 10000000
	//2026-06-24 22:40:18.6551964 +0200 CEST m=+0.161354701 - Sorting a random unsorted array of size: 10000000
	//2026-06-24 22:40:20.1222172 +0200 CEST m=+1.628375501 - Finished sorting an array of size: 10000000
	//Validating that the array is sorted...
	//2026-06-24 22:40:20.1357311 +0200 CEST m=+1.641889401 - Finished validating an array of size: 10000000. Sorted: true

	// ONE_HUNDRED_MILLION with values in range [-1000000; 1000000]
	//2026-06-24 22:41:02.5846464 +0200 CEST m=+0.000000001 - Generating a random unsorted array of size: 100000000
	//2026-06-24 22:41:04.1088092 +0200 CEST m=+1.524162801 - Sorting a random unsorted array of size: 100000000
	//2026-06-24 22:41:21.837668 +0200 CEST m=+19.253021601 - Finished sorting an array of size: 100000000
	//Validating that the array is sorted...
	//2026-06-24 22:41:21.9633242 +0200 CEST m=+19.378677801 - Finished validating an array of size: 100000000. Sorted: true

	// ONE_HUNDRED_MILLION with values in range [-1000000; 1000000] - much longer! - sorting was 6.5 minutes
	//2026-06-24 22:42:43.4374361 +0200 CEST m=+0.000000001 - Generating a random unsorted array of size: 1000000000
	//2026-06-24 22:43:01.7047376 +0200 CEST m=+18.267301501 - Sorting a random unsorted array of size: 1000000000
	//2026-06-24 22:50:27.2164805 +0200 CEST m=+463.779044401 - Finished sorting an array of size: 1000000000
	//Validating that the array is sorted...
	//2026-06-24 22:50:29.2163715 +0200 CEST m=+465.778935401 - Finished validating an array of size: 1000000000. Sorted: true

	//arr := []int{5, 1, 1, 2, 0, 0}

	fmt.Printf("%v - Sorting a random unsorted array of size: %d \n", time.Now(), len(arr))

	//fmt.Printf("Unsorted array: %v \n", arr)

	sortedArray := sortArray(arr)

	fmt.Printf("%v - Finished sorting an array of size: %d \n", time.Now(), len(sortedArray))

	fmt.Printf("Validating that the array is sorted... \n")
	sorted := validateThatArrayIsSorted(arr)
	fmt.Printf("%v - Finished validating an array of size: %d. Sorted: %v \n", time.Now(), len(sortedArray), sorted)

	//fmt.Printf("Sorted array: %v \n", sortedArray)
}

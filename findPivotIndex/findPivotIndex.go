package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)
	ps := make([]int, n+1)

	ps[0] = 0

	for i, v := range nums {
		ps[i+1] = ps[i] + v
	}

	fmt.Printf("Prefix sums: %v \n", ps)

	for i, _ := range nums {
		fmt.Printf("i: %v \n", i)

		// sum from a[0] to a[i - 1] -> ps[i] - ps[0]
		// left sum = ps[i] - ps[0]
		leftSum := ps[i] - ps[0]

		// sum from a[i + 1] to a[n - 1] -> ps[n] - ps[i + 1]
		// right sum = ps[n] - ps[i + 1]
		// in case of i = n - 1 (last element in the array), it will be ps[n] - ps[n] = 0
		rightSum := ps[n] - ps[i+1]

		//fmt.Printf("leftSum: %v rightSum: %v \n", leftSum, rightSum)

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func main() {
	// 724. Find Pivot Index
	arr := []int{1, 7, 3, 6, 5, 6}
	expected := 3

	fmt.Printf("Array: %v \n", arr)
	result := pivotIndex(arr)
	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expected)

}

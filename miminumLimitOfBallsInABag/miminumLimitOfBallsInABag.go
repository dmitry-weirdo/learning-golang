package main

import "fmt"

func minimumSize(nums []int, maxOperations int) int {
	left := 1                 // we're splitting into positive amounts, so minimum count of balls in a bag is 1
	right := maxInArray(nums) // max value in array

	fmt.Printf("Max value in array %v: %v \n", nums, right)

	// binary search for a result that is feasible with the given maxOperations (sum operations for every array value)
	// runs in log(right - left)
	result := right

	for left <= right { // yes, we have to check if left and right are the same
		// mid is the expected max(all bags) after maxOperations
		mid := left + ((right - left) / 2)

		fmt.Printf("\n=========================\n")
		fmt.Printf("left: %v, right: %v, mid: %v \n", left, right, mid)

		// we need to minimize the mid
		feasible := isFeasible(nums, maxOperations, mid)

		if feasible { // mid ok -> try to lower the mid (target)
			right = mid - 1
			result = mid

			fmt.Printf("Mid %v feasible -> moving right down: [%v, %v] \n", mid, left, right)
		} else { // mid failed -> try the upper the mid (targe)
			left = mid + 1

			fmt.Printf("Mid %v non-feasible -> moving left up: [%v, %v] \n", mid, left, right)
		}
	}

	return result
}

func maxInArray(nums []int) int {
	// we assume nums is not empty
	result := nums[0]

	for _, v := range nums {
		if v > result {
			result = v
		}
	}

	return result
}

func isFeasible(nums []int, maxOperations int, bagSizeTarget int) bool {
	totalOperations := 0

	for i := 0; i < len(nums); i++ {
		operationsForNum := calculate(nums[i], bagSizeTarget)
		fmt.Printf("%v / %v -> %v \n", nums[i], bagSizeTarget, operationsForNum)

		totalOperations += operationsForNum
		if totalOperations > maxOperations {
			fmt.Printf("Splitting an array into maxBagSize = %d with %d operations non-feasible. \n", bagSizeTarget, maxOperations)

			return false
		}
	}

	fmt.Printf("For array %v, maxBagSize = %v, totalOperations = %v \n", nums, bagSizeTarget, totalOperations)
	fmt.Printf("Splitting an array into maxBagSize = %d with %d operations is feasible! \n", bagSizeTarget, maxOperations)
	return true
}

func calculate(n int, max int) int { // how many operations are required to split N balls into max of int bags
	// todo: we can just return this (will be faster)
	return (n - 1) / max

	// to split n into bags of max, we need at least ceil(n / max) bags as a result, -1 bag that already exists
	// basically, we just subtract b (a / b) times, -1 on initial bag, and probable + 1 for the remainder part

	// ceil(a / b) = (a + b - 1 / b)
	ceil := (n + max - 1) / max

	if ceil <= 0 {
		return 0
	} else {
		return ceil - 1
	}
}

func test(nums []int, maxOperations int, expected int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("nums: %v \n", nums)
	fmt.Printf("maxOperations: %v \n", maxOperations)

	result := minimumSize(nums, maxOperations)

	fmt.Println()
	fmt.Printf("Expected result: %v \n", expected)
	fmt.Printf("Actual result: %v \n", result)
}

func test1() {
	nums := []int{9}
	maxOperations := 2
	expected := 3

	test(nums, maxOperations, expected)
}

func test2() {
	nums := []int{2, 4, 8, 2}
	maxOperations := 4
	expected := 2

	test(nums, maxOperations, expected)
}

func test3() {
	nums := []int{9}
	maxOperations := 200
	expected := 1

	test(nums, maxOperations, expected)
}

func main() {
	/*	max := 3
		for i := 0; i < 20; i++ {
			result := calculate(i, max)
			fmt.Printf("%v / %v -> %v \n", i, max, result)
		}
	*/

	//test1()
	test2()
	//test3()
}

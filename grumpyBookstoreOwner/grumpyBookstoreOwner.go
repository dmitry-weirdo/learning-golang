package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	// minutes = sliding window size
	n := len(customers)

	windowSize := minutes

	nonGrumpyClientsCount := 0

	// collect window from the start
	i := 0

	grumpyClientsInWindow := 0

	for i = 0; i < windowSize; i++ {
		if grumpy[i] == 0 { // non-grumpy will count anyway
			nonGrumpyClientsCount += customers[i]
		} else { // grumpy counts for the current window
			grumpyClientsInWindow += customers[i]
		}
	}

	maxGrumpyClients := grumpyClientsInWindow
	left := 0

	for ; i < n; i++ {
		// exclude left if grumpy
		if grumpy[left] == 1 {
			grumpyClientsInWindow -= customers[left]
		}

		left++

		// add new right if grumpy
		if grumpy[i] == 1 {
			grumpyClientsInWindow += customers[i]
		} else {
			nonGrumpyClientsCount += customers[i]
		}

		maxGrumpyClients = max(maxGrumpyClients, grumpyClientsInWindow)
	}

	return maxGrumpyClients + nonGrumpyClientsCount
}

func test(customers []int, grumpy []int, minutes int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Customers per minute: %v \n", customers)
	fmt.Printf("Grumpy per minute:    %v \n", grumpy)
	fmt.Printf("Minutes window where we can set non-grumpy: %v \n", minutes)

	result := maxSatisfied(customers, grumpy, minutes)

	fmt.Printf("Max satisfied customers: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	expected := 16

	test(customers, grumpy, minutes, expected)
}

func test2() {
	customers := []int{1}
	grumpy := []int{0}
	minutes := 1
	expected := 1

	test(customers, grumpy, minutes, expected)
}

func test3() {
	customers := []int{10, 1, 7}
	grumpy := []int{0, 0, 0}
	minutes := 2
	expected := 18

	test(customers, grumpy, minutes, expected)
}

func main() {
	// 1052. Grumpy Bookstore Owner
	test1()
	test2()
	test3()
}

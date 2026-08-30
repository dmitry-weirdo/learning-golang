package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	// we just need to sum every row and get the max
	maxWealth := 0 // all values are positive

	for _, bankSums := range accounts {
		sum := 0

		for _, v := range bankSums {
			sum += v
		}

		maxWealth = max(maxWealth, sum)
	}

	return maxWealth
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix of accounts: %v \n", m)

	result := maximumWealth(m)

	fmt.Printf("Max customer wealth: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{1, 2, 3}, // 6
		{3, 2, 1}, // 6
	}

	test(m, 6)
}

func test2() {
	m := [][]int{
		{1, 5}, // 6
		{7, 3}, // 10
		{3, 5}, // 8
	}

	test(m, 10)
}

func test3() {
	m := [][]int{
		{2, 8, 7}, // 17
		{7, 1, 3}, // 11
		{1, 9, 5}, // 15
	}

	test(m, 17)
}

func main() {
	// 1672. Richest Customer Wealth
	test1()
	test2()
	test3()
}

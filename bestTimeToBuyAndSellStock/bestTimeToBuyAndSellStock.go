package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := 100_000 // max price is 1_000

	maxWin := 0 // if there could be no profit, we should return 0

	for _, v := range prices {
		maxWin = max(maxWin, v-minPrice)

		minPrice = min(minPrice, v)
	}

	return maxWin
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Prices by day: %v \n", arr)

	result := maxProfit(arr)

	fmt.Printf("Max possible profit: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{7, 1, 5, 3, 6, 4},
		5,
	)
}

func test2() {
	test(
		[]int{7, 6, 4, 3, 1},
		0, // price only decreases -> no profit possible
	)
}

func test3() {
	test(
		[]int{1, 10},
		9,
	)
}

func main() {
	// 121. Best Time to Buy and Sell Stock
	test1()
	test2()
	test3()
}

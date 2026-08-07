package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	i := 0 // current position of who is buying tickets
	time := 0

	for tickets[k] > 0 {
		if tickets[i] != 0 { // if tickets exhausted -> skip
			tickets[i]--
			time++
		}

		i++
		i = i % len(tickets) // if we go after the last in the queue, start from the beginning
	}

	return time
}

func test(arr []int, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array of people requiring tickets: %v \n", arr)
	fmt.Printf("K - our position in the queue: %v \n", k)

	result := timeRequiredToBuy(arr, k)

	fmt.Printf("Required time for person %v to buy the tickets: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{2, 3, 2}
	k := 2
	expected := 6

	test(arr, k, expected)
}

func test2() {
	arr := []int{5, 1, 1, 1}
	k := 0
	expected := 8

	test(arr, k, expected)
}

func main() {
	// 2073. Time Needed to Buy Tickets
	test1()
	test2()
}

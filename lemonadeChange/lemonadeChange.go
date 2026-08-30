package main

import "fmt"

func lemonadeChange(bills []int) bool {
	// bills are only 5, 10, 20
	// we want to give a change in a greedy manner:
	// 5 bill - we just take it
	// 10 bill - we have to give back 5
	// 20 bill - try to give 10 and 5, else give 3 x 5

	count5 := 0
	count10 := 0
	count20 := 0

	for _, v := range bills {
		if v == 5 {
			count5++
		} else if v == 10 {
			if count5 <= 0 { // cannot give 5 change from 10
				return false
			}

			count5--  // give 5 back
			count10++ // increase count of 10 bills
		} else if v == 20 {
			if count10 <= 0 {
				if count5 < 3 { // cannot give 15 change from 20
					return false
				}

				count5 -= 3
				count20++
			} else { // there are 10 bills
				if count5 < 1 { // cannot give 15 change from 20
					return false
				}

				count10--
				count5--
				count20++
			}
		}
	}

	return true
}

func test(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Bills of the customers: %v \n", arr)

	result := lemonadeChange(arr)

	fmt.Printf("Possible to give change for every customer: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{5, 5, 5, 10, 20}, true)
}

func test2() {
	test([]int{5, 5, 10, 10, 20}, false)
}

func main() {
	// 860. Lemonade Change
	test1()
	test2()
}

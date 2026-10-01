package main

import "fmt"

func minimumCosts(regular []int, express []int, expressCost int) []int64 {
	// for every position[i], we can be there:
	// - in optimal state, currently on regular line
	// - in optimal state, currently on express line
	n := len(regular)

	result := make([]int64, n)

	i := 0

	optimalRegular := regular[i]
	optimalExpress := express[i] + expressCost
	result[i] = int64(min(optimalRegular, optimalExpress))

	regularToRegular := 0
	regularToExpress := 0
	expressToRegular := 0
	expressToExpress := 0

	for i < n-1 {
		i++

		// all 4 options from the [i-1]
		regularToRegular = optimalRegular + regular[i]
		regularToExpress = optimalRegular + express[i] + expressCost // regular -> express switch, add the expressCost

		expressToRegular = optimalExpress + regular[i] // express -> regular switch has no cost
		expressToExpress = optimalExpress + express[i]

		optimalRegular = min(regularToRegular, expressToRegular)
		optimalExpress = min(regularToExpress, expressToExpress)

		result[i] = int64(min(optimalRegular, optimalExpress))
	}

	return result
}

func test(regular, express []int, expressCost int, expectedResult []int64) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Regular prices: %v \n", regular)
	fmt.Printf("Express prices: %v \n", express)
	fmt.Printf("Switch to express cost: %v \n", expressCost)

	result := minimumCosts(regular, express, expressCost)

	fmt.Printf("Optimal prices: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		[]int{1, 6, 9, 5},
		[]int{5, 2, 3, 10},
		8,
		[]int64{1, 7, 14, 19},
	)
}

func test2() {
	test(
		[]int{11, 5, 13},
		[]int{7, 10, 6},
		3,
		[]int64{10, 15, 24},
	)
}

func main() {
	// 2361. Minimum Costs Using the Train Line
	test1()
	test2()
}

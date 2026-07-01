package main

import "fmt"

func numBusesToDestination(routes [][]int, source int, target int) int {
	return 0
}

func test(routes [][]int, source int, target int, expectedResult int) {
	fmt.Println()
	fmt.Println("=========================")

	fmt.Printf("routes: %v \n", routes)
	fmt.Printf("source = %v, target = %v \n", source, target)

	result := numBusesToDestination(routes, source, target)

	fmt.Printf("result (count of buses from %v to %v): %v \n", source, target, result)
	fmt.Printf("expected result:  %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	routes := [][]int{
		{1, 2, 7},
		{3, 6, 7},
	}

	source := 1
	target := 6
	expectedResult := 2

	test(routes, source, target, expectedResult)
}

func test2() {
	routes := [][]int{
		{1, 2, 3},
		{2, 4},
		{3, 5},
	}

	source := 1
	target := 5
	expectedResult := 2

	test(routes, source, target, expectedResult)
}

func test3() {
	routes := [][]int{
		{1, 2, 3},
		{3, 4, 5},
		{5, 6},
	}

	source := 1
	target := 6
	expectedResult := 3

	test(routes, source, target, expectedResult)
}

func main() {
	test1()
	test2()
	test3()
}

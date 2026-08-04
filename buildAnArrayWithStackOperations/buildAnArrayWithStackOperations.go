package main

import "fmt"

func buildArray(target []int, n int) []string {
	result := make([]string, 0)

	i := 1

	for _, v := range target {
		for i < v { // skip values missing in target
			result = append(result, "Push", "Pop")
			i++
		}

		// push the next value from target
		result = append(result, "Push")
		i++
	}

	return result
}

func test(arr []int, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	n := arr[len(arr)-1]

	fmt.Printf("Target array: %v \n", arr)
	fmt.Printf("N (max value from array): %v \n", n)

	result := buildArray(arr, n)

	fmt.Printf("Required stack operations: %v \n", result)
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
	arr := []int{1, 3}
	expectedResult := []string{"Push", "Push", "Pop", "Push"}

	test(arr, expectedResult)
}

func test2() {
	arr := []int{1, 2, 3}
	expectedResult := []string{"Push", "Push", "Push"}

	test(arr, expectedResult)
}

func test3() {
	arr := []int{1, 2}
	expectedResult := []string{"Push", "Push"}

	test(arr, expectedResult)
}

func main() {
	// 1441. Build an Array With Stack Operations
	test1()
	test2()
	test3()
}

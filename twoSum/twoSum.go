package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // value to index

	for i, v := range nums {
		requiredValue := target - v

		if q, ok := m[requiredValue]; ok {
			return []int{q, i}
		}

		m[v] = i
	}

	// this should never be reached, according to the constraints
	return []int{}
}

func test(nums []int, target int, expected []int) {
	fmt.Println()
	fmt.Println("===========================")
	fmt.Printf("nums: %v \n", nums)
	fmt.Printf("target: %v \n", target)

	result := twoSum(nums, target)
	fmt.Printf("expected result: %v \n", expected)
	fmt.Printf("actual result: %v \n", result)
}

func test1() {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectedResult := []int{0, 1}

	test(nums, target, expectedResult)
}

func test2() {
	nums := []int{3, 2, 4}
	target := 6
	expectedResult := []int{1, 2}

	test(nums, target, expectedResult)
}

func main() {
	test1()
	test2()
}

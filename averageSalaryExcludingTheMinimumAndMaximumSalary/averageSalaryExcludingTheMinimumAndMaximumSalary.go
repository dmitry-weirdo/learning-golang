package main

import "fmt"

func average(salary []int) float64 {
	// we subtract min and max values from the total sum and divide by n - 2
	n := len(salary)

	minValue := salary[0]
	maxValue := salary[0]
	sum := 0

	for _, v := range salary {
		minValue = min(minValue, v)
		maxValue = max(maxValue, v)
		sum += v
	}

	return float64(sum-minValue-maxValue) / float64(n-2)
}

func test(arr []int, expectedResult float64) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Salaries: %v \n", arr)

	result := average(arr)

	fmt.Printf("Average of salaries excluding min and max salaries: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{4000, 3000, 1000, 2000}, float64(2500))
}

func test2() {
	test([]int{1000, 2000, 3000}, float64(2000))
}

func main() {
	// 1491. Average Salary Excluding the Minimum and Maximum Salary
	test1()
	test2()
}

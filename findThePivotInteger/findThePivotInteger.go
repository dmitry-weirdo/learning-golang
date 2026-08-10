package main

import "fmt"

func pivotInteger(n int) int {
	return pivotInteger_runningSum(n) // don't store the prefixSums array
	//return pivotInteger_prefixSumArrays(n) // with calculating prefixSums array - uses O(n) memory
}

func pivotInteger_runningSum(n int) int {
	totalSum := getTotalSum(n)

	sum := 0

	for i := 1; i <= n; i++ {
		if (sum + i) == (totalSum - sum) { // ps[0] is actually 0
			return i
		}

		sum += i
	}

	return -1
}

func getTotalSum(n int) int {
	sum := 0

	for i := range n + 1 {
		sum += i
	}

	return sum
}

func pivotInteger_prefixSumArrays(n int) int {
	ps := getPrefixSums(n)

	for i := 1; i <= n; i++ {
		if (ps[i+1] - ps[0]) == (ps[n+1] - ps[i]) { // ps[0] is actually 0
			return i
		}
	}

	return -1
}

func getPrefixSums(n int) []int {
	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSums := make([]int, n+2)

	prefixSums[0] = 0

	for i := 1; i <= n; i++ {
		prefixSumIndex := i + 1
		prefixSums[prefixSumIndex] = prefixSums[i] + i
	}

	return prefixSums
}

func test(n int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", n)

	result := pivotInteger(n)

	fmt.Printf("Pivot integer for [1:%v]: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(8, 6)
}

func test2() {
	test(1, 1)
}

func test3() {
	test(4, -1)
}

func main() {
	// 2485. Find the Pivot Integer
	test1()
	test2()
	test3()
}

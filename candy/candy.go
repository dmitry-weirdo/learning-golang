package main

import "fmt"

func candy(ratings []int) int {
	// todo: can be solved with a slope method and constant space

	// O(n)
	// passes in 0-10 ms, already good
	return candy_oneArray(ratings)
}

func candy_oneArray(ratings []int) int {
	n := len(ratings)

	a := make([]int, n)

	// satisfy the left neighbor condition
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			a[i] = a[i-1] + 1
		}
	}

	// satisfy the right neighbor condition
	result := a[n-1] // add the last result

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			a[i] = max(a[i], a[i+1]+1) // condition might be already satisfied by going from left
		}

		result += a[i]
	}

	// we need to start with 1 candy for every position -> add n
	return result + n
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Children ratings: %v \n", arr)

	result := candy(arr)

	fmt.Printf("Minimum candies required: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 0, 2}, 5) // 2, 1, 2
}

func test2() {
	test([]int{1, 2, 2}, 4) // 1, 2, 1
}

func main() {
	// 135. Candy
	test1()
	test2()
}

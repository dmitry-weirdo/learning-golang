package main

import "fmt"

func missingMultiple(nums []int, k int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		if v%k == 0 {
			m[v] = true
		}
	}

	i := 1

	for {
		if !m[i*k] {
			return i * k
		}

		i++
	}

	panic("This must never happen.")
}

func test(arr []int, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K (divisor): %v \n", arr)

	result := missingMultiple(arr, k)

	fmt.Printf("Smallest positive multiple of K = %v missing from the array: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{8, 2, 3, 4, 6},
		2,
		10,
	)
}

func test2() {
	test(
		[]int{1, 4, 7, 10, 15},
		5,
		5,
	)
}

func main() {
	// 3718. Smallest Missing Multiple of K
	test1()
	test2()
}

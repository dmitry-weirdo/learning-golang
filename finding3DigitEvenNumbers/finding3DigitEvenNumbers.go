package main

import "fmt"

func findEvenNumbers(digits []int) []int {
	m := make([]int, 10) // count of digits

	for _, v := range digits {
		m[v]++
	}

	if (m[0] == 0) && (m[2] == 0) && (m[4] == 0) && (m[6] == 0) && (m[8] == 0) { // no even digits -> cannot generate even numbers
		return []int{}
	}

	result := make([]int, 0)

	d := make([]int, 10)

	for i := 100; i <= 998; i += 2 { // only iterate even numbers
		digit2 := i / 100
		digit1 := (i % 100) / 10
		digit0 := i % 10
		//fmt.Printf("i: %v, digits: %v %v %v \n", i, digit2, digit1, digit0)

		// count the digits of the current number
		d[digit0]++
		d[digit1]++
		d[digit2]++

		// check whether there are enough digits for the current number
		if d[digit0] <= m[digit0] &&
			d[digit1] <= m[digit1] &&
			d[digit2] <= m[digit2] {
			result = append(result, i)
		}

		// backtrack the counts to 0
		d[digit0]--
		d[digit1]--
		d[digit2]--
	}

	return result
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Digits: %v \n", arr)

	result := findEvenNumbers(arr)

	fmt.Printf("All even 3-digit numbers that could be generated from the digits : %v \n", result)
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
		[]int{2, 1, 3, 0},
		[]int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
	)
}

func test2() {
	test(
		[]int{2, 2, 8, 8, 2},
		[]int{222, 228, 282, 288, 822, 828, 882},
	)
}

func test3() {
	test(
		[]int{3, 7, 5},
		[]int{}, // only odd numbers -> not a single even number could be generated
	)
}

func main() {
	// 2094. Finding 3-Digit Even Numbers
	// It's the same as "3483. Unique 3-Digit Even Numbers"
	test1()
	test2()
	test3()
}

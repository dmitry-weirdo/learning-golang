package main

import "fmt"

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	i := 0 // pos in encoded1
	j := 0 // pos in encoded2

	result := make([][]int, 0)

	for i < len(encoded1) && j < len(encoded2) {
		e1 := encoded1[i]
		e2 := encoded2[j]

		minFrequency := min(e1[1], e2[1])

		product := e1[0] * e2[0]

		if len(result) > 0 && result[len(result)-1][0] == product {
			// if the product is the same -> merge with previous product (add frequency)
			result[len(result)-1][1] += minFrequency
		} else {
			// product is different -> append it to the result
			productFrequency := []int{product, minFrequency}

			result = append(result, productFrequency)
		}

		// todo: we're modifying the current arrays. If we want to keep the input data unchanged, we can use copy of it or revert e1[1] when we do i++
		// decrease frequency of encoded1[i]
		e1[1] -= minFrequency
		if e1[1] == 0 {
			i++
		}

		// decrease frequency of encoded2[j]
		e2[1] -= minFrequency
		if e2[1] == 0 {
			j++
		}
	}

	return result
}

func test(e1, e2 [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Encoded array 1: %v \n", e1)
	fmt.Printf("Encoded array 2: %v \n", e2)

	result := findRLEArray(e1, e2)

	fmt.Printf("Compressed product array: %v \n", result)
	fmt.Printf("Expected result:          %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := 0; i < len(expectedResult); i++ {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] {
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}

func test1() {
	e1 := [][]int{
		{1, 3},
		{2, 3},
	}

	e2 := [][]int{
		{6, 3},
		{3, 3},
	}

	expected := [][]int{
		{6, 6},
	}

	test(e1, e2, expected)
}

func test2() {
	e1 := [][]int{
		{1, 3},
		{2, 1},
		{3, 2},
	}

	e2 := [][]int{
		{2, 3},
		{3, 3},
	}

	expected := [][]int{
		{2, 3},
		{6, 1},
		{9, 2},
	}

	test(e1, e2, expected)
}

func test3() {
	e1 := [][]int{
		{2, 3},
		{3, 2},
	}

	e2 := [][]int{
		{5, 2},
		{1, 3},
	}

	expected := [][]int{
		{10, 2},
		{2, 1},
		{3, 2},
	}

	test(e1, e2, expected)
}

func main() {
	// 1868. Product of Two Run-Length Encoded Arrays
	test1()
	test2()
	test3()
}

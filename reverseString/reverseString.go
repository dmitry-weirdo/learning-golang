package main

import "fmt"

func reverseString(s []byte) {
	// slices.Reverse(s) // this is a built-in solution, let's do it manually

	n := len(s)

	for i := range n / 2 {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func test(arr []byte, expectedResult []byte) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array of characters: %s \n", arr)

	reverseString(arr)
	result := arr // replacement in-place

	fmt.Printf("Reversed string: %s \n", result)
	fmt.Printf("Expected result: %s \n", expectedResult)

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
		[]byte{'h', 'e', 'l', 'l', 'o'},
		[]byte{'o', 'l', 'l', 'e', 'h'},
	)
}

func test2() {
	test(
		[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
		[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
	)
}

func main() {
	// 344. Reverse String
	test1()
	test2()
}

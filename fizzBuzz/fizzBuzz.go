package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}

func test(n int, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Number N: %v \n", n)

	result := fizzBuzz(n)

	fmt.Printf("Fizz-buzz array for [1; %v]: %v \n", n, result)
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
	test(3, []string{"1", "2", "Fizz"})
}

func test2() {
	test(5, []string{"1", "2", "Fizz", "4", "Buzz"})
}

func test3() {
	test(15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"})
}

func main() {
	// 412. Fizz Buzz
	test1()
	test2()
	test3()
}

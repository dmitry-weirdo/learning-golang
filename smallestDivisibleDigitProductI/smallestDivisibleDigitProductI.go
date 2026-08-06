package main

import "fmt"

func smallestNumber(n int, t int) int {
	// we only have to check current and 9 next numbers, since 1 of them will have last digit = 0 that will lead to solution
	for i := n; i <= n+9; i++ {
		if getDigitsProduct(i)%t == 0 {
			return i
		}
	}

	panic("This must never happen.")
}

func getDigitsProduct(n int) int {
	product := 1

	for n > 0 {
		product *= n % 10
		n /= 10
	}

	return product
}

func test(n, t int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("N - min dividable number: %v \n", n)
	fmt.Printf("T - divisor: %v \n", t)

	result := smallestNumber(n, t)

	fmt.Printf("Min target value >= N = %v with digits product divisible by T = %v: %v \n", n, t, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	n := 10
	t := 2
	expectedResult := 10

	test(n, t, expectedResult)
}

func test2() {
	n := 15
	t := 3
	expectedResult := 16

	test(n, t, expectedResult)
}

func test3() {
	n := 1
	t := 10
	expectedResult := 10

	test(n, t, expectedResult)
}

func main() {
	// 3345. Smallest Divisible Digit Product I
	test1()
	test2()
	test3()
}

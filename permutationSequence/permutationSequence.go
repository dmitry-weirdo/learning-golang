package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	factorials := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880} // 0! is 1

	// todo: maybe better structure then the array to remove from central positions?
	a := make([]int, n)

	for i := range n {
		a[i] = i + 1 // numbers go from 1 to N
	}

	//fmt.Printf("N = %v permutation numbers: %v \n", n, a)
	//fmt.Printf("%v! = %v \n", n, factorials[n])
	//fmt.Printf("K = %v (K-th permutation to find) \n", k)

	// !!! note that 0! is 1, not 0
	// k = a[n-1] * (n - 1)! + a[n-2] * (n - 2)! + ... + a[1] * 1! + a[0] * 0!

	// todo: understand this magic. It moves K instead of [1 ; N!] into [0 ; N!- 1] range.
	//remaining := k
	remaining := k - 1

	currentDigit := n

	var sb strings.Builder

	for currentDigit > 0 {
		// how many combinations there are in the previous digits
		prevDigitFactorial := factorials[currentDigit-1]

		div := remaining / prevDigitFactorial
		mod := remaining % prevDigitFactorial

		//fmt.Println()
		//fmt.Printf("Power: %v. Remaining: %v. Factorial of prev power: %v. Div: %v, mod: %v. \n", currentDigit, remaining, prevDigitFactorial, div, mod)

		digitIndexToRemove := div

		/*		if mod != 0 { // we take exactly div-th digit from the remaining
					digitIndexToRemove = div
				} else { // we take (div - 1)-th digit from the remaining
					if div == 0 {
						digitIndexToRemove = 0
					} else {
						digitIndexToRemove = div - 1
					}
				}
		*/
		digit := a[digitIndexToRemove]
		a = append(a[:digitIndexToRemove], a[digitIndexToRemove+1:]...)

		// append to result
		sb.WriteString(strconv.Itoa(digit))

		//fmt.Printf("Digit: %v \n", digit)
		//fmt.Printf("Remaining digits: %v \n", a)

		//remaining -= div * prevDigitFactorial
		remaining = mod

		currentDigit--
	}

	return sb.String()
}

func test(n, k int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N (number of digits to permutate): %v \n", n)
	fmt.Printf("K (1-based what permutation is required): %v \n", k)

	result := getPermutation(n, k)

	fmt.Printf("%v-th permutation of numbers from 1 to %v: %v \n", k, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(4, 9, "2314")
}

func test2() {
	test(3, 3, "213")
}

func test3() {
	test(3, 1, "123")
}

func test4() {
	test(3, 2, "132")
}

func main() {
	// 60. Permutation Sequence
	test1()
	test2()
	test3()
	test4()
}

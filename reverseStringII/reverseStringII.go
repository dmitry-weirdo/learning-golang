package main

import "fmt"

func reverseStr(s string, k int) string {
	if k == 1 {
		return s
	}

	n := len(s)

	b := []byte(s) // byte array is mutable, string is immutable

	// let's do it in place
	i := 0 // block start
	j := 0 // swapping character indexes

	for i < n {
		if /*(i+2*k) <= n ||*/ i+k <= n { // normal case -> reverse the first K characters from current pos
			for j = range k / 2 {
				b[i+j], b[i+k-j-1] = b[i+k-j-1], b[i+j]
			}

			i += 2 * k
		} else { // less than K chars left -> reverse everything that is left
			charsLeft := n - i
			k = charsLeft

			for j = range k / 2 {
				b[i+j], b[i+k-j-1] = b[i+k-j-1], b[i+j]
			}

			i = n
		}
	}

	return string(b)
}

func test(s string, k int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("K: %v \n", k)

	result := reverseStr(s, k)

	fmt.Printf("String reverted by K = %v characters: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abcdefg", 2, "bacdfeg")
}

func test2() {
	test("abcd", 2, "bacd")
}

func main() {
	// 541. Reverse String II
	test1()
	test2()
}

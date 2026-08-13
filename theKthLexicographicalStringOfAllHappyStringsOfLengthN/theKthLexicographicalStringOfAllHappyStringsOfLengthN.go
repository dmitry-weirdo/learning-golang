package main

import "fmt"

func getHappyString(n int, k int) string {
	// n - string length
	// k - number of valid "happy" string to return - counting from 1

	foundStringNumber := 0
	kthString := "" // if there are less than K happy strings

	var dfs func(s string)

	dfs = func(s string) {
		if kthString != "" { // k-th string already found - no further iteration required
			return
		}

		if len(s) > n { // too many characters
			return
		}

		if len(s) == n {
			foundStringNumber++

			if foundStringNumber == k { // target result found
				kthString = s
			}
		}

		if canAdd(s, 'a') {
			dfs(s + "a")
		}

		if canAdd(s, 'b') {
			dfs(s + "b")
		}

		if canAdd(s, 'c') {
			dfs(s + "c")
		}
	}

	dfs("")

	return kthString
}

func canAdd(s string, char byte) bool {
	if s == "" { // to empty string, we can add ahy character
		return true
	}

	// check whether last char != char
	return s[len(s)-1] != char
}

func test(n, k int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Length of strings: %v \n", n)
	fmt.Printf("K (1-based) - : %v \n", k)

	result := getHappyString(n, k)

	fmt.Printf("%v-th string or length %v: %v \n", k, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, 4, "") // only 3 happy strings of length 1
}

func test2() {
	test(1, 3, "c") // "a", "b", "c" - counting from 1
}

func test3() {
	// "aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"
	test(3, 9, "cab")
}

func main() {
	// 1415. The k-th Lexicographical String of All Happy Strings of Length n
	test1()
	test2()
	test3()
}

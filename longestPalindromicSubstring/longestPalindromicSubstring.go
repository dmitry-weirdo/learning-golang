package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)

	// DP-memo matrix of whether string [i, j] is a palindrome
	// values of i > j (bottom-left) are not used.
	m := make([][]bool, n)

	// put true values to [i][i] (since 1 character is a palindrome), other values to false
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
		m[i][i] = true
	}

	fmt.Println("Initialized the matrix:")
	printMatrix(m)

	return ""
}

func printMatrix(mat [][]bool) {
	rows := len(mat)
	columns := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var v string

			if i > j { // we don't care about these values
				v = "."
			} else if mat[i][j] {
				v = "T"
			} else {
				v = "F"
			}

			fmt.Printf("%v ", v)
		}

		fmt.Println()
	}
}

func test(s string) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("String: %v \n", s)

	palindrome := longestPalindrome(s)

	fmt.Printf("Longest palindrome: %v \n", palindrome)
}

func test1() {
	s := "babab"

	test(s)
}

func main() {
	test1()
}

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

	// initialize the max palindrome as just the first character
	maxPalindromeStart := 0
	maxPalindromeLength := 1

	// we skip i = n - 1, it's just the last character that is already true
	for i := n - 2; i >= 0; i-- {
		// go from smaller intervals to bigger since we need the values of [i + 1; j - 1]
		for j := i + 1; j <= n-1; j++ {
			if s[i] != s[j] {
				// first and last characters of substring are non-equal -> not a palindrome
				m[i][j] = false // todo: this is not super-necessary since initialized to false anyway
				continue
			}

			isPalindrome := false

			if ((i + 1) >= (j - 1)) || // if i + 1 = j, it is a palindrome of 2 characters!
				m[i+1][j-1] { // substring within i and j is a palindrome
				isPalindrome = true
			}

			if !isPalindrome {
				m[i][j] = false // todo: this is not super-necessary since initialized to false anyway
				continue
			}

			// [i;j] is a palindrome -> check whether it's longer than the current max palindrome
			m[i][j] = true
			currentLength := j - i + 1

			if currentLength > maxPalindromeLength {
				maxPalindromeStart = i
				maxPalindromeLength = currentLength

				fmt.Printf("New max palindrome [%v, %v] \"%v\" of length %v found. \n", i, j, s[i:j+1], currentLength)
			}
		}
	}

	fmt.Println("Matrix after all checks:")
	printMatrix(m)

	return s[maxPalindromeStart : maxPalindromeStart+maxPalindromeLength]
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
	s := "babad"

	test(s)
}

func test2() {
	s := "cbbd"

	test(s)
}

func main() {
	// 5. Longest Palindromic Substring
	test1()
	test2()
}

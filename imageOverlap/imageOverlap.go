package main

import (
	"demo/matrixCommon"
	"fmt"
)

func largestOverlap(img1 [][]int, img2 [][]int) int {
	// stupid O(N^6) or like this
	// Passes in 165-175 ms

	// !!! It's not required to shift M2, we can just shift M1 in both directions.
	// Then it's twice faster -
	// Passes in 78-81 ms
	return largestOverlap_bruteForce(img1, img2)
}

func largestOverlap_bruteForce(m1 [][]int, m2 [][]int) int {
	n := len(m1)

	ones1 := countOneBits(m1)
	ones2 := countOneBits(m2)

	// We can only get the maximum of min ones from both tables.
	// If we reach this result with any transformation -> we can return this immediately.
	maxPossibleResult := min(ones1, ones2)
	//fmt.Printf("Max possible intersection result: %v \n", maxPossibleResult)

	maxResult := -1
	count := -1

	// use a single allocated memory to shift the matrix
	bufX := createIntMatrix(n, n)
	bufY := createIntMatrix(n, n)

	// we can shift from ( -(n - 1) ) to (n - 1) in both coordinates (x and y)
	for xShift := -(n - 1); xShift <= (n - 1); xShift++ {
		for yShift := -(n - 1); yShift <= (n - 1); yShift++ {
			m1Shifted := shift(m1, bufX, bufY, xShift, yShift)
			count = countIntersection(m1Shifted, m2)
			maxResult = max(maxResult, count)

			/*			fmt.Println()
						fmt.Printf("M1 shifted by [%v; %v]: \n", xShift, yShift)
						matrixCommon.PrintIntMatrix(m1Shifted)

						fmt.Printf("Intersection with M2: %v \n", count)
			*/
			if maxResult == maxPossibleResult {
				return maxPossibleResult
			}

			// actually, we don't need to shift M2, shifting M1 is enough.
			/*
				m2Shifted := shift(m2, bufX, bufY, xShift, yShift)
				count = countIntersection(m2Shifted, m1)
				maxResult = max(maxResult, count)

				//fmt.Println()
				//fmt.Printf("M2 shifted by [%v; %v]: \n", xShift, yShift)
				//matrixCommon.PrintIntMatrix(m2Shifted)
				//
				//fmt.Printf("Intersection with M1: %v \n", count)

				if maxResult == maxPossibleResult {
					return maxPossibleResult
				}
			*/
		}
	}

	return maxResult
}

func shift(m, t, u [][]int, x, y int) [][]int {
	n := len(m)

	// clear T memory with all 0 bits
	for i := range n {
		for j := range n {
			t[i][j] = 0
		}
	}

	// copy from M to T with X shift
	tj := 0 // write pos

	if x > 0 { // shift right
		// we skip X columns and they remain 0
		tj += x

		// all rows -> copy all other columns from the 0 to (n - 1 - x)
		for i := range n {
			for j := 0; j < (n - x); j++ {
				t[i][tj+j] = m[i][j]
			}
		}
	} else if x < 0 { // shift left
		// we skip |X| columns at the left of the original matrix
		tj = 0

		absX := abs(x)

		// all rows -> copy all columns from -x to (n -1)
		for i := range n {
			for j := 0; j < (n - absX); j++ {
				t[i][tj+j] = m[i][j+absX]
			}
		}
	} else { // no X shift -> copy as is
		for i := range n {
			for j := range n {
				t[i][j] = m[i][j]
			}
		}
	}

	// Y shifts are executed on T already

	// clear U memory with all 0 bits
	for i := range n {
		for j := range n {
			u[i][j] = 0
		}
	}

	ui := 0 // write pos

	if y > 0 { // shift down
		// we skip Y rows and they remain 0
		ui += y

		// all columns -> copy all other rows from 0 to (n -1 - y)
		for i := 0; i < (n - y); i++ {
			for j := range n {
				u[ui+i][j] = t[i][j]
			}
		}
	} else if y < 0 { // shift up
		ui = 0

		absY := abs(y)

		// all columns -> copy all row from -y to (n -1)
		for i := 0; i < (n - absY); i++ {
			for j := range n {
				u[ui+i][j] = t[i+absY][j]
			}
		}
	} else { // no Y shift -> return T as is
		return t
	}

	return u
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func countOneBits(m [][]int) int {
	n := len(m)

	count := 0

	for i := range n {
		for j := range n {
			if m[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func countIntersection(m1 [][]int, m2 [][]int) int {
	n := len(m1)

	count := 0

	for i := range n {
		for j := range n {
			if m1[i][j] == 1 && m2[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func copyIntMatrix(m [][]int) [][]int {
	r := make([][]int, len(m))

	for i, row := range m {
		// todo: will copy array work faster for every row?
		r[i] = make([]int, len(row))

		for j, v := range row {
			r[i][j] = v
		}
	}

	return r
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func test(m1, m2 [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Image matrix 1: \n")
	matrixCommon.PrintIntMatrix(m1)

	fmt.Println()
	fmt.Printf("Image matrix 2: \n")
	matrixCommon.PrintIntMatrix(m2)

	result := largestOverlap(m1, m2)

	fmt.Println()
	fmt.Printf("Largest overlap: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m1 := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	m2 := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}

	expected := 3

	test(m1, m2, expected)
}

func test2() {
	m1 := [][]int{
		{1},
	}

	m2 := [][]int{
		{1},
	}

	expected := 1

	test(m1, m2, expected)
}

func test3() {
	m1 := [][]int{
		{0},
	}

	m2 := [][]int{
		{0},
	}

	expected := 0

	test(m1, m2, expected)
}

func main() {
	// 835. Image Overlap
	test1()
	test2()
	test3()
}

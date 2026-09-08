package main

import (
	"demo/matrixCommon"
	"fmt"
	"slices"
)

func highFive(items [][]int) [][]int {
	// sort by id asc, score desc
	slices.SortFunc(items, func(a, b []int) int {
		if a[0] != b[0] { // id asc
			return a[0] - b[0]
		}

		// score desc
		return b[1] - a[1]
	})

	//fmt.Printf("Sorted by id asc, score desc: \n%v \n", items)

	// for every student, calculate top 5
	result := make([][]int, 0)

	id := -1
	score := 0
	count := 0

	for _, v := range items {
		if v[0] != id { // switch to new student
			if id > 0 { // add the current student to the result
				result = append(result, []int{id, score / count})
			}

			id = v[0]
			score = v[1]
			count = 1
			continue
		}

		if count < 5 { // for the current student, add count
			score += v[1]
			count++
		}
	}

	// add last result at the end of the array
	if id > 0 { // add the current student to the result
		result = append(result, []int{id, score / count})
	}

	return result
}

func test(m [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Student single scores: \n")
	matrixCommon.PrintIntMatrix(m)

	result := highFive(m)

	fmt.Printf("Average of top 5 scores of every student, ordered by student id: \n")
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Expected result: \n")
	matrixCommon.PrintIntMatrix(expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	//Input: items = [[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
	//Output: [[1,87],[2,88]]

	test(
		[][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}},
		[][]int{{1, 87}, {2, 88}},
	)
}

func test2() {
	//Input: items = [[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100]]
	//Output: [[1,100],[7,100]]

	test(
		[][]int{{1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100}},
		[][]int{{1, 100}, {7, 100}},
	)
}

func main() {
	// 1086. High Five
	test1()
	test2()
}

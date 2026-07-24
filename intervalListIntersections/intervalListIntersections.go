package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// if the lists are not guaranteed to be disjoint,
	// we first merge all the intervals in both lists and then proceed with the same algorithm
	i := 0
	j := 0

	result := make([][]int, 0)

	for i < len(firstList) && j < len(secondList) {
		s1 := firstList[i][0]
		e1 := firstList[i][1]

		s2 := secondList[j][0]
		e2 := secondList[j][1]

		maxStart := max(s1, s2)
		minEnd := min(e1, e2)

		// intersection is [maxStart; minEnd] if maxStart <= minEnd
		if maxStart <= minEnd {
			intersection := []int{maxStart, minEnd}

			result = append(result, intersection)
		}

		// advance the interval that ends earlier,
		// since another interval (ending later) can still intersect with the next intervals in the opposite list
		if e1 < e2 {
			i++
		} else {
			j++
		}
	}

	return result
}

func test(a, b [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("First list: %v \n", a)
	fmt.Printf("Second list: %v \n", b)

	result := intervalIntersection(a, b)

	fmt.Printf("Intersections:   %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
	}

	for i := 0; i < len(expectedResult); i++ {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] {
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}

func test1() {
	a := [][]int{
		{1, 3},
	}

	b := [][]int{
		{2, 4},
	}

	expected := [][]int{
		{2, 3},
	}

	test(a, b, expected)
}

func test2() {
	a := [][]int{
		{1, 3},
	}

	b := [][]int{
		{5, 9},
	}

	expected := [][]int{}

	test(a, b, expected)
}

func test3() {
	a := [][]int{
		{0, 2},
		{5, 10},
		{13, 23},
		{24, 25},
	}

	b := [][]int{
		{1, 5},
		{8, 12},
		{15, 24},
		{25, 26},
	}

	expected := [][]int{
		{1, 2},
		{5, 5},
		{8, 10},
		{15, 23},
		{24, 24},
		{25, 29},
	}

	test(a, b, expected)
}

func main() {
	// 986. Interval List Intersections
	test1()
	test2()
	test3()
}

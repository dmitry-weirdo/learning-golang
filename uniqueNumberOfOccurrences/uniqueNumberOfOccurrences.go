package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	// fill elements by count map
	m := make(map[int]int)

	// collect elements to count map
	for _, v := range arr {
		frequency, ok := m[v]

		if !ok { // add value the first time
			m[v] = 1
		} else {
			m[v] = frequency + 1
		}
	}

	fmt.Printf("map of counts: %v \n", m)

	// check whether counts are unique
	counts := make(map[int]int)

	for value, count := range m {
		existingValueWithSameCount, ok := counts[count]

		if ok {
			fmt.Printf("Value %v has the same count of occurrences = %v as the number %v. Returning false.\n", value, count, existingValueWithSameCount)

			return false
		} else {
			counts[count] = value
		}
	}

	return true
}

func test(a []int) {
	fmt.Println()
	fmt.Println("============================")
	fmt.Printf("Array: %v \n", a)
	result := uniqueOccurrences(a)

	fmt.Printf("Result: %v \n", result)
}

func test1() {
	a := []int{1, 2, 2, 1, 1, 3}
	test(a)
}

func test2() {
	a := []int{1, 2}
	test(a)
}

func test3() {
	a := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	test(a)
}

func main() {
	test1()
	test2()
	test3()
}

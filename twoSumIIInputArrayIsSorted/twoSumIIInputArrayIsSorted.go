package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target { // found the result!
			return []int{left + 1, right + 1}
		}

		if sum < target {
			// we need to increase -> increase the smaller
			left++
		} else {
			// we need to decrease -> decrease the bigger
			right--
		}
	}

	// this must never happen
	// todo: we don't handle if the solution does not exist
	return []int{}
}

func test(arr []int, target int, expectedResult []int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", target)

	result := twoSum(arr, target)

	// return indexes are 1-based!
	value1 := arr[result[0]-1]
	value2 := arr[result[1]-1]

	fmt.Printf("1-based indexes with targetSum = %v: %v \n", target, result)
	fmt.Printf("Values: %v + %v = %v. Target = %v. \n", value1, value2, value1+value2, target)
	fmt.Printf("Expected result:          %v \n", expectedResult)

	if value1+value2 != target {
		fmt.Printf("FAILURE: %v + %v = %v != target = %v \n", value1, value2, value1+value2, target)
	}

	if result[0] != expectedResult[0] {
		fmt.Printf("FAILURE: expected result[0] = %v, actual result[0] = %v \n", expectedResult[0], result[0])
		return
	}

	if result[1] != expectedResult[1] {
		fmt.Printf("FAILURE: expected result[1] = %v, actual result[1] = %v \n", expectedResult[1], result[1])
		return
	}
}

func test1() {
	arr := []int{2, 7, 11, 15}
	target := 9
	expected := []int{1, 2} // 1-based indexes!

	test(arr, target, expected)
}

func test2() {
	arr := []int{2, 3, 4}
	target := 6
	expected := []int{1, 3} // 1-based indexes!

	test(arr, target, expected)
}

func test3() {
	arr := []int{-1, 0}
	target := -1
	expected := []int{1, 2} // 1-based indexes! We have to return different indexes!

	test(arr, target, expected)
}

func main() {
	// 167. Two Sum II - Input Array Is Sorted
	test1()
	test2()
	test3()
}

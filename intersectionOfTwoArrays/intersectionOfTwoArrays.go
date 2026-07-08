package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	// since value range is just [0; 1000], we can use an array instead of map
	arr := make([]bool, 1001) // from 0 to 1000

	for _, v := range nums1 {
		arr[v] = true
	}

	result := make([]int, 0)

	for _, v := range nums2 {
		if arr[v] {
			arr[v] = false // do not count this element anymore
			result = append(result, v)
		}
	}

	return result
}

func test(a1 []int, a2 []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array 1: %v \n", a1)
	fmt.Printf("Array 2: %v \n", a2)

	result := intersection(a1, a2)

	fmt.Printf("Intersection: %v \n", result)
}

func test1() {
	a1 := []int{1, 2, 2, 1}
	a2 := []int{2, 2}

	test(a1, a2)
}

func test2() {
	a1 := []int{4, 9, 5}
	a2 := []int{9, 4, 9, 8, 4}

	test(a1, a2)
}

func main() {
	// 349. Intersection of Two Arrays
	test1()
	test2()
}

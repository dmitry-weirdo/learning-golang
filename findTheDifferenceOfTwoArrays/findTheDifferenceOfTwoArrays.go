package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	// since value range is just [-1000; 1000], we can use an array instead of map
	arr1 := make([]bool, 2001) // from -1000 to 1000
	arr2 := make([]bool, 2001) // from -1000 to 1000

	for _, v := range nums1 {
		arr1[getIndex(v)] = true
	}

	res0 := make([]int, 0)

	for _, v := range nums2 {
		i := getIndex(v)

		arr2[i] = true

		if !arr1[i] {
			arr1[i] = true // do not count this element anymore
			res0 = append(res0, v)
		}
	}

	res1 := make([]int, 0)

	for _, v := range nums1 {
		i := getIndex(v)

		if !arr2[i] {
			arr2[i] = true // do not count this element anymore
			res1 = append(res1, v)
		}
	}

	return [][]int{res1, res0}
}

func getIndex(v int) int {
	// -1000 maps to 0
	// 1000 maps to 2000
	return v + 1000
}

func test(a1 []int, a2 []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array 1: %v \n", a1)
	fmt.Printf("Array 2: %v \n", a2)

	result := findDifference(a1, a2)

	fmt.Printf("Distinct integers in arr1 not present in arr2: %v \n", result[0])
	fmt.Printf("Distinct integers in arr2 not present in arr1: %v \n", result[1])
}

func test1() {
	a1 := []int{1, 2, 3}
	a2 := []int{2, 4, 6}

	test(a1, a2)
}

func test2() {
	a1 := []int{1, 2, 3, 3}
	a2 := []int{1, 1, 2, 2}

	test(a1, a2)
}

func main() {
	// 2215. Find the Difference of Two Arrays
	test1()
	test2()
}

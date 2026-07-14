package main

import "fmt"

var resultGlobal [][]int

func subsets(nums []int) [][]int {
	resultGlobal = make([][]int, 0)

	path := make([]int, 0)

	dfs(nums, 0, &path)

	return resultGlobal
}

func dfs(nums []int, i int, path *[]int) {
	fmt.Println()
	fmt.Println("==============================")
	fmt.Printf("i: %v, path: %v \n", i, *path)

	if i >= len(nums) {
		// reached the end of the array -> add the current to the result
		pathCopy := make([]int, len(*path))
		copy(pathCopy, *path)

		fmt.Printf("End of the array reached. Adding current path = %v to the result. \n", pathCopy)

		resultGlobal = append(resultGlobal, pathCopy)
		fmt.Printf("Result after adding: %v. \n", resultGlobal)

		return
	}

	// don't take the current element
	dfs(nums, i+1, path)

	// take the current element
	*path = append(*path, nums[i])
	dfs(nums, i+1, path)

	// remove the current element
	*path = (*path)[0 : len(*path)-1]
}

func test(nums []int) {
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("Array: %v \n", nums)

	result := subsets(nums)

	fmt.Printf("All possible subsets: %v \n", result)
}

func test1() {
	arr := []int{1, 2, 3}

	test(arr)
}

func test2() {
	arr := []int{0}

	test(arr)
}

func main() {
	// 78. Subsets
	test1()
	//test2()
}

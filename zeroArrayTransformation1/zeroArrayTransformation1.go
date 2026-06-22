package main

import "fmt"

func isZeroArray(nums []int, queries [][]int) bool {
	l := len(nums)

	diff := make([]int, l+1) // can be + 1

	// fill the diff array ( +1 at query start, - 1 at query end + 1 )
	for _, v := range queries {
		left := v[0]
		right := v[1]

		diff[left]++
		diff[right+1]-- // note it's right + 1, because the range is inclusive and ends only on the next index
	}

	fmt.Printf("Range array: %v \n", diff)

	prefixSums := make([]int64, l+2) // 0th + all sums of diffs
	prefixSums[0] = 0
	//prefixSums[l+1] = 666 // test that the last value is also calculated

	for i, v := range diff {
		prefixSums[i+1] = prefixSums[i] + int64(v)
	}

	fmt.Printf("Prefix sum array: %v \n", prefixSums)

	for i, n := range nums {
		if int64(n) > prefixSums[i+1] {
			fmt.Printf("nums[%v] = %v > prefixSums[%v] = %v. Returning false. \n", i, n, i+1, prefixSums[i+1])
			return false
		}
	}

	return true
}

func test1() {
	nums := []int{4, 3, 2, 1}

	queries := [][]int{
		{1, 3},
		{0, 2},
	}

	fmt.Printf("numbers array: %v \n", nums)
	fmt.Printf("queries array: %v \n", queries)

	possible := isZeroArray(nums, queries)

	fmt.Printf("Possible: %v \n", possible)
}

func test2() {
	nums := []int{3, 2, 1, 2}

	queries := [][]int{
		{0, 2},
		{1, 3},
		{2, 3},
	}

	fmt.Printf("numbers array: %v \n", nums)
	fmt.Printf("queries array: %v \n", queries)

	possible := isZeroArray(nums, queries)

	fmt.Printf("Possible: %v \n", possible)
}

func main() {
	//test1()
	test2()
}

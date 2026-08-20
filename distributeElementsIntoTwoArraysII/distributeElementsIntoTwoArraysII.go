package main

import (
	"fmt"
	"slices"
)

func resultArray(nums []int) []int {
	// todo: probably solving with Java's TreeMap instead of Fenwick tree would be faster,
	// since it already supports uniqueness (no need for complex "find unique sorted values" and no binary search for every number to find its position,
	// The counts will be put into map values.

	// Not the fastest solution
	// Runs in 90+ ms

	// Collect unique sorted values: O(2*n + n * log n).
	// Allocate 2 FT: O(2 * n) - for worst case when all values are unique.

	// For every value of n
	// Find index for the trees -> O(log n)
	// Find the range sum in the tree -> O(log n)
	// Add to one of the trees -> O(log n)
	// I.e. O(3 * n * log n)

	// Reverse the right part of the array - O(n) worst case (probably O(n / 2) actually)

	// Total:
	// 2 * n + (n * log n) + 2 * n + (3 * n * log n) + n =
	// = 5 * n + (4 * n * log n)
	// It's still O(n * log n), but works pretty slow.
	return resultArray_twoFenwickTrees(nums)
}

func resultArray_twoFenwickTrees(nums []int) []int {
	n := len(nums)

	uniqueSortedValues := uniqueSortedValuesInArray(nums)
	uniqueValuesCount := len(uniqueSortedValues) // will be <= 10^5

	//fmt.Printf("Unique sorted values (total %v values): %v \n", uniqueValuesCount, uniqueSortedValues)

	// FT initialized with empty arrays, i.e. all sums are initially 0
	// Both trees have the same size.
	// Every value in FT tree array contains count of the values put into this array.
	// We're adding 1 for every added value.
	// To find position of value in the array, we're using the binary search on unique values
	t1 := createEmptyFenwickTree(uniqueValuesCount)
	t2 := createEmptyFenwickTree(uniqueValuesCount)

	// add first values to tree1 and tree2 (i.e. set counts of these values to 1)
	t1.Add(getIndexForFenwickTree(uniqueSortedValues, nums[0]), 1)
	t2.Add(getIndexForFenwickTree(uniqueSortedValues, nums[1]), 1)

	// 2 arrays arr1 and arr2 to save the appended values
	// We can make it one array and then reverse the 2nd, as in "3069. Distribute Elements Into Two Arrays I"
	result := make([]int, n)
	result[0] = nums[0]
	result[n-1] = nums[1]

	// left goes (left -> right)
	left := 0

	// right goes (right -> left)
	right := n - 1

	a1TotalElements := 1
	a2TotalElements := 1

	for i := 2; i < n; i++ {
		value := nums[i]

		treeIndex := getIndexForFenwickTree(uniqueSortedValues, value)

		// In Fenwick trees, we're counting elements <= value.
		// We need to subtract this value from total element in this array,
		// totalElementsGreater = totalElements - totalElementsLessOrEqual
		a1LessOrEqualCount := t1.RangeSum(1, treeIndex)
		a2LessOrEqualCount := t2.RangeSum(1, treeIndex)

		a1GreaterCount := a1TotalElements - a1LessOrEqualCount
		a2GreaterCount := a2TotalElements - a2LessOrEqualCount

		// decide whether we append to a1 or a2
		addToA1 := true

		if a1GreaterCount > a2GreaterCount { // add to a1
			addToA1 = true
		} else if a1GreaterCount < a2GreaterCount { // add to a2
			addToA1 = false
		} else { // a1GreaterCount == a2GreaterCount -> add to array with fewer values
			if a1TotalElements <= a2TotalElements { // if lengths of a1 and a2 are equal -> append to a1
				addToA1 = true
			} else { // a1 is longer -> append to a2
				addToA1 = false
			}
		}

		if addToA1 { // add to a1
			left++
			a1TotalElements++
			result[left] = value

			t1.Add(treeIndex, 1) // add 1 count of the current element to tree1
		} else { // add to a2
			right--
			a2TotalElements++
			result[right] = value

			t2.Add(treeIndex, 1) // add 1 count of the current element to tree1
		}
	}

	// reverse the part [right; n-1]
	slices.Reverse(result[right:])
	return result
}

func uniqueSortedValuesInArray(arr []int) []int {
	// overall complexity is O(n + n * log n + n) = O(2 * n + n * log n) = O(n * log n)

	clone := slices.Clone(arr) // don't change the original array, O(n)
	slices.Sort(clone)         // sorts in-place, O(n * log n)

	return slices.Compact(clone) // removes consecutive duplicates, O(n)
}

func searchExactValueLeftmost(arr []int, target int) int { // returns -1 if element is not found
	if len(arr) < 1 { // empty array -> nothing to search
		return -1
	}

	condition := func(x int) bool {
		//return x == target // this will NOT work, e.g for {1, 1, 2, 3, 3}, target = 1 we will jump right
		return x >= target
	}

	index := binarySearchGeneric(arr, 0, len(arr)-1, condition)

	if arr[index] != target {
		return -1
	}

	return index
}

func binarySearchGeneric(
	arr []int,
	left int, // usually it starts with 0, if we search in the complete array
	right int, // set len(arr) - 1 if you want to be within array. Set len(arr) if index after the array can be returned.
	condition func(int) bool, // we will find the leftmost index satisfying this condition within [left; right] range
) int {
	// todo: this method can return an incorrect value for the empty array

	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		if condition(arr[mid]) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func getIndexForFenwickTree(uniqueSortedValues []int, value int) int {
	return searchExactValueLeftmost(uniqueSortedValues, value) + 1 // FT has 1-based indexes
}

type FenwickTree struct {
	t []int
}

func createEmptyFenwickTree(n int) FenwickTree {
	t := make([]int, n+1) // t-tree is 1-indexed
	return FenwickTree{t: t}
}

func createFenwickTree(a []int) FenwickTree {
	n := len(a)

	t := make([]int, n+1) // t-tree is 1-indexed
	copy(t[1:], a)        // t[0] is not used, we copy from pos 1

	for i := 1; i <= n; i++ {
		j := i + lsb(i) // add to the LSB of i

		if j <= n { // no overflow of array size
			t[j] += t[i]
		}
	}

	return FenwickTree{t: t}
}

func (ft *FenwickTree) prefixSum(i int) int { // sums elements with indexes [1; i], 1-based, inclusive
	sum := 0

	for i > 0 {
		sum += ft.t[i]

		// clear (set to 0) the least significant bit of i
		i &= ^lsb(i)
	}

	return sum
}

func (ft *FenwickTree) RangeSum(i, j int) int { // 1-based indexes, returns sum of [i, j] inclusive
	if i > j {
		panic(fmt.Sprintf("Incorrect range. i = %v > j = %v.", i, j))
	}

	return ft.prefixSum(j) - ft.prefixSum(i-1) // sum[i - 1] to include i
}

func (ft *FenwickTree) Set(i, value int) {
	currentValue := ft.RangeSum(i, i) // ft[i] is NOT value of the array, it's a partial sum [i - lsb(i) + 1, i]

	delta := value - currentValue // calc the value to add

	ft.Add(i, delta)
}

func (ft *FenwickTree) Add(i, delta int) {
	for i < len(ft.t) {
		ft.t[i] += delta

		// add 1 to the LSB, i.e. propagate to parents
		i += lsb(i)
	}
}

func lsb(x int) int { // least significant bit
	return x & -x
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text

	result := resultArray(arr)

	fmt.Printf("Transformed/distributed array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		[]int{2, 1, 3, 3},
		[]int{2, 3, 1, 3},
	)
}

func test2() {
	test(
		[]int{5, 14, 3, 1, 2},
		[]int{5, 3, 1, 2, 14},
	)
}

func test3() {
	test(
		[]int{3, 3, 3, 3},
		[]int{3, 3, 3, 3},
	)
}

func main() {
	// 3072. Distribute Elements Into Two Arrays II
	test1()
	test2()
	test3()
}

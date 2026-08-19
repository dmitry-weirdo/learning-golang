package binarySearchCommon

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

func searchExactValueRightmost(arr []int, target int) int { // returns -1 if element is not found
	if len(arr) < 1 { // empty array -> nothing to search
		return -1
	}

	condition := func(x int) bool {
		return x > target
	}

	// right is len(arr), so we'll jump over the array (to the insertion point)
	index := binarySearchGeneric(arr, 0, len(arr), condition)

	if index == 0 { // target smaller than the 0-th element of the array
		return -1
	}

	// rightmost is the previous from what we sought for
	index--

	if arr[index] != target {
		return -1
	}

	return index
}

func searchInsertPosition(arr []int, target int) int {
	condition := func(x int) bool {
		return x >= target
	}

	return binarySearchGeneric(
		arr,
		0,
		len(arr), // insert position can be after the end of the array
		condition,
	)
}

func binarySearchGeneric(
	arr []int, // todo: we can also generalize the type in the array
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

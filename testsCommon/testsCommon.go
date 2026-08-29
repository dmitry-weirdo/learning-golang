package testsCommon

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use ListNode without prefix
	"demo/matrixCommon"
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
	"strconv"
)

// common test formats to copy-paste just from this file instead of searching by many files

// ==================== Linked List functions ==================== //
func ListToListFunction(head *ListNode) *ListNode {
	return head
}

func testListToList(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := ListToListFunction(list) // todo: replace with your function
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with removed duplicates: \n") // todo: replace with your text
	listsCommon.PrintList(result)

	fmt.Printf("Result as array: %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func TwoListsToListFunction(l1, l2 *ListNode) *ListNode {
	if l1 != nil {
		return l1
	}

	return l2
}

func testTwoListsToList(a1 []int, a2 []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("===========================")

	l1 := listsCommon.ArrayToList(a1)
	l2 := listsCommon.ArrayToList(a2)

	fmt.Println("List 1:")
	listsCommon.PrintList(l1)

	fmt.Println("List 2:")
	listsCommon.PrintList(l2)

	result := TwoListsToListFunction(l1, l2) // todo: replace with your function
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Println("Sum list:") // todo: replace with your text
	listsCommon.PrintList(result)

	fmt.Printf("Result as array: %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func ListToBooleanFunction(head *ListNode) bool {
	return head != nil
}

func testListToBoolean(values []int, expectedResult bool) {
	fmt.Println()
	fmt.Println("===========================")

	list := listsCommon.ArrayToList(values)

	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("List: ")
	listsCommon.PrintList(list)

	result := ListToBooleanFunction(list) // todo: replace with your function

	fmt.Printf("List of %v is a palindrome: %v \n", values, result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func ListToIntArrayFunction(head *ListNode) []int {
	return []int{head.Val}
}

func testListToIntArray(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n") // todo: replace with your text if required
	listsCommon.PrintList(list)

	result := ListToIntArrayFunction(list) // todo: replace with your function

	fmt.Printf("Result: %v \n", result)
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

// ==================== []int int arrays functions ==================== //
func IntArrayToIntArrayFunction(arr []int) []int {
	return arr
}

func testIntArrayToIntArray(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text

	result := IntArrayToIntArrayFunction(arr) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
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

func IntArrayToIntMatrixFunction(arr []int) [][]int {
	return make([][]int, len(arr))
}

func testIntArrayToIntMatrix(arr []int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text

	result := IntArrayToIntMatrixFunction(arr) // todo: replace with your function

	fmt.Printf("Result: \n") // todo: replace with your text
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

func IntArrayToIntFunction(arr []int) int {
	return len(arr)
}

func testIntArrayToInt(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := IntArrayToIntFunction(arr) // todo: update to your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntArrayToBoolFunction(arr []int) bool {
	return len(arr) > 0
}

func testIntArrayToBool(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := IntArrayToBoolFunction(arr) // todo: update to your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntArrayToPairsArrayFunction(arr []int) [][]int {
	return [][]int{{len(arr), 1}}
}

func testIntArrayToPairsArrayFunction(arr []int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := IntArrayToPairsArrayFunction(arr)

	fmt.Printf("Pairs array:   %v \n", result) // todo: replace with your text if required
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := range expectedResult {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] { // we only compare pairs, i.e. [0] and [1] elements
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}

// ==================== []byte byte arrays functions ==================== //
func ByteArrayToByteFunction(arr []byte) byte {
	return byte(len(arr))
}

func testByteArrayToByte(arr []byte, expectedResult byte) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %s \n", arr) // %s used to display as characters, it will be just one string without spaces // todo: replace with your text if required

	result := ByteArrayToByteFunction(arr) // todo: update to your function

	fmt.Printf("Result: %c \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %c \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %c, actual result = %c \n", expectedResult, result)
	}
}

// ==================== Int functions ==================== //

func IntToBoolFunction(x int) bool {
	return x > 666
}

func testIntToBool(x int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := IntToBoolFunction(x) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntToByteFunction(x int) byte { // if we need to return a 1-byte char in Go
	return byte(x)
}

func testIntToByte(x int, expectedResult byte) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := IntToByteFunction(x) // todo: replace with your function

	fmt.Printf("Result: %c \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %c \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %c, actual result = %c \n", expectedResult, result)
	}
}

func IntToIntFunction(x int) int {
	return x + 1
}

func testIntToInt(x int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := IntToIntFunction(x) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntToStringFunction(x int) string {
	return strconv.Itoa(x)
}

func testIntToString(x int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := IntToStringFunction(x) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntToIntArrayFunction(x int) []int {
	return []int{x}
}

func testIntToIntArray(x int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text

	result := IntToIntArrayFunction(x) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
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

func IntToIntMatrixFunction(x int) [][]int {
	return make([][]int, x)
}

func testIntToIntMatrix(x int, expectedResult [][]int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := IntToIntMatrixFunction(x) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

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

func IntToStringMatrixFunction(x int) [][]string {
	return make([][]string, x)
}

func testIntToStringMatrix(x int, expectedResult [][]string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := IntToStringMatrixFunction(x) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

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

// ==================== String functions ==================== //
func StringToByteFunction(s string) byte { // if we need to return a 1-byte char in Go
	return s[0]
}

func testStringToByte(s string, expectedResult byte) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text if required

	result := StringToByteFunction(s) // todo: replace with your function

	fmt.Printf("Result: %c \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %c \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %c, actual result = %c \n", expectedResult, result)
	}
}

func StringToIntFunction(s string) int {
	return len(s)
}

func testStringToInt(s string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text if required

	result := StringToIntFunction(s) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func StringToStringFunction(s string) string {
	return s
}

func testStringToString(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text if required

	result := StringToStringFunction(s) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func StringToBoolFunction(s string) bool {
	return len(s) > 666
}

func testStringToBool(s string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text if required

	result := StringToBoolFunction(s) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func StringToIntArrayFunction(s string) []int {
	return []int{len(s)}
}

func testStringToIntArray(s string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text

	result := StringToIntArrayFunction(s) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
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

func StringToStringArrayFunction(s string) []string {
	return []string{s}
}

func testStringToStringArray(s string, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text

	result := StringToStringArrayFunction(s) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
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

// ==================== Tree functions ==================== //
func TwoIntArraysToTreeFunction(a1 []int, a2 []int) *TreeNode { // typically used to build a tree from 2 order of traversals
	return &TreeNode{a1[0] + a2[0], nil, nil}
}

func testTwoIntArraysToTree(preorder []int, inorder []int, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Preorder traversal: %v \n", preorder) // todo: replace with your text if required
	fmt.Printf("Inorder traversal: %v \n", inorder)   // todo: replace with your text if required

	result := TwoIntArraysToTreeFunction(preorder, inorder) // todo: replace with your function

	fmt.Printf("Built tree: \n") // todo: replace with your text if required
	trees.PrintTreeTopDown(result)

	resultAsArray := trees.TreeToArray(result)
	fmt.Printf("Tree as array:   %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func IntArrayToTreeFunction(arr []int) *TreeNode { // typically used to build a tree from 2 order of traversals
	return &TreeNode{len(arr), nil, nil}
}

func testIntArrayToTree(arr []int, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Sorted array: %v \n", arr) // todo: replace with your text if required

	result := IntArrayToTreeFunction(arr) // todo: replace with your function

	fmt.Printf("Built balanced BST: \n") // todo: replace with your text if required
	trees.PrintTreeTopDown(result)

	resultAsArray := trees.TreeToArray(result)
	fmt.Printf("Tree as array:   %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

// ==================== Matrix functions - often used for graphs ==================== //
func IntMatrixToIntMatrixFunction(m [][]int) [][]int {
	return m
}

func testIntMatrixToIntMatrix(m [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: \n") // todo: replace with your text if required
	matrixCommon.PrintIntMatrix(m)

	result := IntMatrixToIntMatrixFunction(m) // todo: replace with your function

	fmt.Printf("Result: \n") // todo: replace with your text
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

func IntMatrixToBoolArray(m [][]int) []bool { // e.g. check whether every pair/interval is matching
	return make([]bool, len(m))
}

func testIntMatrixToBoolArray(m [][]int, expectedResult []bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m) // todo: replace with your text if required

	result := IntMatrixToBoolArray(m) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

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

func IntMatrixToIntArray(m [][]int) []int { // e.g. find a path within a matrix
	return make([]int, len(m))
}

func testIntMatrixToIntArray(m [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m) // todo: replace with your text if required

	result := IntMatrixToIntArray(m) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

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

func IntMatrixToStringArray(m [][]int) []string { // e.g. find a path within a matrix
	return make([]string, len(m))
}

func testIntMatrixToStringArray(m [][]int, expectedResult []string) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m) // todo: replace with your text if required

	result := IntMatrixToStringArray(m) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

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

func IntMatrixToInt(m [][]int) int { // calculate some element or length of path in the matrix
	return len(m)
}

func testIntMatrixToInt(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m) // todo: replace with your text if required

	result := IntMatrixToInt(m) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntMatrixToBool(m [][]int) bool { // e.g. check whether something is possible in the graph (matrix of edges)
	return len(m) > 2
}

func testIntMatrixToBool(m [][]int, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m) // todo: replace with your text if required

	result := IntMatrixToBool(m) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntMatrixToFloat64(m [][]int) float64 { // calculate some element or length of path in the matrix
	return float64(len(m))
}

func testIntMatrixToFloat64(m [][]int, expectedResult float64) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m) // todo: replace with your text if required

	result := IntMatrixToFloat64(m) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

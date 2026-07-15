package main

import (
	"container/list"
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	fmt.Printf("Split path by /: %v \n", split)

	stack := list.New()

	for _, v := range split {
		fmt.Printf("Next value from the split: %v \n", v)

		// current directory or empty space between spaces -> do nothing
		if v == "" || v == "." {
			continue
		}

		// parent directory -> pop last directory from the stack
		if v == ".." {
			if stack.Len() <= 0 { // if already a root path -> do nothing
				fmt.Printf("\"..\" command given, but nothing to pop from the stack. Staying at the root directory. \n")
				continue
			}

			removed := stack.Remove(stack.Back())

			fmt.Printf("Removed parent path \"%v\" from the stack. \n", removed)
			continue
		}

		// usual directory -> add to the stack
		stack.PushBack(v)
	}

	pathArray := make([]string, 0)

	for e := stack.Front(); e != nil; e = e.Next() {
		pathArray = append(pathArray, e.Value.(string))
	}

	return "/" + strings.Join(pathArray, "/")
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := simplifyPath(s)

	fmt.Printf("Simplified path: %v \n", result)
	fmt.Printf("Expected simplified path: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "/home/user/Documents/../Pictures"
	expected := "/home/user/Pictures"

	test(s, expected)
}

func test2() {
	s := "/.../a/../b/c/../d/./"
	expected := "/.../b/d"

	test(s, expected)
}

func test3() {
	s := "/../"
	expected := "/"

	test(s, expected)
}

func test4() {
	s := "/"
	expected := "/"

	test(s, expected)
}

func test5() {
	s := "/neetcode/practice//...///../courses"
	expected := "/neetcode/practice/courses"

	test(s, expected)
}

func test6() {
	s := "/..//_home/a/b/..///"
	expected := "/_home/a"

	test(s, expected)
}

func main() {
	// 71. Simplify Path
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}

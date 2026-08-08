package main

import "fmt"

//const NO_PARENT = "-" // nil won't work, let's use the special string

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	m := make(map[string]string) // region to parent

	for _, v := range regions {
		parent := v[0]

		// save this `if` for time optimization when building the tree -> map will return an empty string "" if there is nothing for the root key
		/*
			if _, ok := m[parent]; !ok { // put explicit special key for the root so that it also has a key
				m[parent] = NO_PARENT
			}
		*/

		for j := 1; j < len(v); j++ {
			child := v[j]

			m[child] = parent
		}
	}

	fmt.Printf("Child to parent map: \n%v \n", m)

	return lowestCommonAncestor(m, region1, region2)
}

func lowestCommonAncestor(m map[string]string, p string, q string) string {
	// this is the same logic as in "1650. Lowest Common Ancestor of a Binary Tree III"

	// if we want O(1) space optimization logic, we can:
	// - calculate depths of both nodes
	// - go from the deeper node up to align the depths
	// - move up step by step until the nodes are the same
	// It will still be O(h) time complexity

	path := make(map[string]bool)

	// go from P to the root, collect all the values in the path into a map
	current := p

	for current != "" { // while root not reached
		path[current] = true
		current = m[current]
	}

	// go from Q to the root. The first node in the path that was already in the path of P is the LCA
	current = q

	for current != "" { // while root not reached
		if _, ok := path[current]; ok {
			return current
		}

		current = m[current]
	}

	// this must never happen, at least the root should be the LCA
	return ""
}

func test(regions [][]string, r1, r2 string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Regions: %v \n", regions)
	fmt.Printf("Region 1: %v \n", r1)
	fmt.Printf("Region 2: %v \n", r2)

	result := findSmallestRegion(regions, r1, r2)

	fmt.Printf("LCA of \"%v\" and \"%v\": %v \n", r1, r2, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	regions := [][]string{
		{"Earth", "North America", "South America"},
		{"North America", "United States", "Canada"},
		{"United States", "New York", "Boston"},
		{"Canada", "Ontario", "Quebec"},
		{"South America", "Brazil"},
	}

	// map[
	//Boston:United States
	//Brazil:South America
	//Canada:North America
	//Earth:-
	//New York:United States
	//North America:Earth
	//Ontario:Canada
	//Quebec:Canada
	//South America:Earth
	//United States:North America
	//]

	region1 := "Quebec"
	region2 := "New York"
	expected := "North America"

	test(regions, region1, region2, expected)
}

func test2() {
	regions := [][]string{
		{"Earth", "North America", "South America"},
		{"North America", "United States", "Canada"},
		{"United States", "New York", "Boston"},
		{"Canada", "Ontario", "Quebec"},
		{"South America", "Brazil"},
	}

	region1 := "Canada"
	region2 := "South America"
	expected := "Earth"

	test(regions, region1, region2, expected)
}

func main() {
	// 1257. Smallest Common Region
	test1()
	test2()
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	// n - count of strings
	// k - max (or average) len of the strings

	// trivial solution: sort chars in all strings, then combine by sorted
	// it will be O(k * log k) for every string
	// O(n * k * log k) in total

	// if we hash every string instead of sorting,
	// it will be O(k) for every string
	// O(n * k) in total
	// calculating a hash string is 26 every time
	// (k + 26) vs (k * log k), log will be bigger for k >= 14
	// todo: we can select the method based on N, but this is an EXTREME optimization

	m := make(map[string][]string, 0)

	for _, s := range strs {
		frequencyString := getFrequencyString(s)

		if stringsByFrequency, ok := m[frequencyString]; ok {
			m[frequencyString] = append(stringsByFrequency, s)
		} else { // new frequencyString
			m[frequencyString] = []string{s}
		}
	}

	// conversion of map to array is additional time :(
	result := make([][]string, len(m))

	i := 0
	for _, v := range m {
		result[i] = v

		i++
	}

	return result
}

func getFrequencyString(s string) string {
	// frequency array from 'a' to 'z'
	f := make([]int, 26)

	for _, ch := range s {
		index := ch - 'a'
		f[index]++
	}

	var sb strings.Builder

	for i, v := range f {
		if v <= 0 { // only append chars from string
			continue
		}

		ch := rune('a' + i)
		sb.WriteRune(ch)
		sb.WriteString(strconv.Itoa(v))
	}

	frequencyString := sb.String()

	fmt.Println()
	fmt.Printf("String: %v \n", s)
	fmt.Printf("Frequency array: %v \n", f)
	fmt.Printf("Frequency string: %v \n", frequencyString)

	return frequencyString
}

func test(strings []string) {
	fmt.Println()
	fmt.Println("===========================")

	fmt.Printf("Strings: %v \n", strings)

	anagramGroups := groupAnagrams(strings)

	fmt.Printf("Strings grouped as anagrams: %v \n", anagramGroups)
}

func test1() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	test(s)
}

func main() {
	test1()
}

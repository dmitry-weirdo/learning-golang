package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	return longestCommonPrefix_trie(strs)
}

func longestCommonPrefix_trie(strs []string) string {
	trie := Constructor()

	for _, s := range strs {
		if s == "" {
			return ""
		}

		trie.Insert(s)
	}

	// search by any string, since we need a common prefix of them all
	return trie.SearchPrefix(strs[0])
}

type TrieNode struct {
	children map[byte]*TrieNode
	// todo: key not necessary required? But I would like it to be
	key  byte
	word bool // marks the end of the word
}

func (this *TrieNode) Size() int {
	return len(this.children)
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
			key:      byte(' '), // special unique key for the root
			word:     false,
		},
	}
}

func (this *Trie) Insert(word string) {
	current := this.root

	for _, v := range word {
		char := byte(v)

		if _, ok := current.children[char]; !ok {
			current.children[char] = &TrieNode{
				children: make(map[byte]*TrieNode),
				key:      char,
				word:     false,
			}
		}

		current = current.children[char]
	}

	// mark the end of the word
	current.word = true
}

func (this *Trie) Search(word string) bool {
	current := this.root

	for _, v := range word {
		char := byte(v)

		if _, ok := current.children[char]; !ok {
			return false
		}

		current = current.children[char]
	}

	// we return true only if it is the end of the word
	return current.word
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root

	for _, v := range prefix {
		char := byte(v)

		if _, ok := current.children[char]; !ok {
			return false
		}

		current = current.children[char]
	}

	// for prefix search, we return true regardless of this node is word or not
	return true
}

func (this *Trie) SearchPrefix(s string) string { // custom implementation for "14. Longest Common Prefix"
	var sb strings.Builder

	// root does NOT correspond to any letter
	node := this.root

	for _, v := range s {
		ch := byte(v)

		if node.word || // any word ended here -> stop
			node.Size() != 1 { // we must have just 1 child, else it's a different path
			return sb.String()
		}

		// next single character must be the char of the current string
		if _, ok := node.children[ch]; !ok {
			return sb.String()
		}

		// proceed to the next single character
		node = node.children[ch]

		sb.WriteByte(ch)
	}

	return sb.String()
}

func test(arr []string, expectedResult string) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := longestCommonPrefix(arr)

	fmt.Printf("Longest common prefix of all strings in the array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]string{"flower", "flow", "flight"},
		"fl",
	)
}

func test2() {
	test(
		[]string{"dog", "racecar", "car"},
		"",
	)
}

func test3() {
	test(
		[]string{"", "aaa", "aaaa"},
		"",
	)
}

func main() {
	// 14. Longest Common Prefix
	test1()
	test2()
	test3()
}

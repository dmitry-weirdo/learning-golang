package main

import "fmt"

func longestWord(words []string) string {
	// passes in 21-26 ms. Using many hashmaps (1 for every node) is not cheap
	return longestWord_trie(words)
}

func longestWord_trie(words []string) string {
	// put all the words into Trie
	t := NewTrie()

	for _, s := range words {
		t.Insert(s)
	}

	// we need to do DFS (or BFS) on Trie, getting the max depth and selecting the lexicographically smaller of the same maxDepth
	// we're only iterating the nodes with word == true
	result := ""

	var dfs func(n *TrieNode, s []byte)

	dfs = func(n *TrieNode, s []byte) {
		if !t.IsRoot(n) && !n.word {
			// not a word node and not the root node -> stop DFS
			return
		}

		if !t.IsRoot(n) {
			s = append(s, n.key)
		}

		if len(s) > len(result) {
			result = string(s)
		} else if len(s) == len(result) {
			// same length -> change result only if it's lexicographically smaller
			str := string(s)

			if str < result {
				result = str
			}
		}

		for _, child := range n.children {
			if child.word {
				dfs(child, s)
			}
		}

		// backtrack the current char
		if !t.IsRoot(n) {
			s = s[:len(s)-1]
		}
	}

	s := make([]byte, 0) // dynamically add/remove at the end of the string
	dfs(t.GetRoot(), s)

	return result
}

type TrieNode struct {
	children map[byte]*TrieNode
	key      byte
	word     bool // word == true marks the end of the word
}

func (this *TrieNode) Size() int {
	return len(this.children)
}

type Trie struct {
	root    *TrieNode
	rootKey byte
}

func NewTrie() Trie {
	// todo: make overrideable in the constructor
	const ROOT_KEY = byte(' ') // special unique key for the root

	return Trie{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
			key:      ROOT_KEY, // special unique key for the root
			word:     false,
		},
		rootKey: ROOT_KEY,
	}
}

func (this *Trie) GetRootKey() byte {
	return this.rootKey
}

func (this *Trie) GetRoot() *TrieNode {
	return this.root
}

func (this *Trie) IsRoot(node *TrieNode) bool {
	return node.key == this.rootKey
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

func test(arr []string, expectedResult string) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := longestWord(arr)

	fmt.Printf("Longest word that can be build char-by-char: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]string{"w", "wo", "wor", "worl", "world"},
		"world",
	)
}

func test2() {
	test(
		[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
		"apple", // "apple" < "apply"
	)
}

func test3() {
	test(
		[]string{"a", "b", "cc"},
		"a",
	)
}

func main() {
	// 720. Longest Word in Dictionary
	test1()
	test2()
	test3()
}

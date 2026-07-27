package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	// todo: key not necessary required? But I would like it to be
	key  byte
	word bool // marks the end of the word
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
			key:      byte(' '),
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

func check(expectedResult bool, f func() bool) {
	result := f()

	if result != expectedResult {
		errorMessage := fmt.Sprintf("expected result: %v, actual result: %v", expectedResult, result)
		panic(errorMessage)
	}
}

func test() {
	t2 := Constructor()
	t := &t2

	t.Insert("apple")
	check(true, func() bool { return t.Search("apple") })
	check(false, func() bool { return t.Search("app") })
	check(true, func() bool { return t.StartsWith("app") })

	t.Insert("app")
	check(true, func() bool { return t.Search("app") })
}

func main() {
	// 208. Implement Trie (Prefix Tree)
	test()
}

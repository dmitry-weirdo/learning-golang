package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	// passes in 0 ms because of small constraints, but can be optimized
	return wordBreak_trieDP_naive(s, wordDict)
}

func wordBreak_trieDP_naive(s string, dict []string) bool {
	t := NewTrie()

	for _, v := range dict { // add all dictionary words to trie
		t.Insert(v)
	}

	n := len(s)

	// DP[i] - whether we can get to s[i-1] with words from the dictionary
	dp := make([]bool, n+1)
	dp[0] = true

	for i := range s {
		if !dp[i] { // we cannot get to position i -> skip it
			continue
		}

		for j := i; j < n; j++ { // current position is valid -> check from it to the end of S whether S[i:j] is in the dictionary
			//fmt.Printf("i: %v, j: %v. Checked string: \"%v\". \n", i, j, s[i:j+1])

			// todo: we shouldn't iterate from Trie root all the time, we should just proceed with 1 position

			if !t.StartsWith(s[i : j+1]) { // there is no prefix for substring -> stop further iteration from this position [i]
				break
			}

			if t.Search(s[i : j+1]) {
				dp[j+1] = true
			}
		}
	}

	//fmt.Printf("DP: %v \n", dp)

	// whether we can get to the last character of S
	return dp[n]
}

// ============================== Trie start ============================== //
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

// ============================== Trie end ============================== //

func test(s string, dict []string, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Dictionary: %v \n", dict)

	result := wordBreak(s, dict)

	fmt.Printf("String can be split to dictionary words: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("leetcode", []string{"leet", "code"}, true)
}

func test2() {
	test("applepenapple", []string{"apple", "pen"}, true)
}

func test3() {
	test("catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false)
}

func main() {
	// 139. Word Break
	test1()
	test2()
	test3()
}

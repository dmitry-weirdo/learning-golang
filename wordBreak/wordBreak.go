package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	// do not iterate the tree from the root, go char by char
	// Notably it runs in 0-2 ms, but it's probably the LeetCode runtime instability
	return wordBreak_trieDP_optimized(s, wordDict)

	// passes in 0 ms because of small constraints, but can be optimized
	//return wordBreak_trieDP_naive(s, wordDict)
}

func wordBreak_trieDP_optimized(s string, dict []string) bool {
	t := NewTrie()

	for _, v := range dict { // add all dictionary words to trie
		t.Insert(v)
	}

	n := len(s)

	// DP[i] - whether we can get to s[i-1] with words from the dictionary
	dp := make([]bool, n+1)
	dp[0] = true

	for i := range s {
		//fmt.Printf("i: %v \n", i)

		if !dp[i] { // we cannot get to position i -> skip it
			continue
		}

		node := t.GetRoot()

		for j := i; j < n; j++ { // current position is valid -> check from it to the end of S whether S[i:j] is in the dictionary
			ch := s[j]

			//fmt.Printf("i: %v, j: %v. Checked string: \"%v\". Ch[j]: %c. \n", i, j, s[i:j+1], ch)

			if !node.HasChild(ch) { // there is no prefix for substring -> stop further iteration from this position [i]
				break
			}

			// go to the node corresponding to s[j]
			node = node.children[ch]

			// we reached a node that maps to a word -> mark the position [i + j] as valid word end.
			if node.word {
				if j+1 == n { // reached the end of the string with a word end -> return true immediately
					return true
				}

				dp[j+1] = true
			}
		}
	}

	//fmt.Printf("DP: %v \n", dp)

	// whether we can get to the last character of S
	return dp[n]
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
		//fmt.Printf("i: %v \n", i)

		if !dp[i] { // we cannot get to position i -> skip it
			continue
		}

		for j := i; j < n; j++ { // current position is valid -> check from it to the end of S whether S[i:j] is in the dictionary
			//fmt.Printf("i: %v, j: %v. Checked string: \"%v\". \n", i, j, s[i:j+1])

			// !!! we shouldn't iterate from Trie root all the time, we should just proceed with 1 position
			// this is done in the wordBreak_trieDP_optimized version of this method

			if !t.StartsWith(s[i : j+1]) { // there is no prefix for substring -> stop further iteration from this position [i]
				break
			}

			if t.Search(s[i : j+1]) {
				if j+1 == n { // reached the end of the string with a word end -> return true immediately
					return true
				}

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

func (this *TrieNode) HasChild(ch byte) bool {
	_, ok := this.children[ch]
	return ok
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

		if !current.HasChild(char) {
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
	found, node := this.FindPrefixNode(word)

	// we return true only if it is the end of the word
	return found && node.word
}

func (this *Trie) StartsWith(prefix string) bool {
	// for prefix search, we return true regardless of this node is word or not
	found, _ := this.FindPrefixNode(prefix)
	return found
}

func (this *Trie) FindPrefixNode(prefix string) (found bool, node *TrieNode) {
	current := this.root

	for _, v := range prefix {
		char := byte(v)

		if !current.HasChild(char) {
			return false, nil
		}

		current = current.children[char]
	}

	// for prefix search, we return true regardless of this node is word or not
	return true, current
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

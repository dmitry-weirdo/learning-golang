package trieCommon

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

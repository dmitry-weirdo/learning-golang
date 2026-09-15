package trieCommon

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

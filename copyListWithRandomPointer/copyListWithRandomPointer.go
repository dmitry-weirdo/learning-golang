package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	return copyRandomList_HashMap_trivial(head)
}

func copyRandomList_HashMap_trivial(head *Node) *Node {
	// index is 0-based order in the linked list
	nodeToIndex := make(map[*Node]int)    // old list - node to index
	newIndexToNode := make(map[int]*Node) // new list - index to node

	index := 0
	n := head

	dummyHead := &Node{-1, nil, nil}
	currentTail := dummyHead
	newNode := dummyHead

	// first run - go via linked list, duplicate into the new nodes
	// Random is not yet set
	for n != nil {
		nodeToIndex[n] = index

		newNode = &Node{n.Val, nil, nil}
		currentTail.Next = newNode
		currentTail = newNode

		newIndexToNode[index] = newNode

		n = n.Next
		index++
	}

	n = head
	index = 0

	// second run - go via linked list, find to what index does random point
	// In the new list, set the Random of the currentIndex node to the new node with RandomIndex
	for n != nil {
		if n.Random != nil {
			randomIndex := nodeToIndex[n.Random]

			newIndexToNode[index].Random = newIndexToNode[randomIndex]
		}

		n = n.Next
		index++
	}

	return dummyHead.Next
}

func main() {
	// 138. Copy List with Random Pointer

	// todo: it requires new list methods for the new type and new testing method, we can add it and write tests
}

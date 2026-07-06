package main

import "fmt"

type MyLinkedList struct {
	head  *ListNode
	tail  *ListNode
	size  int
	debug bool
}

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func Constructor() MyLinkedList {
	dummyHead := &ListNode{
		val:  -666,
		prev: nil,
		next: nil,
	}

	dummyTail := &ListNode{
		val:  777,
		prev: dummyHead,
		next: nil,
	}

	dummyHead.next = dummyTail

	return MyLinkedList{
		head:  dummyHead,
		tail:  dummyTail,
		size:  0,
		debug: false,
	}
}

func (this *MyLinkedList) PrintWithDummyNodes() {
	var node = this.head // dummyHead

	for node != nil {
		fmt.Printf("%d ", node.val)

		node = node.next
	}

	fmt.Printf(" | Size: %v", this.size)

	fmt.Println()
	fmt.Println()
}

func (this *MyLinkedList) PrintWithoutDummyNodes() {
	// skip dummy head
	var node = this.head.next

	for node != this.tail { // skip dummy tail
		fmt.Printf("%d ", node.val)

		node = node.next
	}

	fmt.Printf(" | Size: %v", this.size)

	fmt.Println()
	fmt.Println()
}

func (this *MyLinkedList) GetAtIndex(index int) *ListNode { // helper function, not in the ticket
	if index >= this.size {
		if this.debug {
			fmt.Printf("Index [%v] is greater than size %v. Returning nil node. \n", index, this.size)
		}

		return nil
	}

	i := 0
	n := this.head.next

	for i < index {
		i++
		n = n.next
	}

	return n
}

func (this *MyLinkedList) Get(index int) int {
	n := this.GetAtIndex(index)

	if n == nil {
		if this.debug {
			fmt.Printf("Index [%v] is greater than size %v. Cannot get element at index [%v]. \n", index, this.size, index)
		}

		return -1
	}

	if this.debug {
		fmt.Printf("Element[%v] = %v. \n", index, n.val)
		this.PrintWithDummyNodes()
	}

	return n.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// we will insert AFTER dummyHead
	firstElement := this.head.next

	newHead := &ListNode{
		val:  val,
		prev: this.head, // dummyHead
		next: firstElement,
	}

	this.head.next = newHead // append new element after dummyHead
	firstElement.prev = newHead

	this.size++

	if this.debug {
		fmt.Printf("Added new element %v at head. \n", val)
		this.PrintWithDummyNodes()
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	// we will insert BEFORE dummyTail
	lastElement := this.tail.prev

	newTail := &ListNode{
		val:  val,
		prev: lastElement,
		next: this.tail, // dummyTail
	}

	this.tail.prev = newTail
	lastElement.next = newTail

	this.size++

	if this.debug {
		fmt.Printf("Added new element %v at tail. \n", val)
		this.PrintWithDummyNodes()
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.size { // special case -> insert at tail
		if this.debug {
			fmt.Printf("Index [%v] is equal to size %v. Adding element %v at index %v as a new tail. \n", index, this.size, val, index)
		}

		this.AddAtTail(val)

		return
	}

	n := this.GetAtIndex(index)

	if n == nil {
		if this.debug {
			fmt.Printf("Index [%v] is greater than size %v. Cannot add element %v at index %v. \n", index, this.size, val, index)
		}

		return
	}

	// we insert BEFORE this node
	newNode := &ListNode{
		val:  val,
		prev: n.prev,
		next: n, // insert BEFORE n
	}

	// change prev.next
	n.prev.next = newNode

	// change n.prev
	n.prev = newNode

	// increase size
	this.size++

	if this.debug {
		fmt.Printf("Added element %v at index [%v]. \n", val, index)
		this.PrintWithDummyNodes()
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	n := this.GetAtIndex(index)

	if n == nil {
		if this.debug {
			fmt.Printf("Index [%v] is greater than size %v. Cannot delete element at index %v. \n", index, this.size, index)
		}

		return
	}

	// skip removed node from its prev
	n.prev.next = n.next

	// skip removed node from its next
	n.next.prev = n.prev

	this.size--

	if this.debug {
		fmt.Printf("Deleted element %v at index [%v]. \n", n.val, index)
		this.PrintWithDummyNodes()
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func test1() {
	list := Constructor()
	list.debug = true

	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2) // linked list becomes 1->2->3
	list.Get(1)           // return 2
	list.DeleteAtIndex(1) // now the linked list is 1->3
	list.Get(1)           // return 3
}

func test2() {
	list := Constructor()
	list.debug = true

	list.AddAtHead(7)
	list.AddAtHead(2)
	list.AddAtHead(1)

	list.AddAtIndex(3, 0)

	list.DeleteAtIndex(2)

	list.AddAtHead(6)

	list.AddAtTail(4)

	list.Get(4)
	list.AddAtHead(4)
	list.AddAtIndex(5, 0)
	list.AddAtHead(6)
}

func main() {
	//test1()
	test2()
}

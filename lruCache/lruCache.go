package main

import (
	"container/list"
	"fmt"
)

type KeyValue struct {
	key   int // need to store the key to remove the LRU key from the map
	value int
}

type LRUCache struct {
	capacity int
	elements *list.List            // list contains *KeyValue, so that we can update the value in place
	m        map[int]*list.Element // key to element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		elements: list.New(),
		m:        make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok { // element not found
		return -1
	}

	// key used -> move the key to the beginning of the list
	el := this.m[key]
	this.elements.MoveToFront(el)

	return el.Value.(*KeyValue).value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok { // element not found -> add it
		el := this.elements.PushFront(&KeyValue{key, value})
		this.m[key] = el

		if this.elements.Len() > this.capacity { // remove the LRU element from the end of the list
			lruElement := this.elements.Back()
			this.elements.Remove(lruElement)

			delete(this.m, lruElement.Value.(*KeyValue).key)
		}
	} else { // key exists -> update its value
		el := this.m[key]
		keyValue := el.Value.(*KeyValue)
		keyValue.value = value

		// move the used element to the start
		this.elements.MoveToFront(el)
	}
}

func testPut(c LRUCache, key, value int) {
	fmt.Println()

	c.Put(key, value)
	fmt.Printf("Put value %v to key %v \n", value, key)

	updatedValue := c.Get(key)
	fmt.Printf("Updated value for key %v: %v \n", key, updatedValue)

	if updatedValue != value {
		fmt.Printf("FAILURE: expected updated value = %v, actual updated value = %v \n", updatedValue, value)
	}
}

func testGet(c LRUCache, key, expectedValue int) {
	fmt.Println()

	value := c.Get(key)
	fmt.Printf("Got value for key %v: %v \n", key, value)

	if value != expectedValue {
		fmt.Printf("FAILURE: expected value = %v, actual value = %v \n", expectedValue, value)
	}
}

func test1() {
	// ["LRUCache","put","put","get","put","put","get"]
	//	[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
	fmt.Println()
	fmt.Println("========================")

	c := Constructor(2)
	testPut(c, 2, 1)
	testPut(c, 2, 2)  // for key 2, change 1 to 2
	testGet(c, 2, 2)  // must be 1
	testPut(c, 1, 1)  // new key 1
	testPut(c, 4, 1)  // new key 4, removes the LRU key 2
	testGet(c, 2, -1) // must return -1
}

func test2() {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	fmt.Println()
	fmt.Println("========================")

	c := Constructor(2)
	testPut(c, 1, 1)
	testPut(c, 2, 2)
	testGet(c, 1, 1)
	testPut(c, 3, 3) // will remove key 2
	testGet(c, 2, -1)
	testPut(c, 4, 4) // will remove key 1
	testGet(c, 1, -1)
	testGet(c, 3, 3)
	testGet(c, 4, 4)
}

func main() {
	// 146. LRU Cache
	test1()
	test2()
}

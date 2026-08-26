package main

import (
	"container/list"
	"fmt"
)

type KeyValueFreq struct {
	key                  int // need to store the key to remove the LRU key from the map
	value                int
	frequencyNodeElement *list.Element // to just keep a link on element in the main map
}

type FrequencyNode struct {
	//name      string // just for printing usability
	frequency int
	elements  *list.List // List<*KeyValueFreq> list contains *KeyValueFreq, so that we can update the value and freq in place
}

type LFUCache struct {
	capacity      int
	totalElements int
	freqList      *list.List            // List<*FrequencyNode>. We need to keep existing frequencies ordered
	m             map[int]*list.Element // Key -> Element<*KeyValueFreq>. Element contains *KeyValueFreq. Element knows its key, value and frequency node element.
}

func Constructor(capacity int) LFUCache {
	head := FrequencyNode{
		//name:      "head",
		frequency: 0,          // so that "frequency: 1" node should be after it
		elements:  list.New(), // no elements in fake head
	}

	tail := FrequencyNode{
		//name:      "tail",
		frequency: -666,
		elements:  list.New(), // no elements in fake tail
	}

	frequencyList := list.New()
	frequencyList.PushFront(&head)
	frequencyList.PushBack(&tail)

	return LFUCache{
		capacity:      capacity,
		totalElements: 0,
		freqList:      frequencyList,
		m:             make(map[int]*list.Element),
	}
}

func (this *LFUCache) GetWithoutFrequencyUpdate(key int) int { // does NOT update frequency -> method is just for test usage
	if _, ok := this.m[key]; !ok { // element not found
		return -1
	}

	return this.m[key].Value.(*KeyValueFreq).value
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok { // element not found
		return -1
	}

	// key used -> increase its frequency, remove from old frequency, add new frequency
	currentValue := this.m[key].Value.(*KeyValueFreq).value

	this.IncreaseFrequencyForExistingKey(key, currentValue) // do NOT change the value

	return currentValue
}

func (this *LFUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok { // element not found -> add new element for this key
		// new key -> increase totalElements, potentially removed the LFU -> LRU element
		this.totalElements++

		if this.totalElements > this.capacity {
			this.RemoveLruElementFromLfuFrequency() // will decrease totalElements
		}

		this.AddNewKeyWithFrequencyOne(key, value)
	} else {
		// element found -> updated it to the new frequency, remove from the old frequency
		// totalElements not changed -> no removal
		this.IncreaseFrequencyForExistingKey(key, value)
	}
}

func (this *LFUCache) AddNewKeyWithFrequencyOne(key int, value int) {
	frequency0Element := this.freqList.Front()

	// get frequencyElement for frequency = 1
	newFrequencyElement := this.GetFrequencyElementForFrequencyPlusOne(frequency0Element)
	newFrequencyNode := newFrequencyElement.Value.(*FrequencyNode)

	newKeyValueFreq := &KeyValueFreq{
		key:                  key,
		value:                value,
		frequencyNodeElement: newFrequencyElement,
	}

	// add new KeyValueFreq to the beginning of newFrequencyNode.elements
	newKeyValueFreqElement := newFrequencyNode.elements.PushFront(newKeyValueFreq)

	// for the global O(1) access by key, set the new element for this key
	this.m[key] = newKeyValueFreqElement

	// there was no old frequency for this key -> we do not remove it from the old frequency
}

func (this *LFUCache) IncreaseFrequencyForExistingKey(key int, newValue int) {
	oldKeyValueFreqElement := this.m[key]
	oldKeyValueFreq := oldKeyValueFreqElement.Value.(*KeyValueFreq)

	oldFrequencyElement := oldKeyValueFreq.frequencyNodeElement

	newFrequencyElement := this.GetFrequencyElementForFrequencyPlusOne(oldFrequencyElement)
	newFrequencyNode := newFrequencyElement.Value.(*FrequencyNode)

	// todo: theoretically we can also move list.Element from oldFrequencyNode.elements to newFrequencyNode.elements. But it will lead to ugly and unsafe code.
	/*
		// if possible, reuse the same keyValueFreq
		newKeyValueFreq := &KeyValueFreq{
			key:                  key,
			value:                newValue,
			frequencyNodeElement: newFrequencyElement,
		}
	*/

	// we can reuse the same KeyValueFreq, just change its value and link to the FrequencyNode
	oldKeyValueFreq.value = newValue
	oldKeyValueFreq.frequencyNodeElement = newFrequencyElement

	// add new KeyValueFreq to the beginning of newFrequencyNode.elements
	newKeyValueFreqElement := newFrequencyNode.elements.PushFront(oldKeyValueFreq)

	// for the global O(1) access by key, set the new element for this key
	this.m[key] = newKeyValueFreqElement

	// remove from old frequency node
	this.RemoveElementFromFrequencyNode(oldFrequencyElement, oldKeyValueFreqElement)
}

func (this *LFUCache) GetFrequencyElementForFrequencyPlusOne(oldFrequencyElement *list.Element) *list.Element {
	// if frequency element for (frequency + 1) exists -> return it
	// if frequency element for (frequency + 1) does not exist -> adds it after oldFrequencyElement and returns it

	oldFrequencyNode := oldFrequencyElement.Value.(*FrequencyNode) // value of oldFrequencyElement
	oldFrequency := oldFrequencyNode.frequency

	newFrequency := oldFrequency + 1

	nextFrequencyElement := oldFrequencyElement.Next()
	nextFrequencyNode := nextFrequencyElement.Value.(*FrequencyNode)
	nextFrequency := nextFrequencyNode.frequency

	if nextFrequency == newFrequency {
		// (oldFrequency + 1) node exists -> return it
		return nextFrequencyElement
	}

	// newFrequency node does not exist -> add new frequencyNode for newFrequency
	newFrequencyNode := &FrequencyNode{
		//name:      this.GetFrequencyNodeName(newFrequency),
		frequency: newFrequency,
		elements:  list.New(),
	}

	// add the newFrequencyNode after the oldFrequencyNode
	newFrequencyElement := this.freqList.InsertAfter(newFrequencyNode, oldFrequencyElement)

	return newFrequencyElement
}

func (this *LFUCache) RemoveLruElementFromLfuFrequency() {
	//fmt.Printf("Total elements = %v > capacity = %v. Removing the LFU -> LRU element. \n", this.totalElements, this.capacity)

	// get LFU frequency
	leastFrequencyElement := this.freqList.Front().Next()
	leastFrequencyNode := leastFrequencyElement.Value.(*FrequencyNode)

	if leastFrequencyNode.elements.Len() <= 0 { // this must never happen, only can happen if capacity is set to 0
		panic(fmt.Sprintf("There are no elements in the LFU list for least frequency = %v.", leastFrequencyNode.frequency))
	}

	// get LRU element within LRU frequency
	lruKeyValueFreqElement := leastFrequencyNode.elements.Back()
	lruKeyValueFreq := lruKeyValueFreqElement.Value.(*KeyValueFreq)

	//fmt.Printf("From the least frequency %v, removing the LRU (key %v -> value %v). \n", leastFrequencyNode.frequency, lruKeyValueFreq.key, lruKeyValueFreq.value)

	this.RemoveElementFromFrequencyNode(leastFrequencyElement, lruKeyValueFreqElement)

	delete(this.m, lruKeyValueFreq.key)
	this.totalElements--
}

func (this *LFUCache) RemoveElementFromFrequencyNode(frequencyElement *list.Element, keyValueFreqElement *list.Element) {
	frequencyNode := frequencyElement.Value.(*FrequencyNode) // value of oldFrequencyElement

	// remove from old frequency node
	frequencyNode.elements.Remove(keyValueFreqElement)

	// no values for old frequency -> remove the frequency node from LFU.freqList
	if frequencyNode.elements.Len() <= 0 {
		this.freqList.Remove(frequencyElement)
	}
}

func (this *LFUCache) GetFrequencyNodeName(frequency int) string {
	return fmt.Sprintf("frequency-%v", frequency)
}

func (this *LFUCache) Print() {
	fmt.Println("====================")
	fmt.Printf("Total elements: %v / %v \n", this.totalElements, this.capacity)

	for freqElement := this.freqList.Front(); freqElement != nil; freqElement = freqElement.Next() {
		// for every frequency
		v := freqElement.Value.(*FrequencyNode)

		fmt.Printf("Frequency %v: ", v.frequency)

		for e := v.elements.Front(); e != nil; e = e.Next() {
			keyValueFreq := e.Value.(*KeyValueFreq)

			fmt.Printf("%v -> %v, ", keyValueFreq.key, keyValueFreq.value)
		}

		fmt.Println()
	}

	fmt.Println("====================")
	fmt.Println()
}

func testPut(c *LFUCache, key, value int) {
	fmt.Println()

	c.Put(key, value)
	fmt.Printf("Put value %v to key %v \n", value, key)

	updatedValue := c.GetWithoutFrequencyUpdate(key) // avoid frequency side effect
	fmt.Printf("Updated value for key %v: %v \n", key, updatedValue)

	if updatedValue != value {
		fmt.Printf("FAILURE: expected updated value = %v, actual updated value = %v \n", updatedValue, value)
	}

	c.Print()
}

func testGet(c *LFUCache, key, expectedValue int) {
	fmt.Println()

	value := c.Get(key)
	fmt.Printf("Got value for key %v: %v \n", key, value)

	if value != expectedValue {
		fmt.Printf("FAILURE: expected value = %v, actual value = %v \n", expectedValue, value)
	}

	c.Print()
}

func test1() {
	// Input
	//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]

	// Output
	//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	q := Constructor(2)
	c := &q // pointer, to update the int fields

	testPut(c, 1, 1)  // freq 1 -> 1
	testPut(c, 2, 2)  // freq 1 ->  2, 1
	testGet(c, 1, 1)  // freq 2 -> 1, freq 1 -> 2
	testPut(c, 3, 3)  // freq 2 -> 1, freq 1 -> 3 | (key 2 removed)
	testGet(c, 2, -1) // freq 2 -> 1, freq 1 -> 3
	testGet(c, 3, 3)  // freq 2 -> 3, 1
	testPut(c, 4, 4)  // freq 2 -> 3, freq 1 -> 4 (key 1 removed, even if it had freq > new key)
	testGet(c, 1, -1) // freq 2 -> 3, freq 1 -> 4
	testGet(c, 3, 3)  // freq 3 -> 3, freq 1 -> 4
	testGet(c, 4, 4)  // freq 3 -> 3, freq 2 -> 4
}

func test2() {
	// Input
	// ["LFUCache","put","put","put","put","get"]
	// [[2],[3,1],[2,1],[2,2],[4,4],[2]]

	// Output
	// [null,null,null,null,null,2]

	q := Constructor(2)
	c := &q // pointer, to update the int fields

	testPut(c, 3, 1) // freq 1 -> (3 -> 1)
	testPut(c, 2, 1) // freq 1 -> (2 -> 1), (3 -> 1)
	testPut(c, 2, 2) // freq 2 -> (2 -> 2), freq 1 -> (3 -> 1)
	testPut(c, 4, 4) // freq 2 -> (2 -> 2), freq 1 -> (4 -> 4) (key 3 removed)
	testGet(c, 2, 2) // freq 3 -> (2 -> 2), freq 1 -> (4 -> 4)
}

func main() {
	// 460. LFU Cache
	test1()
	test2()
}

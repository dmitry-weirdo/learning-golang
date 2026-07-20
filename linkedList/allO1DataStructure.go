package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	name      string
	frequency int
	keys      map[string]int // this should be a map as well, to be O(1). So we map to 1
}

type AllOne struct {
	m map[string]*list.Element // key to List Element
	l *list.List
}

func Constructor() AllOne {
	var head = Node{
		"head",
		-1,
		make(map[string]int, 0),
	}

	var tail = Node{
		"tail",
		-100,
		make(map[string]int, 0),
	}

	var l = list.New()
	l.PushFront(&head)
	l.PushBack(&tail)

	// printList(l)

	return AllOne{
		make(map[string]*list.Element),
		l,
	}
}

func getFrequencyNodeName(frequency int) string {
	return fmt.Sprintf("frequency-%v", frequency)
}

func (this *AllOne) Inc(key string) {
	var v, ok = this.m[key]

	if !ok { // key not present in map -> add new
		frequency := 1

		// we assume there is always a tail
		minFrequencyElement := this.l.Front().Next()
		minFrequencyNode := minFrequencyElement.Value.(*Node)

		if minFrequencyNode.frequency != frequency {
			var node1 = Node{
				name:      getFrequencyNodeName(frequency),
				frequency: frequency,
				keys:      map[string]int{key: 1},
			}

			newMinFrequencyElement := this.l.InsertAfter(&node1, this.l.Front())

			this.m[key] = newMinFrequencyElement

			fmt.Printf("Appended new key \"%v\" a new node \"%v\"\n", key, minFrequencyNode.name)
		} else { // add the key to frequency 1
			minFrequencyNode.keys[key] = 1
			fmt.Printf("Appended new key \"%v\" to existing node \"%v\"\n", key, minFrequencyNode.name)

			this.m[key] = minFrequencyElement
		}
	} else { // key present in map -> reposition it to (frequency + 1) node
		var oldFrequencyElement = v
		oldFrequencyNode := v.Value.(*Node)

		oldFrequency := oldFrequencyNode.frequency
		newFrequency := oldFrequency + 1

		fmt.Printf("Key \"%v\" is already present with frequency %v. Updating it to frequency %v. \n", key, oldFrequency, newFrequency)

		// get next node to the key frequency node
		var nextElement *list.Element = v.Next()
		var nextNode = nextElement.Value.(*Node)

		if nextNode.frequency != newFrequency { // insert a new node with frequency
			var newNodeWithFrequency = Node{
				name:      getFrequencyNodeName(newFrequency),
				frequency: newFrequency,
				keys:      map[string]int{key: 1},
			}

			newFrequencyElement := this.l.InsertAfter(&newNodeWithFrequency, oldFrequencyElement)

			this.m[key] = newFrequencyElement

			fmt.Printf("Appended new key \"%v\" a new node \"%v\"\n", key, newNodeWithFrequency.name)
		} else { // add to the existing frequency node
			nextNode.keys[key] = 1
			fmt.Printf("Appended new key \"%v\" to existing node \"%v\"\n", key, nextNode.name)

			this.m[key] = nextElement
		}

		this.removeKeyFromOldFrequencyNode(key, oldFrequencyNode, oldFrequencyElement)
	}
}

func (this *AllOne) removeKeyFromOldFrequencyNode(key string, oldFrequencyNode *Node, oldFrequencyElement *list.Element) {
	// remove the key from the old node
	delete(oldFrequencyNode.keys, key)
	fmt.Printf("Removed the key \"%v\" from the old frequency node \"%v\" \n", key, oldFrequencyNode.name)

	// if the old node has no keys anymore -> remove it (so that min & max are always near head & tail)
	if len(oldFrequencyNode.keys) == 0 {
		this.l.Remove(oldFrequencyElement)

		fmt.Printf("Removed the old frequency node \"%v\" since it doesn't have any keys anymore. \n", oldFrequencyNode.name)
	}
}

func (this *AllOne) Dec(key string) {
	var v, ok = this.m[key]

	if !ok { // key not present in the map -> do nothing
		fmt.Printf("!!! Key \"%v\" is not present in the map. Nothing to decrease. \n", key)

		return
	} else { // key present in map -> reposition it to (frequency - 1) node
		var oldFrequencyElement = v
		oldFrequencyNode := v.Value.(*Node)

		oldFrequency := oldFrequencyNode.frequency
		newFrequency := oldFrequency - 1

		if newFrequency <= 0 {
			fmt.Printf("Key \"%v\" is already present with frequency %v. Updating it to frequency %v -> remove it from the map. \n", key, oldFrequency, newFrequency)

			delete(this.m, key)
		} else {
			fmt.Printf("Key \"%v\" is already present with frequency %v. Updating it to frequency %v. \n", key, oldFrequency, newFrequency)

			// get previous node to the key frequency node
			var previousElement *list.Element = v.Prev()
			var previousNode = previousElement.Value.(*Node)

			if previousNode.frequency != newFrequency { // insert a new node with frequency
				var newNodeWithFrequency = Node{
					name:      getFrequencyNodeName(newFrequency),
					frequency: newFrequency,
					keys:      map[string]int{key: 1},
				}

				newFrequencyElement := this.l.InsertBefore(&newNodeWithFrequency, oldFrequencyElement)

				this.m[key] = newFrequencyElement

				fmt.Printf("Appended new key \"%v\" a new node \"%v\"\n", key, newNodeWithFrequency.name)
			} else { // add to the existing frequency node
				previousNode.keys[key] = 1
				fmt.Printf("Appended new key \"%v\" to existing node \"%v\"\n", key, previousNode.name)

				this.m[key] = previousElement
			}
		}

		this.removeKeyFromOldFrequencyNode(key, oldFrequencyNode, oldFrequencyElement)
	}
}

func (this *AllOne) GetMaxKey() string {
	maxFrequencyElement := this.l.Back().Prev()
	maxFrequencyNode := maxFrequencyElement.Value.(*Node)

	if maxFrequencyNode.frequency < 0 {
		fmt.Printf("There is no maxFrequencyNode element. Returning empty string. \n")
		return ""
	}

	for k := range maxFrequencyNode.keys {
		fmt.Printf("maxFrequencyNode element has frequency %v. Returning first key \"%v\". \n", maxFrequencyNode.frequency, k)

		return k
	}

	// this must NOT happen, keys must be not empty
	return ""
}

func (this *AllOne) GetMinKey() string {
	minFrequencyElement := this.l.Front().Next()
	minFrequencyNode := minFrequencyElement.Value.(*Node)

	if minFrequencyNode.frequency < 0 {
		fmt.Printf("There is no minFrequencyNode element. Returning empty string. \n")
		return ""
	}

	for k := range minFrequencyNode.keys {
		fmt.Printf("minFrequencyNode element has frequency %v. Returning first key \"%v\". \n", minFrequencyNode.frequency, k)

		return k
	}

	// this must NOT happen, keys must be not empty
	return ""
}

func printAllOne(allOne AllOne) {
	fmt.Println("=====================")

	fmt.Printf("Min Key: %v \n", allOne.GetMinKey())
	fmt.Println()

	fmt.Printf("Max Key: %v \n", allOne.GetMaxKey())
	fmt.Println()

	fmt.Println(allOne.m)

	printMap(allOne.m)
	fmt.Println()
	printList(allOne.l)
	fmt.Println("=====================")
}

func printMap(m map[string]*list.Element) {
	for k, v := range m {
		fmt.Printf("%v: %v \n", k, v.Value.(*Node))
	}
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	// 432. All O`one Data Structure
	allOne := Constructor()
	printAllOne(allOne)

	allOne.Inc("hello")
	printAllOne(allOne)

	allOne.Inc("world")
	printAllOne(allOne)

	allOne.Inc("hello")
	printAllOne(allOne)

	allOne.Inc("world")
	printAllOne(allOne)

	allOne.Inc("leet")
	printAllOne(allOne)

	allOne.Dec("bad-key")
	printAllOne(allOne)

	allOne.Dec("hello")
	printAllOne(allOne)

	allOne.Dec("leet")
	printAllOne(allOne)

	allOne.Dec("hello")
	printAllOne(allOne)
}

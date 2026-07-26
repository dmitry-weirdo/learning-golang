package main

import (
	"fmt"
	"math/rand/v2"
)

type RandomizedSet struct {
	m map[int]int // map of value to index
	a []int       // array with values, guarantees the randomization with equal probability
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		a: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok { // value already present -> do nothing
		return false
	}

	// we add to the end of the array -> average O(1)
	this.a = append(this.a, val)

	this.m[val] = len(this.a) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if deletedElementIndex, ok := this.m[val]; !ok { // value not present -> do nothing
		return false
	} else {
		lastIndex := len(this.a) - 1
		lastElement := this.a[lastIndex]

		// lastElement will be at the index of the deleted element
		this.m[lastElement] = deletedElementIndex
		delete(this.m, val)

		// actually, we don't need to swap, we just want to save the last element to the index of the deleted element
		// swap a[index] && a[lastIndex]
		this.a[deletedElementIndex] = lastElement

		// delete last element from array -> O(1)
		//clear(this.a[lastIndex:]) // clear last element for GC
		this.a = this.a[:lastIndex]

		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.IntN(len(this.a)) // returns random [0; n), exactly what we need for the array size
	return this.a[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func insert(s *RandomizedSet, v int, expectedResult bool) { // !!! this MUST be *Randomized set, else the slice field update will NOT work
	result := s.Insert(v)
	fmt.Printf("Inserted %v into the set. Result: %v. \n", v, result)

	if result != expectedResult {
		errorMessage := fmt.Sprintf("Failure on inserting value %v into the set. Expected result: %v, actual result: %v.", v, expectedResult, result)
		panic(errorMessage)
	}
}

func remove(s *RandomizedSet, v int, expectedResult bool) { // !!! this MUST be *Randomized set, else the slice field update will NOT work
	result := s.Remove(v)
	fmt.Printf("Deleted %v from the set. Result: %v. \n", v, result)

	if result != expectedResult {
		errorMessage := fmt.Sprintf("Failure on deleting value %v into the set. Expected result: %v, actual result: %v.", v, expectedResult, result)
		panic(errorMessage)
	}
}

func random(s *RandomizedSet, count int) {
	for i := 0; i < count; i++ {
		value := s.GetRandom()
		fmt.Printf("Got random value from set: %v. \n", value)
	}
}

func test() {
	s := Constructor()
	set := &s

	insert(set, 1, true)
	insert(set, 1, false)

	insert(set, 2, true)
	insert(set, 2, false)

	insert(set, 3, true)

	random(set, 10)

	remove(set, 3, true)
	remove(set, 3, false)

	random(set, 10)
}

func main() {
	// 380. Insert Delete GetRandom O(1)
	test()
}

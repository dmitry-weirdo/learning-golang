package main

import (
	"fmt"
)

type MyStruct struct {
	m      map[int]int // element to count O(1), size N
	unique map[int]int // just unique elements, map to be O(1), size N
}

func create() MyStruct {
	return MyStruct{
		m:      make(map[int]int),
		unique: make(map[int]int),
	}
}

func (this MyStruct) add(x int) {
	if _, ok := this.m[x]; ok { // element already there
		this.m[x] = this.m[x] + 1

		// remove from unique
		delete(this.unique, x)
	} else { // first addition -> add to uniques
		this.m[x] = 1
		this.unique[x] = 1 // just to hold it in a map
	}
}

func (this MyStruct) get_unique() (result int, ok bool) {
	if len(this.unique) == 0 {
		return -1, false
	}

	for k, _ := range this.unique {
		return k, true
	}

	// this should never happen -> we handled it above
	return -1, false
}

// ok true - means deleted, ok false - element was not there
func (this MyStruct) delete(x int) (result int, ok bool) {
	if _, ok := this.m[x]; ok { // element is present

		newFrequency := this.m[x] - 1
		if newFrequency == 0 { // delete from the structure
			delete(this.m, x)

			delete(this.unique, x)
		} else { // decrease frequency
			this.m[x] = this.m[x] - 1

			if newFrequency == 1 {
				this.unique[x] = 1
			}

		}

		return x, true
	} else { // element is not present
		return -1, false
	}
}

func testGetUnique(m MyStruct, expectedValues []int, expectedOk bool) {
	uq, ok := m.get_unique() // todo: ok should be name error

	if ok != expectedOk {
		errMsg := fmt.Sprintf("Expected ok: %v, actual ok: %v", expectedOk, ok)
		panic(errMsg)
	}

	if ok {
		fmt.Printf("getUnique returned: %v \n", uq)

		vFoundInExpected := false
		for _, v := range expectedValues {
			if v == uq {
				vFoundInExpected = true
			}
		}

		if !vFoundInExpected {
			errMsg := fmt.Sprintf("Returned unique value %v, not found in expected values %v", uq, expectedValues)
			panic(errMsg)
		}

	} else {
		fmt.Printf("getUnique returned nothing. No unique elements in the map \n")
	}
}

func testDelete(m MyStruct, valueToDelete int, expectedOk bool) {
	deleted, ok := m.delete(valueToDelete)

	if ok != expectedOk {
		errMsg := fmt.Sprintf("Expected ok: %v, actual ok: %v", expectedOk, ok)
		panic(errMsg)
	}

	if ok {
		fmt.Printf("delete returned: %v \n", deleted)

		if deleted != valueToDelete {
			errMsg := fmt.Sprintf("Returned deleted value %v, not equal to the deleted value %v", deleted, valueToDelete)

			panic(errMsg)
		}

	} else {
		fmt.Printf("delete returned nothing. Element %v was not present.", valueToDelete)
	}
}

// add(x) - add element `x` to the structure
// get_unique() - return some element that is uniquely represented in
//
//	the structure
//
// delete(x) - remove element `x` from the structure
func main() {
	myStruct := create()

	myStruct.add(10)
	testGetUnique(myStruct, []int{10}, true)

	myStruct.add(10)
	testGetUnique(myStruct, []int{}, false)

	myStruct.add(20)

	testGetUnique(myStruct, []int{20}, true)

	myStruct.add(30)
	testGetUnique(myStruct, []int{20, 30}, true)

	testDelete(myStruct, 20, true)
	testGetUnique(myStruct, []int{30}, true)

	testDelete(myStruct, 30, true)
	//fmt.Printf("map: %v \n", myStruct.m)
	//fmt.Printf("uniques: %v \n", myStruct.unique)

	// no uniques expected
	testGetUnique(myStruct, []int{}, false)

	// 10 -> 2, so we have to remove 2 times
	testDelete(myStruct, 10, true)
	testGetUnique(myStruct, []int{10}, true)

	testDelete(myStruct, 10, true)
	testGetUnique(myStruct, []int{}, false)

	myStruct.add(10)
	testGetUnique(myStruct, []int{10}, true)
}

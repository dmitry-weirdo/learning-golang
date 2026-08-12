package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func sliceFunctions() {
	fmt.Println()
	fmt.Println("==================== slices/arrays functions ====================")

	// remove a central element from the slice
	a := []int{1, 2, 3}

	fmt.Printf("Slice before removal: %v \n", a)

	indexToRemove := 1
	removedElement := a[indexToRemove]
	a = append(a[:indexToRemove], a[indexToRemove+1:]...)

	fmt.Printf("Element removed from a[%v] = %v \n", indexToRemove, removedElement)
	fmt.Printf("Slice after removing a[%v]: %v \n", indexToRemove, a)

	// todo: remove last element from the slice (actually the "middle removal" code should also work as general
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func listFunctions() {
	fmt.Println()
	fmt.Println("==================== list.List functions ====================")

	var l = list.New() // list is NOT type, we can push any values

	// push to back
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// push to front
	l.PushFront(0)
	l.PushFront(-1)

	fmt.Printf("list.List: %v \n", l)

	// remove from front
	fromFront := l.Remove(l.Front()).(int) // cast to your type
	fromBack := l.Remove(l.Back()).(int)   // cast to your type

	fmt.Printf("Element removed from list front: %v \n", fromFront)
	fmt.Printf("Element removed from list back: %v \n", fromBack)

	// iterate list forwards
	fmt.Println("Iterating list forwards: ")

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		fmt.Printf("%v ", v)
	}

	fmt.Println()

	// iterate list backwards
	fmt.Println("Iterating list backwards: ")

	for e := l.Back(); e != nil; e = e.Prev() {
		v := e.Value.(int)
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}

func stringBuilder() {
	fmt.Println()
	fmt.Println("==================== StringBuilder functions ====================")

	var r rune = 'a'
	var ch byte = 'b' // 1-byte character is of type byte

	var intValue = 666

	var sb = strings.Builder{}

	sb.WriteRune(r)
	sb.WriteByte(ch)                       // one-byte character
	sb.WriteString(strconv.Itoa(intValue)) // there is no WriteInt

	s := sb.String() // toString

	fmt.Printf("s from StringBuilder: %v \n", s)
}

func main() {
	stringBuilder()
	listFunctions()
	sliceFunctions()
}

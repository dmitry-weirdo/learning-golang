package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	// Keys in the map are unsorted
	var m map[string]int
	fmt.Println(m) // map[] (nil)
	fmt.Println(len(m))

	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)
	fmt.Println(len(m))

	// add value to map
	m["baz"] = 666
	m["foo"] = 777

	fmt.Println()
	fmt.Println("After addition: ")
	fmt.Println(m)
	fmt.Println(len(m))

	fmt.Println()
	fmt.Println("Values from map: ")
	fmt.Println(m["foo"])     // 777
	fmt.Println(m["bad key"]) // 0, doesn't fail, always returns a result

	const goodKey = "foo"
	goodKeyValue, ok := m[goodKey]
	fmt.Printf("Key \"%v\" Exists in the map: %v. Value: %v \n", goodKey, ok, goodKeyValue)

	const badKey = "badKey"
	badKeyValue, ok := m[badKey]
	fmt.Printf("Key \"%v\" Exists in the map: %v. Value: %v \n", badKey, ok, badKeyValue)

	delete(m, "bad key") // will not fail
	delete(m, "foo")

	fmt.Println()
	fmt.Println("After deletion: ")
	fmt.Println(m)
	fmt.Println(len(m))

	var mapStringToString = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	fmt.Println()
	fmt.Println("Values from map<String, String>: ")
	fmt.Println(mapStringToString["key1"])    // value1
	fmt.Println(mapStringToString["bad key"]) // empty string, doesn't fail

	// using the "maps" package
	// see also https://pkg.go.dev/golang.org/x/exp/maps
	maps.Clear(mapStringToString)
	fmt.Println()
	fmt.Println("After clearing: ")
	fmt.Println(mapStringToString)
	fmt.Println(len(mapStringToString))

	// copy will be by reference
	fmt.Println()
	fmt.Println("Copying the maps by reference: ")

	m1 := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
		//"baz": 3, // duplicate keys will not compile
	}

	// not clone, will be referencing the same map (since it's a reference type)
	m2 := m1

	m1["foo"], m2["bar"] = 42, 666
	fmt.Println(m1)
	fmt.Println(m2)

	//m1 == m2 // there is no "==" operator for map
}

package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	// Keys in the map are unsorted
	var m map[string]int
	fmt.Println(m)        // map[] (nil)
	fmt.Println(m == nil) // true // we can compare to nil
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

	m1["foo"], m1["bar"] = 42, 666
	fmt.Println(m1)
	fmt.Println(m2)

	//m1 == m2 // there is no "==" operator for map
	m1EqualsM2 := maps.Equal(m1, m2)
	fmt.Printf("m1 equals m2: %v \n", m1EqualsM2)

	// we can clone with special function, it will be not by reference
	m3 := maps.Clone(m1)

	m1["foo"], m1["bar"] = 1111, 2222
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	m1EqualsM2 = maps.Equal(m1, m2)
	fmt.Printf("m1 equals m2: %v \n", m1EqualsM2)

	m1EqualsM3 := maps.Equal(m1, m3)
	fmt.Printf("m1 equals m3: %v \n", m1EqualsM3)

	useMaps()
}

func useMaps() {
	fmt.Println()
	fmt.Println("======= useMaps ======= ")

	var m map[string][]string // string -> string[]
	fmt.Println(m)
	fmt.Println(m == nil) // true // we can compare to nil

	// multi-line initialization
	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"}, // trailing comma required
	}

	fmt.Println(m)
	fmt.Println(m == nil)    // true // we can compare to nil
	fmt.Println(m["coffee"]) // true // we can compare to nil

	m["other category"] = []string{"Hot Chocolate"}
	fmt.Println(m)

	// note it's NOT m.delete("key")
	delete(m, "tea")
	fmt.Println(m)
	fmt.Println(m["tea"]) // will be []

	v, ok := m["tea"]
	fmt.Println(v, ok) // [], false

	// m2 update will also update m
	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m2["tea"] = []string{"Hot Tea"}
	fmt.Println(m)
	fmt.Println(m2)
}

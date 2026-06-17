package main

import "fmt"

func main() {

	var s struct { // anonymous struct
		name string
		id   int
	}

	fmt.Println(s)

	s.name = "Arthur"
	s.id = 666
	fmt.Println(s)
	fmt.Println(s.name)
	fmt.Println(s.id)

	s2 := s // will clone, separate instance
	fmt.Println(s, s2, s2 == s)

	s2.name = "Tricia" // will NOT update s
	fmt.Println(s, s2, s2 == s)

	// use custom type
	var ms myStruct
	fmt.Println(ms)

	ms = myStruct{
		id:   666,
		name: "Mike",
		age:  42, // trailing comma required
	}

	fmt.Println(ms)

}

type myStruct struct { // add custom type
	id   int
	name string
	age  int
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []int    // slices of ints, no size specified
	fmt.Println(s) // [] (nil)

	s = []int{1, 2, 3}
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = 99
	fmt.Println(s)

	s2 := append(s, 5, 10, 15)
	fmt.Println(s)
	fmt.Println(s2)

	s[0] = 666
	fmt.Println(s)
	fmt.Println(s2) // is NOT updated with s

	s2 = slices.Delete(s2, 1, 3) // will remove elements with indices [1; 3) from slice
	fmt.Println(s)
	fmt.Println(s2)

	slicesLikeArrays()
}

func slicesLikeArrays() {
	fmt.Println()
	fmt.Println("======= slicesLikeArrays ======= ")

	var s []string
	fmt.Println(s)
	fmt.Println(len(s))

	s = []string{"Coffee", "Espresso", "Capuccino"}
	fmt.Println(s)
	fmt.Println(len(s))

	s[1] = "Chai Tea"
	fmt.Println(s)

	s2 := s
	s2[2] = "Chai Latte" // also updates s, unlike the array
	fmt.Println(s, s2)

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)
	fmt.Println(len(s))

	//slices.Delete(s, 1, 2) // if we don't re-assign to s, the len does not change, although the value is deleted
	s = slices.Delete(s, 1, 2) // if we don't re-assign to s, the len does not change, although the value is deleted
	fmt.Println(s)
	fmt.Println(len(s))
}

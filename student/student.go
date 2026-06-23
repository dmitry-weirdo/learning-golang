package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	ID int

	Name string
}

func main() {

	var students []Student

	for i := 0; i < 2; i++ {
		var s Student
		s.ID = i
		s.Name = "some name " + strconv.Itoa(i)

		fmt.Printf("\n--- student #%d ---\n", i+1)

		fmt.Print("Student ID: ")
		fmt.Scanln(&s.ID) // a pointer should be passed to Scanln

		fmt.Print("Student Name: ")
		fmt.Scanln(&s.Name) // a pointer should be passed to Scanln

		students = append(students, s)
	}

	fmt.Println(students)
}

func checkNumber(n int) {
	if n >= 0 {
		fmt.Println("Positive")

	} else {
		fmt.Println("Negative")
	}

}

func call() {
	myNum := 8
	checkNumber(myNum)
}

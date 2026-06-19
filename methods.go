package main

import "fmt"

func main() {
	var i myInt = 666
	even := i.isEven()

	fmt.Printf("MyInt %v is even: %v \n", i, even)

	u := user{
		id:       666,
		username: "Initial name",
	}

	fmt.Printf("User before change: %v \n", u.String())

	u.UpdateName("New name")
	fmt.Printf("User after change: %v \n", u.String())
}

// method receiver cannot be a built-in type like int
//func (i int) isEven() bool {
//}

type myInt int

func (i myInt) isEven() bool {
	return i%2 == 0 // super-ugly standard formatting, no space around arithmetic operators
}

type user struct {
	id       int
	username string
}

// value receiver
func (u user) String() string {
	return fmt.Sprintf("[User ID: %v, User Name: %v]", u.id, u.username)
}

// pointer receiver
func (u *user) UpdateName(name string) {
	u.username = name
}

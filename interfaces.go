package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type printableUser struct {
	id       int
	username string
}

func (u printableUser) Print() string {
	//TODO implement me
	return fmt.Sprintf("%v [%v] \n", u.id, u.username)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (m menuItem) Print() string {
	var b bytes.Buffer

	b.WriteString(m.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")

	for size, cost := range m.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}

func main() {
	var p printer

	p = printableUser{
		id:       666,
		username: "John Doe",
	}

	fmt.Println(p.Print())

	p = menuItem{
		name: "Coffee",

		prices: map[string]float64{
			"small":  1.65,
			"medium": 1.80,
			"large":  1.95,
		},
	}

	fmt.Println(p.Print())

	// will panic in case of ClassCastException
	//u := p.(printableUser)

	// form with 2 values return won't fail in case of failed assertion
	u, ok := p.(printableUser)
	fmt.Println(u, ok)

	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	// type switch, using (type)
	switch v := p.(type) {
	case printableUser:
		// v will be printableUser
		fmt.Println("Found a user!", v)

	case menuItem:
		// v will be menuItem
		fmt.Println("Found a menu item!", v)

	default:
		fmt.Println("Unknown value type!", v)
	}
}

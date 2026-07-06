package main

import (
	"container/list"
	"fmt"
)

type BrowserHistory struct {
	l       *list.List
	current *list.Element
	debug   bool
}

func Constructor(homepage string) BrowserHistory {
	history := list.New()
	history.PushBack(homepage)

	return BrowserHistory{
		l:       history,
		current: history.Front(),
		debug:   false,
	}
}

func (this *BrowserHistory) Print() {
	for e := this.l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(string))
	}

	fmt.Printf("| Current: %v \n\n", this.current.Value.(string))
}

func (this *BrowserHistory) Visit(url string) {
	// you CANNOT modify list.Element.next
	// remove elements after the current

	if this.debug {
		fmt.Printf("Removing all elements after \"%v\"... \n", this.current.Value)
	}

	// todo: fix here. It is now removing just one element
	for element := this.current.Next(); element != nil; {
		next := element.Next()
		this.l.Remove(element)

		if this.debug {
			fmt.Printf("Removed element \"%v\" from history. \n", element.Value)
		}

		element = next
	}

	newElement := this.l.InsertAfter(url, this.current)
	this.current = newElement

	if this.debug {
		fmt.Printf("Added new element \"%v\" to history. \n", url)
		this.Print()
	}
}

func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps; i++ {
		if this.current.Prev() == nil {
			if this.debug {
				fmt.Printf("No previous element in history. Moved just %v/%v steps and reached the start of the history. \n", i, steps)
			}

			break
		}

		this.current = this.current.Prev()
	}

	current := this.current.Value.(string)

	if this.debug {
		fmt.Printf("Moved %v steps back to \"%v\". \n", steps, current)
		this.Print()
	}

	return current
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps; i++ {
		if this.current.Next() == nil {
			if this.debug {
				fmt.Printf("No next element in history. Moved just %v/%v steps and reached the end of the history. \n", i, steps)
			}

			break
		}

		this.current = this.current.Next()
	}

	current := this.current.Value.(string)

	if this.debug {
		fmt.Printf("Moved %v steps forward to \"%v\". \n", steps, current)
		this.Print()
	}

	return current
}

func test1() {
	browserHistory := Constructor("neetcode.com")
	browserHistory.debug = true

	browserHistory.Visit("google.com")   // You are in "neetcode.com". Visit "google.com"
	browserHistory.Visit("facebook.com") // You are in "google.com". Visit "facebook.com"
	browserHistory.Visit("youtube.com")  // You are in "facebook.com". Visit "youtube.com"
	browserHistory.Back(1)               // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
	browserHistory.Back(1)               // You are in "facebook.com", move back to "google.com" return "google.com"
	browserHistory.Forward(1)            // You are in "google.com", move forward to "facebook.com" return "facebook.com"
	browserHistory.Visit("linkedin.com") // You are in "facebook.com". Visit "linkedin.com"
	browserHistory.Forward(2)            // You are in "linkedin.com", you cannot move forward any steps.
	browserHistory.Back(2)               // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
	browserHistory.Back(7)               // You are in "google.com", you can move back only one step to "neetcode.com". return "neetcode.com"
}

func main() {
	test1()
}

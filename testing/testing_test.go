package testing

import "testing"

// first letter after Test must be capitalized
func TestAdd(t *testing.T) {
	// arrange
	l, r := 1, 2
	expect := 3

	// act
	result := Add(l, r)

	// assert
	if expect != result {
		// Errorf will also call Fail()
		t.Errorf("Failed to add %v and %v. Expected %v, got %v. \n", l, r, expect, result)
	}
}

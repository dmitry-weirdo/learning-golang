package main

import "fmt"

func main() {
	testScores := []float32{ // slice, no fixed size
		87.3,
		105,
		63.5,
		27,
	}

	//c := testScores // this will map to the same slice in memory
	c := clone(testScores)

	// addresses will be different for a clone
	sameMemory := &testScores[0] == &c[0]

	fmt.Println(&testScores[0], &c[0], sameMemory, c)
}

// [V type] is a generic type constraint
func clone[V any](s []V) []V {
	result := make([]V, len(s))

	for i, v := range s {
		result[i] = v
	}

	return result
}

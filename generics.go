package main

import "fmt"

func main() {
	cloneSlice()
	cloneMapExample()
}

func cloneSlice() {
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

// !! no method overloads in Go, you cannot override the same name with different arguments
func cloneMapExample() {
	testScores := map[string]float32{ // slice, no fixed size
		"Harry":    87.3,
		"Hermione": 105,
		"Ronald":   63.5,
		"Neville":  27,
	}

	c := cloneMap(testScores)

	fmt.Println(c)
}

// map keys must be comparable
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))

	for k, v := range m {
		result[k] = v
	}

	return result
}

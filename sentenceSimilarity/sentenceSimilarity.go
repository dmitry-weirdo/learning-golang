package main

import "fmt"

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}

	// map word -> similar words (for O(1) access)
	m := make(map[string]map[string]bool)

	for _, p := range similarPairs {
		a := p[0]
		b := p[1]

		// but a -> b similarity
		if _, ok := m[a]; !ok {
			m[a] = make(map[string]bool)
		}

		m[a][b] = true

		// but b -> a similarity
		if _, ok := m[b]; !ok {
			m[b] = make(map[string]bool)
		}

		m[b][a] = true
	}

	for i := range sentence1 {
		s1 := sentence1[i]
		s2 := sentence2[i]

		if s1 == s2 { // word is similar to itself
			continue
		}

		if v, ok := m[s1]; !ok { // no mappings for s1 overall
			return false
		} else if !v[s2] { // s1 has some mappings, but has no mapping to s2
			return false
		}
	}

	return true
}

func test(s1, s2 []string, pairs [][]string, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Sentence 1: %v \n", s1)
	fmt.Printf("Sentence 2: %v \n", s2)
	fmt.Printf("Similar pairs: %v \n", pairs)

	result := areSentencesSimilar(s1, s2, pairs)

	fmt.Printf("Sentences are similar: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]string{"great", "acting", "skills"},
		[]string{"fine", "drama", "talent"},
		[][]string{
			{"great", "fine"},
			{"drama", "acting"},
			{"skills", "talent"},
		},
		true,
	)
}

func test2() {
	test(
		[]string{"great"},
		[]string{"great"},
		[][]string{},
		true, // word maps to itself
	)
}

func test3() {
	test(
		[]string{"great"},
		[]string{"doubleplus", "good"},
		[][]string{
			{"great", "doubleplus"},
		},
		false, // different lengths of s1 and s2
	)
}

func main() {
	// 734. Sentence Similarity
	test1()
	test2()
	test3()
}

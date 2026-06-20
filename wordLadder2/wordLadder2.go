package main

import (
	"container/list"
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	return nil
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	// todo: implement method
	m := arrayToMap(wordList)
	fmt.Printf("List\n %v converted to map\n %v \n", wordList, m)

	_, ok := m[endWord]
	if !ok {
		fmt.Printf("End word \"%v\" is not in the map %v. Returning 0. \n", endWord, m)

		return 0
	}

	// word to list of nodes of the previous level
	var previousNodes = make(map[string][]string)

	// word to distance from beginWord
	var distances = make(map[string]int)

	length := len(beginWord)

	queue := list.New()
	appendToQueue(queue, beginWord)

	//fmt.Printf("Queue state: %v \n", queue)

	// beginWord counts as 1
	depth := 1

	// beginWord distance counts as 0
	distance := 0
	distances[beginWord] = distance

	var chars = make([]byte, length)

	for queue.Len() > 0 {
		currentLevelElements := queue.Len()
		fmt.Println()
		fmt.Printf("Current depth: %v, queue elements on this level: %v \n", depth, currentLevelElements)
		depth++
		distance++

		for j := 0; j < currentLevelElements; j++ {
			// removing just the current level
			word := removeFirstFromQueue(queue)

			fmt.Printf("Removed from queue: %v, depth: %v \n", word, depth)

			chars = []byte(word)
			for i := 0; i < length; i++ { // for every position in a word
				//fmt.Printf("\n")
				//fmt.Printf("Character position: %v \n", i)

				originalChar := chars[i]

				for c := byte('a'); c <= byte('z'); c++ { // iterate all the 26 letters
					chars[i] = c

					potentialWord := string(chars)
					//fmt.Printf("Potential word: \"%v\" \n", potentialWord)

					// if we already found this word with the same distance
					// it was already removed from the map, but we need to add another predecessor
					if v, distanceExists := distances[potentialWord]; distanceExists {
						if v == distance {
							addToMap(previousNodes, potentialWord, word)

							fmt.Printf("Next word \"%v\" is already present with distance %v. Added a predecessor \"%v\" to it. Previous map: %v \n", potentialWord, v, word, previousNodes)
						}
					}

					if _, ok := m[potentialWord]; ok {
						fmt.Printf("From word \"%v\", potential word found: \"%v\" \n", word, potentialWord)

						addToMap(previousNodes, potentialWord, word)

						fmt.Printf("Added previous word \"%v\" to next word \"%v\". Previous map: %v \n", word, potentialWord, previousNodes)

						/*						if potentialWord == endWord {
													fmt.Printf("End word \"%v\" found as potential word! Returning depth = %v.", potentialWord, depth)

													return depth
												}
						*/
						appendToQueue(queue, potentialWord)

						// todo: think about removing from a separate copy
						// todo: this should be different for word ladder 2

						// word found the first time -> save its distance
						distances[potentialWord] = distance
						fmt.Printf("Setting distance %v to next word \"%v\". Distances: %v \n", distance, potentialWord, distances)

						delete(m, potentialWord)
					}
				}

				// restore to the original word
				chars[i] = originalChar
			}
		}
	}

	return 0
}

func addToMap(m map[string][]string, word string, previousWord string) {
	existingPreviousWords, ok := m[word]

	if !ok { // add new key
		m[word] = []string{previousWord}
	} else {
		m[word] = append(existingPreviousWords, previousWord)
	}
}

func appendToQueue(queue *list.List, s string) {
	queue.PushBack(s)
}

func removeFirstFromQueue(queue *list.List) string {
	return queue.Remove(queue.Front()).(string)
}

func arrayToMap(s []string) map[string]int {
	var m = make(map[string]int)

	for _, v := range s {
		m[v] = 1
	}

	return m
}

func main() {
	var beginWord, endWord string
	var wordList []string

	// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log", "cog"}

	// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
	// Output should be 0 since the endWord is not in wordList
	//beginWord = "hit"
	//endWord = "cog"
	//wordList = []string{"hot", "dot", "dog", "lot", "log"}

	//beginWord = "red"
	//endWord = "tax"
	//wordList = []string{"ted", "tex", "rex", "tax"}

	ladderLength2(beginWord, endWord, wordList)
}

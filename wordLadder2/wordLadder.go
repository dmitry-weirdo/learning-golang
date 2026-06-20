package main

import (
	"container/list"
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// todo: implement method
	m := arrayToMap(wordList)
	fmt.Printf("List\n %v converted to map\n %v \n", wordList, m)

	_, ok := m[endWord]
	if !ok {
		fmt.Printf("End word \"%v\" is not in the map %v. Returning 0. \n", endWord, m)

		return 0
	}

	length := len(beginWord)

	queue := list.New()
	appendToQueue(queue, beginWord)

	//fmt.Printf("Queue state: %v \n", queue)

	// beginWord counts as 1
	depth := 1

	var chars = make([]byte, length)

	for queue.Len() > 0 {
		currentLevelElements := queue.Len()
		fmt.Println()
		fmt.Printf("Current depth: %v, queue elements on this level: %v \n", depth, currentLevelElements)
		depth++

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

					if _, ok := m[potentialWord]; ok {
						fmt.Printf("From word \"%v\", potential word found: \"%v\" \n", word, potentialWord)

						if potentialWord == endWord {
							fmt.Printf("End word \"%v\" found as potential word! Returning depth = %v.", potentialWord, depth)

							return depth
						}

						appendToQueue(queue, potentialWord)

						// todo: think about removing from a separate copy
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

	ladderLength(beginWord, endWord, wordList)
}

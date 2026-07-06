package main

import (
	"container/list"
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	m := arrayToMap(wordList)
	fmt.Printf("List\n %v converted to map\n %v \n", wordList, m)

	_, ok := m[endWord]
	if !ok {
		fmt.Printf("End word \"%v\" is not in the map %v. Returning 0. \n", endWord, m)

		return make([][]string, 0)
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

	endWordFound := false

	for (queue.Len() > 0) && !endWordFound { // if end word already reached -> no need to process the further distances
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

					if _, ok := m[potentialWord]; ok { // only proceed if potential char-generated word is in wordList
						fmt.Printf("From word \"%v\", potential word found: \"%v\" \n", word, potentialWord)

						addToMap(previousNodes, potentialWord, word)

						fmt.Printf("Added previous word \"%v\" to next word \"%v\". Previous map: %v \n", word, potentialWord, previousNodes)

						// in Word Ladder 1, we returned the depth immediately on reaching the endWord
						/*						if potentialWord == endWord {
													fmt.Printf("End word \"%v\" found as potential word! Returning depth = %v.", potentialWord, depth)

													return depth
												}
						*/
						appendToQueue(queue, potentialWord)

						// todo: think about removing from a separate copy

						// word found the first time -> save its distance
						distances[potentialWord] = distance
						fmt.Printf("Setting distance %v to next word \"%v\". Distances: %v \n", distance, potentialWord, distances)

						// word already handled -> removed it from the dictionary, we cannot use it in the further path
						delete(m, potentialWord)

						if potentialWord == endWord {
							fmt.Printf("End word \"%v\" found! \n", endWord)
							endWordFound = true
						}
					}
				}

				// restore to the original word
				chars[i] = originalChar
			}
		}
	}

	fmt.Printf("End word \"%v\" found: %v \n", endWord, endWordFound)

	// with DFS, go from the endWord to the beginWord via previousNodes, collecting all possible paths
	var result = make([][]string, 0)

	if endWordFound { // should always happen
		fmt.Println()
		fmt.Printf("============================")
		fmt.Printf("Building paths from endWord \"%v\" to beginWord \"%v\"... \n", endWord, beginWord)

		path := list.New() // just the current path in the recursion
		appendToQueue(path, endWord)

		buildPaths(path, &result, previousNodes, beginWord, endWord)

		fmt.Printf("Returning %v results: \n%v \n", len(result), result)
	}

	return result
}

func buildPaths(
	path *list.List,
	result *[][]string,
	previousNodes map[string][]string,
	beginWord string,
	currentWord string,
) {
	if currentWord == beginWord { // reached the beginWord
		fmt.Printf("beginWord \"%v\" reached. \n", beginWord)

		resultPath := queueToArray(path)

		fmt.Printf("Adding path %v to result. \n", resultPath)
		*result = append(*result, resultPath)

		return
	}

	// iterate through all the predecessors of the currentWord
	predecessors := previousNodes[currentWord]
	fmt.Printf("Predecessors of word \"%v\": %v \n", currentWord, predecessors)

	for _, predecessor := range predecessors {
		appendToQueue(path, predecessor)

		buildPaths(path, result, previousNodes, beginWord, predecessor)

		// remove the handled predecessor
		removeLastFromQueue(path)
	}
}

func addToMap(m map[string][]string, word string, previousWord string) {
	existingPreviousWords, ok := m[word]

	if !ok { // add new key
		m[word] = []string{previousWord}
	} else { // add one more previous word of the same distance to the map of the word
		m[word] = append(existingPreviousWords, previousWord)
	}
}

func queueToArray(queue *list.List) []string {
	// todo: probably we need to iterate backwards
	var array = make([]string, 0)

	element := queue.Back()

	for element != nil {
		array = append(array, element.Value.(string))

		element = element.Prev()
	}

	return array
}

func appendToQueue(queue *list.List, s string) {
	queue.PushBack(s)
}

func removeFirstFromQueue(queue *list.List) string {
	return queue.Remove(queue.Front()).(string)
}

func removeLastFromQueue(queue *list.List) string {
	return queue.Remove(queue.Back()).(string)
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
	wordList = []string{"hot", "dot", "dog", "lot", "log", "cog", "cox"}

	// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
	// Output should be 0 since the endWord is not in wordList
	//beginWord = "hit"
	//endWord = "cog"
	//wordList = []string{"hot", "dot", "dog", "lot", "log"}

	//beginWord = "red"
	//endWord = "tax"
	//wordList = []string{"ted", "tex", "rex", "tax"}

	findLadders(beginWord, endWord, wordList)
}

package main

import (
	"cmp"
	"fmt"
	"slices"
)

type VideoFrequency struct {
	name      string // video name
	frequency int    // video watch frequency
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	// watchedVideos[i] is the videos to calculate the frequencies
	// friends[i] is already the adj[i]
	// id is the starting node (level = 0)

	queue := make([]int, 0)
	queue = append(queue, id) // starting node

	visited := make(map[int]bool)
	visited[id] = true

	currentLevel := 0

	for len(queue) > 0 {
		currentLevelLength := len(queue)

		for range currentLevelLength {
			// poll node from queue
			current := queue[0]
			queue = queue[1:]

			neighbors := friends[current]

			for _, v := range neighbors {
				if visited[v] { // don't go back in the graph
					continue
				}

				queue = append(queue, v)
				visited[v] = true
			}
		}

		currentLevel++

		if level == currentLevel { // reached the current level
			break
		}
	}

	//fmt.Printf("Nodes of level %v: %v \n", level, queue)

	// for the target level, collect video to frequency
	m := make(map[string]int)

	for _, user := range queue {
		for _, video := range watchedVideos[user] {
			m[video]++
		}
	}

	// convert map to VideoFrequency array
	videos := make([]VideoFrequency, len(m))

	i := 0

	for videoName, videoFrequency := range m {
		videos[i] = VideoFrequency{videoName, videoFrequency}
		i++
	}

	// sort videos by frequency, name
	slices.SortFunc(videos, func(a, b VideoFrequency) int {
		if a.frequency == b.frequency { // frequencies are the same -> compare by name
			return cmp.Compare(a.name, b.name)
		}

		// frequencies are different -> sort by frequency
		return cmp.Compare(a.frequency, b.frequency)
	})

	// collect just video names to result
	result := make([]string, len(videos))

	i = 0

	for _, video := range videos {
		result[i] = video.name
		i++
	}

	return result
}

func test(watchedVideos [][]string, friends [][]int, id int, level int, expectedResult []string) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Watched videos by user: %v \n", watchedVideos)
	fmt.Printf("Friends by user: %v \n", friends)
	fmt.Printf("Start user id (0-based): %v \n", id)
	fmt.Printf("Friends level to search videos: %v \n", level)

	result := watchedVideosByFriends(watchedVideos, friends, id, level)

	fmt.Printf("Watched videos of level %v ordered by frequency asc: %v \n", level, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	videos := [][]string{
		{"A", "B"},
		{"C"},
		{"B", "C"},
		{"D"},
	}

	friends := [][]int{
		{1, 2},
		{0, 3},
		{0, 3},
		{1, 2},
	}

	id := 0
	level := 1

	expected := []string{"B", "C"}

	test(videos, friends, id, level, expected)
}

func test2() {
	videos := [][]string{
		{"A", "B"},
		{"C"},
		{"B", "C"},
		{"D"},
	}

	friends := [][]int{
		{1, 2},
		{0, 3},
		{0, 3},
		{1, 2},
	}

	id := 0
	level := 2

	expected := []string{"D"}

	test(videos, friends, id, level, expected)
}

func main() {
	// 1311. Get Watched Videos by Your Friends
	test1()
	test2()
}

package main

import (
	"fmt"
	"slices"
	"strings"
)

func findItinerary(tickets [][]string) []string {
	// sort tickets by src, destination
	// so that when we collect the adjacency list, for every node we'll add the destinations in the order
	slices.SortFunc(tickets, func(a, b []string) int {
		if a[0] == b[0] { // same from -> compare by to
			return strings.Compare(a[1], b[1])
		}

		// compare by from
		return strings.Compare(a[0], b[0])
	})

	fmt.Printf("Tickets sorted by from, to: %v \n", tickets)

	// map string -> list<string> sorted alphabetically
	adj := make(map[string][]string)

	// create adjacency list
	for _, v := range tickets {
		from := v[0]
		to := v[1]

		if _, ok := adj[from]; !ok {
			adj[from] = []string{}
		}

		adj[from] = append(adj[from], to)
	}

	fmt.Printf("Adjacency list, destinations sorted within every node: \n%v \n", adj)

	// start node is "JFK"
	return nil
}

func test(tickets [][]string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tickets: \n%v \n", tickets)

	result := findItinerary(tickets)

	fmt.Printf("Reconstructed itinerary: %v \n", result)
}

func test1() {
	tickets := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}

	test(tickets)
}

func test2() {
	tickets := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}

	test(tickets)
}

func main() {
	// 332. Reconstruct Itinerary
	test1()
	test2()
}

package main

import (
	"container/list"
	"fmt"
	"slices"
	"strings"
)

func findItinerary(tickets [][]string) []string {
	// see https://www.youtube.com/watch?v=8MpoO2zA2l4
	// for Hierholzer's algorithm description

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
	fmt.Printf("Total nodes: %v \n", len(adj))

	// fill the outDegree map -> we need to map from string, not just by index
	out := make(map[string]int)

	for k, v := range adj {
		out[k] = len(v)
	}

	fmt.Printf("Out map: \n%v \n", out)

	// we're skipping the checks of in/out nodes, since the condition says that the Eulerian path will exist in the graph

	// start node is "JFK" -> we don't need to search for the starting node

	// Eulerian path
	path := list.New() // todo: we can use slice as well

	dfs("JFK", adj, out, path)

	result := []string{}

	for e := path.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(string))
	}

	return result
}

func dfs(
	node string,
	adj map[string][]string,
	out map[string]int,
	path *list.List,
) {
	outgoingEdges := adj[node]

	// while the current node has outgoing edges
	for out[node] > 0 {
		// imagine neighbors = a, b, c
		// out starts with 3 -> we need to take 0 in this case
		nextNodeIndex := len(outgoingEdges) - out[node]
		nextNode := outgoingEdges[nextNodeIndex]

		// decrease the node
		out[node]--

		dfs(nextNode, adj, out, path)
	}

	// add current node post-order BEFORE the dfs-later neighbours
	path.PushFront(node) // todo: should we push back?

	fmt.Printf("Added node %v to the front of the list after going through all of its outgoing edges. \n", node)
	fmt.Println("Current path:")
	printList(path)
	//path.PushBack()
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Println()
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

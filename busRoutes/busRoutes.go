package main

import (
	"container/list"
	"fmt"
)

type RouteStop struct {
	stop     int
	busCount int
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	// BFS on stops, but we track visitedRoutes and visitedStops

	// store just stop, distance used is the BFS level
	return numBusesToDestination_new(routes, source, target)

	// store RouteStop pair of stop + dist
	//return numBusesToDestination_old(routes, source, target)
}

func numBusesToDestination_new(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stopsToRoutes := fillMapsToRoutes(routes)

	if _, ok := stopsToRoutes[source]; !ok {
		return -1
	}

	if _, ok := stopsToRoutes[target]; !ok {
		return -1
	}

	visitedRoutes := make(map[int]bool) // map to have O(1) access, there is no Set in Go
	visitedStops := make(map[int]bool)  // map to have O(1) access, there is no Set in Go

	queue := list.New()
	queue.PushBack(source) // just push stops

	distance := 0

	for queue.Len() > 0 {
		currentLevelElements := queue.Len()

		for range currentLevelElements {
			stop := queue.Remove(queue.Front()).(int)

			if stop == target {
				return distance
			}

			stopRoutes := stopsToRoutes[stop]

			for _, route := range stopRoutes {
				if _, ok := visitedRoutes[route]; ok { // route already visited -> skip
					continue
				}

				visitedRoutes[route] = true

				routeStops := routes[route]

				for _, routeStop := range routeStops {
					if _, ok := visitedStops[routeStop]; ok { // stop already visited -> skip
						continue
					}

					if routeStop == target {
						return distance + 1
					}

					visitedStops[routeStop] = true

					queue.PushBack(routeStop)
				}
			}
		}

		distance++
	}

	return -1
}

func numBusesToDestination_old(routes [][]int, source int, target int) int {
	// bus number = bus index = route number = index in routes[]
	// map every stop to list of routes

	if source == target {
		fmt.Printf("Source stop %v is the same as target stop %v. Returning 0. \n", source)

		return 0
	}

	// fill a map stop -> routes[]
	stopsToRoutes := fillMapsToRoutes(routes)

	fmt.Printf("Filled map stop -> routes: \n%v \n", stopsToRoutes)

	if _, ok := stopsToRoutes[source]; !ok {
		fmt.Printf("Source stop %v does not belong to any route. Returning -1. \n", source)
		return -1
	}

	if _, ok := stopsToRoutes[target]; !ok {
		fmt.Printf("Target stop %v does not belong to any route. Returning -1. \n", target)
		return -1
	}

	queue := list.New()
	visitedRoutes := make(map[int]bool) // map to have O(1) access, there is no Set in Go
	visitedStops := make(map[int]bool)  // map to have O(1) access, there is no Set in Go

	// start with the source with bus distance 0
	sourceStop := RouteStop{source, 0}

	queue.PushBack(sourceStop)
	visitedStops[source] = true

	for queue.Len() > 0 {
		fmt.Println()
		fmt.Printf("Queue: \n")
		printQueue(queue)

		fmt.Printf("Visited routes: %v \n", visitedRoutes)
		fmt.Printf("Visited stops: %v \n", visitedStops)

		currentStop := queue.Remove(queue.Front()).(RouteStop)
		stop := currentStop.stop

		fmt.Printf("Current stop: %v, busCount: %v \n", stop, currentStop.busCount)

		if currentStop.stop == target {
			fmt.Printf("Reached the target stop %v. Returning busCount = %v. \n", target, currentStop.busCount)
			return currentStop.busCount
		}

		stopRoutes := stopsToRoutes[stop]
		fmt.Printf("Stop %v routes: %v \n", stop, stopRoutes)

		// handle all routes including the stop
		for _, route := range stopRoutes {
			_, routePresentInVisitedRoute := visitedRoutes[route]
			if routePresentInVisitedRoute { // route already visited -> skip it
				fmt.Printf("Route %v from stop %v is already visited. Skipping this route. \n", route, stop)
				continue
			}

			// add route to visited routes
			visitedRoutes[route] = true

			routeStops := routes[route]
			fmt.Printf("Route %v stops: %v \n", route, routeStops)

			for _, routeStop := range routeStops {
				_, stopPresentInVisitedStops := visitedStops[routeStop]
				if stopPresentInVisitedStops {
					fmt.Printf("Stop %v from route %v is already visited. Skipping this stop. \n", routeStop, route)
					continue
				}

				// if we're about to add the target stop -> we already found it,
				// no need to do the further operations until we reach it via the queue
				if routeStop == target {
					fmt.Printf("Stop %v from route %v is the target stop. Returning %v. \n", routeStop, route, currentStop.busCount+1)
					return currentStop.busCount + 1
				}

				// add stop to visited stops
				visitedStops[routeStop] = true

				// add stop to the queue with (currentDistance + 1)
				// Since there can be stops of different distance in the map, we save the level (busCount)
				// in every node in the queue.
				// We could also just use a map (stop -> distance) and check it earlier :think:
				stopForQueue := RouteStop{routeStop, currentStop.busCount + 1}
				queue.PushBack(stopForQueue)
			}
		}
	}

	// never reached the target stop -> no solution
	return -1
}

func fillMapsToRoutes(routes [][]int) map[int][]int {
	stopsToRoutes := make(map[int][]int)

	for busNumber, routeStops := range routes {
		for _, routeStop := range routeStops {
			if stopRoutes, ok := stopsToRoutes[routeStop]; !ok {
				// no route yet in this stop -> set an array of just the current route
				stopsToRoutes[routeStop] = []int{busNumber}
			} else {
				// there is already a routes array for this stop -> append current route to it
				stopsToRoutes[routeStop] = append(stopRoutes, busNumber)
			}
		}
	}

	return stopsToRoutes
}

func printQueue(queue *list.List) {
	for v := queue.Front(); v != nil; v = v.Next() {
		fmt.Printf("%v -> %v, ", v.Value.(RouteStop).stop, v.Value.(RouteStop).busCount)
	}

	fmt.Println()
}

func test(routes [][]int, source int, target int, expectedResult int) {
	fmt.Println()
	fmt.Println("=========================")

	fmt.Printf("routes: %v \n", routes)
	fmt.Printf("source = %v, target = %v \n", source, target)

	result := numBusesToDestination(routes, source, target)

	fmt.Printf("result (count of buses from %v to %v): %v \n", source, target, result)
	fmt.Printf("expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	routes := [][]int{
		{1, 2, 7},
		{3, 6, 7},
	}

	source := 1
	target := 6
	expectedResult := 2

	test(routes, source, target, expectedResult)
}

func test2() {
	routes := [][]int{
		{1, 2, 3},
		{2, 4},
		{3, 5},
	}

	source := 1
	target := 5
	expectedResult := 2

	test(routes, source, target, expectedResult)
}

func test3() {
	routes := [][]int{
		{1, 2, 3},
		{3, 4, 5},
		{5, 6},
	}

	source := 1
	target := 6
	expectedResult := 3

	test(routes, source, target, expectedResult)
}

func main() {
	// 815. Bus Routes
	test1()
	test2()
	test3()
}

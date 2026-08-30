package main

import "fmt"

type RouteAggregate struct { // route is aggregation of multiple trips on this route
	// todo: we can add routeKey as well
	totalTime  int
	tripsCount int
}

type TripStart struct {
	startStation string
	startTime    int
}

type UndergroundSystem struct {
	routes map[string]*RouteAggregate // routeKey -> routeInfo
	trips  map[int]*TripStart         // cardId -> started trip. We're guaranteed that the customer (by cardId) is checked in just 1 trip at a time, therefore we're mapping to a single trip. After the trip is ended, we're removing from this map.
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		routes: make(map[string]*RouteAggregate),
		trips:  make(map[int]*TripStart),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if _, ok := this.trips[id]; ok {
		panic(fmt.Sprintf("Customer with card id = %v has already started a non-finished trip.", id))
	}

	this.trips[id] = &TripStart{
		startStation: stationName,
		startTime:    t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, ok := this.trips[id]; !ok {
		panic(fmt.Sprintf("Customer with card id = %v does not have a started and non-finished trip.", id))
	}

	tripStart := this.trips[id]

	routeKey := this.GetRouteKey(tripStart.startStation, stationName)
	tripTime := t - tripStart.startTime

	if _, ok := this.routes[routeKey]; !ok {
		// no data for this route yet -> initialize RouteAggregate for this route
		this.routes[routeKey] = &RouteAggregate{
			totalTime:  0,
			tripsCount: 0,
		}
	}

	// update route data for this route
	this.routes[routeKey].totalTime += tripTime
	this.routes[routeKey].tripsCount++

	// remove the non-finished trip for this cardId
	delete(this.trips, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	routeKey := this.GetRouteKey(startStation, endStation)

	if _, ok := this.routes[routeKey]; !ok {
		// there must be at least 1 trip on this route when getAverageTime is called for this route
		panic(fmt.Sprintf("There is no route data for startStation = \"%v\", endStation = \"%v\".", startStation, endStation))
	}

	routeAggregate := this.routes[routeKey]

	return float64(routeAggregate.totalTime) / float64(routeAggregate.tripsCount)
}

func (this *UndergroundSystem) GetRouteKey(startStation string, endStation string) string {
	return startStation + "_" + endStation
}

func main() {
	// 1396. Design Underground System
}

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

func testAverage(us *UndergroundSystem, startStation, endStation string, expectedResult float64) {
	result := us.GetAverageTime(startStation, endStation)

	fmt.Printf("Average time from \"%v\" to \"%v\": %v \n", startStation, endStation, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	fmt.Println()
	fmt.Println("========================")

	// ["UndergroundSystem","checkIn","checkIn","checkIn","checkOut","checkOut","checkOut","getAverageTime","getAverageTime","checkIn","getAverageTime","checkOut","getAverageTime"]
	// 	[[],[45,"Leyton",3],[32,"Paradise",8],[27,"Leyton",10],
	//	[45,"Waterloo",15],[27,"Waterloo",20],[32,"Cambridge",22],
	//	["Paradise","Cambridge"],["Leyton","Waterloo"],
	//	[10,"Leyton",24],
	//	["Leyton","Waterloo"],
	//	[10,"Waterloo",38],
	//	["Leyton","Waterloo"]]

	// [null,null,null,null,null,null,null,14.00000,11.00000,null,11.00000,null,12.00000]

	us := Constructor()
	us.CheckIn(45, "Leyton", 3)
	us.CheckIn(32, "Paradise", 8)
	us.CheckIn(27, "Leyton", 10)

	us.CheckOut(45, "Waterloo", 15)
	us.CheckOut(27, "Waterloo", 20)
	us.CheckOut(32, "Cambridge", 22)

	testAverage(&us, "Paradise", "Cambridge", float64(14))
	testAverage(&us, "Leyton", "Waterloo", float64(11))

	us.CheckIn(10, "Leyton", 24)
	testAverage(&us, "Leyton", "Waterloo", float64(11)) // started trip does not affect the average yet

	us.CheckOut(10, "Waterloo", 38)
	testAverage(&us, "Leyton", "Waterloo", float64(12))
}

func test2() {
	fmt.Println()
	fmt.Println("========================")

	// ["UndergroundSystem","checkIn","checkOut","getAverageTime","checkIn","checkOut","getAverageTime","checkIn","checkOut","getAverageTime"]
	// [[],[10,"Leyton",3],[10,"Paradise",8],["Leyton","Paradise"],
	// [5,"Leyton",10],[5,"Paradise",16],["Leyton","Paradise"],
	// [2,"Leyton",21],[2,"Paradise",30],["Leyton","Paradise"]]

	// [null,null,null,5.00000,null,null,5.50000,null,null,6.66667]

	us := Constructor()

	us.CheckIn(10, "Leyton", 3)
	us.CheckOut(10, "Paradise", 8)
	testAverage(&us, "Leyton", "Paradise", float64(5)) // 5 / 1 = 5

	us.CheckIn(5, "Leyton", 10)
	us.CheckOut(5, "Paradise", 16)
	testAverage(&us, "Leyton", "Paradise", 5.5) // (5 + 6) / 2 =  11 / 2 = 5.5

	us.CheckIn(2, "Leyton", 21)
	us.CheckOut(2, "Paradise", 30)
	testAverage(&us, "Leyton", "Paradise", float64(20)/float64(3)) // (5 + 6 + 9) / 3 =  20 / 3 = 6.666..
}

func main() {
	// 1396. Design Underground System
	test1()
	test2()
}

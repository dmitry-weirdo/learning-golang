package main

const BIG = 1
const MEDIUM = 2
const SMALL = 3

type ParkingSystem struct {
	// m map[int]int // available places.
	m []int // hashMap is slow -> let's use an array
}

func Constructor(big int, medium int, small int) ParkingSystem {
	m := make([]int, 4)
	m[BIG] = big
	m[MEDIUM] = medium
	m[SMALL] = small

	return ParkingSystem{m}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.m[carType] <= 0 {
		return false
	}

	this.m[carType]--
	return true
}

func main() {
	// 1603. Design Parking System
	// todo: implement tests
}

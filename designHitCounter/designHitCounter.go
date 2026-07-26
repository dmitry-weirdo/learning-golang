package main

import (
	"container/list"
	"fmt"
)

// option 1 - queue with every hit, not grouped by count
// This assumes the queries are only for the last values, since we remove the old values on every getHits() call
/*
type HitCounter struct {
	queue *list.List
}

func Constructor() HitCounter {
	return HitCounter{
		queue: list.New(),
	}
}

func (this *HitCounter) Hit(timestamp int) {
	// we assume that we're pushing the increasing timestamp all the times
	this.queue.PushBack(timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	// O(N) in worst case, but actually every timestamp is removed just 1 time, so it's amortized O(1)

	// we remove the old values (from the queue front) before timestamp - 299
	for this.queue.Len() > 0 && timestamp-this.queue.Front().Value.(int) > 299 {
		this.queue.Remove(this.queue.Front())
	}

	return this.queue.Len()
}
*/

// option 2 - queue that groups hit counts by timestamp.
// This structure is better for cases of many duplicate hits for the same timestamp -> we will store and delete fewer values.
// !!! We assume that timestamps are always increasing, i.e. duplicate count can ONLY add to the latest hit
// The count of hits is stored as a separate variable.

/*type Hit struct {
	timestamp int
	count     int
}

type HitCounter struct {
	queue *list.List
	count int
}

func Constructor() HitCounter {
	return HitCounter{
		queue: list.New(),
		count: 0,
	}
}

func (this *HitCounter) Hit(timestamp int) {
	// we assume that we're pushing the increasing timestamp all the times
	this.count++

	// !!! We need to store *Hit (pointer), NOT Hit, else the count++ will not work
	if this.queue.Len() > 0 && this.queue.Back().Value.(*Hit).timestamp == timestamp {
		// increase hits for the last timestamp
		lastHit := this.queue.Back().Value.(*Hit)
		lastHit.count++
	} else {
		// push new timestamp with count = 1

		newHit := &Hit{timestamp: timestamp, count: 1}
		this.queue.PushBack(newHit)
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	// O(N) in worst case, but actually every timestamp is removed just 1 time, so it's amortized O(1)

	//fmt.Printf("getHits(%v) Before any removal (count = %v): \n", timestamp, this.count)
	//printQueue(*this.queue)

	// we remove the old values (from the queue front) before timestamp - 299
	for this.queue.Len() > 0 && timestamp-this.queue.Front().Value.(*Hit).timestamp > 299 {
		// subtract the count of the removed element from the total count
		this.count -= this.queue.Front().Value.(*Hit).count

		this.queue.Remove(this.queue.Front())
	}

	//fmt.Printf("getHits(%v) After removal (count = %v): \n", timestamp, this.count)
	//printQueue(*this.queue)

	return this.count
}
*/

// option 3 - binary search
// we're saving all the values (and assume they're adding with the non-decreasing timestamp (i.e. the array will be sorted).
// This is if we need to query retrospectively - we don't delete the values and can search any interval.
/*
type HitCounter struct {
	arr []int
}

func Constructor() HitCounter {
	return HitCounter{
		arr: make([]int, 0),
	}
}

func (this *HitCounter) Hit(timestamp int) {
	// O(1) - we're just appending
	this.arr = append(this.arr, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	targetLeft := timestamp - 299

	// search for leftmost value that is >= timestamp - 299
	left := 0
	right := len(this.arr) // - 1 // insert position -> we can return value after the right border

	for left < right {
		mid := (left + right) / 2

		if this.arr[mid] >= targetLeft {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// not retrospective -> don't search the right border
	if len(this.arr) == 0 || timestamp >= this.arr[len(this.arr)-1] {
		return len(this.arr) - left
	}

	leftIndex := left

	// search for leftmost value that is >= timestamp
	targetRight := timestamp

	left = 0
	right = len(this.arr) // insert position -> we can return value after the right border // todo: in this case we can be sure it will be less, but let's stay consistent

	for left < right {
		mid := (left + right) / 2

		if this.arr[mid] >= targetRight {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - leftIndex
}
*/

// option 4 - store 2 arrays of size 300: timestamps array and hits array
// Index = timestamp % 300
// if timestamps[index] == timestamp, we increase hits[index]++
// if timestamps[index] != timestamp, we set new timestamps[index] and set hits[index] = 1
// This solution is better so that it can handle hits added unordered.
// Although we'll reset the values if they're on more than 300 diff for timestamp % 300

type HitCounter struct {
	timestamps []int
	hits       []int
	size       int
}

func Constructor() HitCounter {
	size := 300

	return HitCounter{
		timestamps: make([]int, size),
		hits:       make([]int, size),
		size:       size,
	}
}

func (this *HitCounter) Hit(timestamp int) {
	// O(1)
	index := timestamp % this.size

	if this.timestamps[index] == timestamp {
		// add count to the existing timestamp
		this.hits[index]++
	} else {
		// reset to a new timestamp
		this.timestamps[index] = timestamp
		this.hits[index] = 1
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	// O(300) = O(1)

	// !!! since some timestamps may be outdated (out of the range), we need to check all the timestamps
	sum := 0

	for i, v := range this.timestamps {
		if v >= timestamp-(this.size-1) {
			sum += this.hits[i]
		}
	}

	return sum
}

// ===================================================================== tests =========================================
func printQueue(queue list.List) {

	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Println()
}

func getHits(c *HitCounter, timestamp int, expectedResult int) {
	result := c.GetHits(timestamp)
	fmt.Printf("Hits for timestamp %v: %v. \n", timestamp, result)

	if result != expectedResult {
		errorMessage := fmt.Sprintf("Failure on getting hits for timestamp %v. Expected result: %v, actual result: %v.", timestamp, expectedResult, result)
		panic(errorMessage)
	}
}

func test1() {
	fmt.Println()
	fmt.Println("=======================")

	hc := Constructor()
	c := &hc

	c.Hit(1)
	c.Hit(1)
	c.Hit(1)
	c.Hit(300)
	getHits(c, 300, 4) // [1:300] = 1, 1, 1, 300

	c.Hit(300)
	getHits(c, 300, 5) // [1:300] = 1, 1, 1, 300, 300

	c.Hit(301)
	c.GetHits(301)
	getHits(c, 301, 3) // [2:301] = 300, 300, 301
}

func test2() {
	fmt.Println()
	fmt.Println("=======================")

	hc := Constructor()
	c := &hc

	c.Hit(2)
	c.Hit(3)
	c.Hit(4)
	getHits(c, 300, 3) // [1:300] = 2, 3, 4
	getHits(c, 301, 3) // [2: 301] = 2, 3, 4
	getHits(c, 302, 2) // [3: 302] = 3, 4
	getHits(c, 303, 1) // [4: 303] = 4
	getHits(c, 304, 0) // [5: 304] =

	c.Hit(501)
	getHits(c, 600, 1) // [301:600] = 501

	// !!! this will only work for binary search option (or any options that allow queries for older timestamps
	// retrospective query - not within the problem, but this is why we're using the binary search option
	//getHits(c, 300, 3) // [1:300] = 2, 3, 4
}

func main() {
	// 362. Design Hit Counter
	test1()
	test2()
}

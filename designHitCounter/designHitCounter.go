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

type Hit struct {
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

func test() {
	hc := Constructor()
	c := &hc

	c.Hit(1)
	c.Hit(1)
	c.Hit(1)
	c.Hit(300)
	getHits(c, 300, 4) // 1, 1, 1, 300

	c.Hit(300)
	getHits(c, 300, 5) // 1, 1, 1, 300, 300

	c.Hit(301)
	c.GetHits(301)
	getHits(c, 301, 3) // 300, 300, 301
}

func main() {
	// 362. Design Hit Counter
	test()
}

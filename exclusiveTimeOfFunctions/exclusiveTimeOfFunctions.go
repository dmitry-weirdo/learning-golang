package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Log struct {
	id    int
	event string // todo: can be an enum of "start" and "end"
	time  int
}

type Event struct {
	id    int // function id
	start int // start time of this function call
}

func exclusiveTime(n int, logs []string) []int {
	totalTimes := make([]int, n) // total execution time of all the N functions

	// just the ids of the executing functions. Only contains the executing functions.
	// If a function ends, this execution is removed from the stack
	stack := list.New()

	for _, s := range logs {
		log := parseLog(s)

		fmt.Println()
		fmt.Printf("Log: %v \n", log)

		fmt.Println("Execution stack:")
		printList(stack)

		if log.event == "start" {
			fmt.Printf("Start event: %v \n", log)

			if stack.Len() > 0 { // there is a previous function executing
				// prevFunction is not ended, its execution is still in the stack
				prevFunction := getStackTop(stack)

				// Add the previous function execution time interval to its total.
				// prev function ends at (log.time - 1) and starts at prev.start
				// So, the length of its execution is (log.time - 1 - prev.start + 1) = log.time - prev.start
				executionTime := log.time - prevFunction.start
				totalTimes[prevFunction.id] += executionTime

				fmt.Printf("Previous executing function %v was interrupted. Added interval [%v, %v] = %v to its execution time. \n", prevFunction.id, prevFunction.start, log.time-1, executionTime)
			}

			// start a new function
			startEvent := Event{log.id, log.time}
			stack.PushFront(&startEvent) // !!! we need to store pointers, since we will be updating the startTime values in the events directly in the stack
		} else if log.event == "end" {
			fmt.Printf("End event: %v \n", log)

			// end the current function

			// current function added -> its current execution is removed from the stack
			currentFunction := removeFromStack(stack)

			if currentFunction.id != log.id {
				// This must never happen on correct data
				errorMsg := fmt.Sprintf("End log event %v, but previous function in the stack has id = %v != %v.", log, currentFunction.id, log.id)
				panic(errorMsg)
			}

			// Current function execution interval = (endTime - startTime + 1),
			// e.g., execution from start 2 to 5 takes all 2, 3, 4, 5 = 4 units of time
			executionTime := log.time - currentFunction.start + 1
			totalTimes[currentFunction.id] += executionTime

			fmt.Printf("Function %d ended. Added interval [%v, %v] = %v to its execution time. \n", currentFunction.id, currentFunction.start, log.time, executionTime)

			// change the start time of the previous executing function (if it exists) to the current end time + 1
			if stack.Len() > 0 {
				// prevFunction is not ended, its execution is still in the stack
				prevFunction := getStackTop(stack)

				// When a function ends on time 5, this minute 5 belongs to the ended function execution.
				// So the previous function starts from time 6 (takes minute 6).
				prevFunction.start = log.time + 1

				fmt.Printf("Previous executing function %v re-executed. Changed its start time to %v. \n", prevFunction.id, log.time)
			}
		}
	}

	return totalTimes
}

func removeFromStack(stack *list.List) *Event {
	return stack.Remove(stack.Front()).(*Event)
}

func getStackTop(stack *list.List) *Event { // only called when stack is not empty
	return stack.Front().Value.(*Event)
}

func parseLog(log string) Log {
	parts := strings.Split(log, ":")
	id, _ := strconv.Atoi(parts[0])
	time, _ := strconv.Atoi(parts[2])

	return Log{id, parts[1], time}
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Println()
}

func test(n int, arr []string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Total functions: %v \n", n)
	fmt.Printf("Logs array: %v \n", arr)

	result := exclusiveTime(n, arr)

	fmt.Printf("Execution time of %v functions: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

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
	n := 2
	arr := []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}
	expectedResult := []int{3, 4}

	test(n, arr, expectedResult)
}

func test2() {
	n := 1
	arr := []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}
	expectedResult := []int{8}

	test(n, arr, expectedResult)
}

func test3() {
	n := 2
	arr := []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}
	expectedResult := []int{7, 1}

	test(n, arr, expectedResult)
}

func test4() {
	n := 2
	arr := []string{"0:start:0", "0:end:1", "1:start:3", "1:end:4"} // In the 2nd minute, no function is executed
	expectedResult := []int{2, 2}

	test(n, arr, expectedResult)
}

func main() {
	// 636. Exclusive Time of Functions
	test1()
	test2()
	test3()
	test4()
}

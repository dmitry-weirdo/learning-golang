package main

import (
	"fmt"
	"slices"
)

type Job struct {
	start  int
	end    int
	profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)

	// build an array of one structure with all the fields instead of using 3 separate arrays
	jobs := convertToJobsArray(startTime, endTime, profit, n)
	fmt.Printf("Converted %v jobs into a job objects array: \n%v \n", n, jobs)

	// sort the array just by endTime
	comparator := func(j1, j2 Job) int {
		return j1.start - j2.start
	}

	slices.SortFunc(jobs, comparator)

	fmt.Printf("Sorted %v jobs array by startTime: \n%v \n", n, jobs)

	// memo array for keeping results of jobs[i]
	memo := make([]int, n)
	// todo: do we need to fill it with -1?

	// starting DFS from the 0-th job
	result := dfs(jobs, memo, 0)

	fmt.Printf("Memo array: %v \n", memo)
	return result
}

func dfs(jobs []Job, memo []int, jobIndex int) int {
	// no more jobs to process -> no profit
	n := len(jobs)
	if jobIndex >= n {
		fmt.Printf("jobIndex = %v exceeds the count of %v jobs. Returning 0. \n", jobIndex, n)

		return 0
	}

	// job[i] already calculated -> return the result
	if memo[jobIndex] != 0 {
		fmt.Printf("memo[%v] already contains the results of jobs[%v]. Returning %v. \n", jobIndex, jobIndex, memo[jobIndex])

		return memo[jobIndex]
	}

	// just handle the next job
	nextJobIndexIfWeSkipJob := jobIndex + 1

	// starting for the next job, get the next job in the array where (newJob.startTime
	currentJobEndTime := jobs[jobIndex].end
	nextJobIndexIfWeTakeJob := getNextJobIndex(jobs, currentJobEndTime, jobIndex+1)

	skipJobResult := dfs(jobs, memo, nextJobIndexIfWeSkipJob)
	takeJobResult := dfs(jobs, memo, nextJobIndexIfWeTakeJob) + jobs[jobIndex].profit

	maxResult := max(skipJobResult, takeJobResult)
	fmt.Printf("jobs[%v]. Skipping the job profit: %v, taking the job profit: %v. Taking the max = %v, \n", jobIndex, skipJobResult, takeJobResult, maxResult)

	// save the found result to the cache
	memo[jobIndex] = maxResult

	return maxResult
}

func getNextJobIndex(jobs []Job, minStartTime int, startIndex int) int { // return N if there are no jobs
	// we cannot just iterate from startIndex to n,
	// so we have to use the binary search for O(log n) complexity

	left := startIndex
	right := len(jobs)

	for left < right {
		mid := (left + right) / 2

		if jobs[mid].start >= minStartTime { // move to the left of mid, even if it's the same value
			right = mid
		} else { // move to the right of mid
			left = mid + 1
		}
	}

	return left
}

func convertToJobsArray(startTime []int, endTime []int, profit []int, n int) []Job {
	a := make([]Job, n)

	for i := 0; i < n; i++ {
		job := Job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}

		a[i] = job
	}

	return a
}

func test() {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}

	maxProfit := jobScheduling(startTime, endTime, profit)
	fmt.Printf("Max profit: %v \n", maxProfit)
}

func main() {
	test()
}

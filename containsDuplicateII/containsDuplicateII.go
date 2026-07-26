package main

func containsNearbyDuplicate(nums []int, k int) bool {
	// window of size k

	left := 0

	m := make(map[int]int) // int to freq

	for right := 0; right < len(nums); right++ {
		rightValue := nums[right]

		if v, ok := m[rightValue]; ok && (v > 0) {
			return true
		}

		m[rightValue]++

		if right >= k {
			leftValue := nums[left]
			m[leftValue]--

			if m[leftValue] <= 0 {
				delete(m, leftValue)
			}

			left++
		}
	}

	return false
}

func main() {
	// 219. Contains Duplicate II
}

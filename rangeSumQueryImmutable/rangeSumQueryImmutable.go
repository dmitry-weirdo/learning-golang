package main

type NumArray struct {
	ps []int
}

func Constructor(nums []int) NumArray {
	ps := make([]int, len(nums)+1)
	ps[0] = 0

	for i, v := range nums {
		ps[i+1] = ps[i] + v
	}

	return NumArray{ps: ps}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.ps[right+1] - this.ps[left]
}

func main() {
	// 303. Range Sum Query - Immutable
}

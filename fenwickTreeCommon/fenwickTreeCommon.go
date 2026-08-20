package fenwickTreeCommon

import "fmt"

type FenwickTree struct {
	t []int
}

func createEmptyFenwickTree(n int) FenwickTree {
	t := make([]int, n+1) // t-tree is 1-indexed
	return FenwickTree{t: t}
}

func createFenwickTree(a []int) FenwickTree {
	n := len(a)

	t := make([]int, n+1) // t-tree is 1-indexed
	copy(t[1:], a)        // t[0] is not used, we copy from pos 1

	for i := 1; i <= n; i++ {
		j := i + lsb(i) // add to the LSB of i

		if j <= n { // no overflow of array size
			t[j] += t[i]
		}
	}

	return FenwickTree{t: t}
}

func (ft *FenwickTree) prefixSum(i int) int { // sums elements with indexes [1; i], 1-based, inclusive
	sum := 0

	for i > 0 {
		sum += ft.t[i]

		// clear (set to 0) the least significant bit of i
		i &= ^lsb(i)
	}

	return sum
}

func (ft *FenwickTree) RangeSum(i, j int) int { // 1-based indexes, returns sum of [i, j] inclusive
	if i > j {
		panic(fmt.Sprintf("Incorrect range. i = %v > j = %v.", i, j))
	}

	return ft.prefixSum(j) - ft.prefixSum(i-1) // sum[i - 1] to include i
}

func (ft *FenwickTree) Set(i, value int) {
	currentValue := ft.RangeSum(i, i) // ft[i] is NOT value of the array, it's a partial sum [i - lsb(i) + 1, i]

	delta := value - currentValue // calc the value to add

	ft.Add(i, delta)
}

func (ft *FenwickTree) Add(i, delta int) {
	for i < len(ft.t) {
		ft.t[i] += delta

		// add 1 to the LSB, i.e. propagate to parents
		i += lsb(i)
	}
}

func lsb(x int) int { // least significant bit
	return x & -x
}

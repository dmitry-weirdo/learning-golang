package prefixSumsCommon

func getPrefixSums(a []int) []int {
	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSums := make([]int, len(a)+1)

	prefixSums[0] = 0

	for i, v := range a {
		prefixSumIndex := i + 1
		prefixSums[prefixSumIndex] = prefixSums[i] + v
	}

	return prefixSums
}

func getTotalSum(a []int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

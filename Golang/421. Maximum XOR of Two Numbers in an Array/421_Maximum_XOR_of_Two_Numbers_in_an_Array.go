func findMaximumXOR(nums []int) int {
	maxXOR, mask := 0, 0
	for i := 31; i >= 0; i-- {
		mask |= 1 << i
		prefixSet := map[int]bool{}
		for _, num := range nums {
			prefixSet[num&mask] = true
		}
		candidate := maxXOR | (1 << i)
		for prefix := range prefixSet {
			if prefixSet[prefix^candidate] {
				maxXOR = candidate
				break
			}
		}
	}
	return maxXOR
}

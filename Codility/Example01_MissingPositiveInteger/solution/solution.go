package solution

// Solution returns the smallest missing positive integer of an array of integers
func Solution(A []int) int {
	intN := make([]int, len(A))
	for _, value := range A {
		if value >= 1 && value <= len(A) {
			intN[value-1]++
		}
	}

	for i, occ := range intN {
		if occ == 0 {
			return i + 1
		}
	}

	return len(A) + 1
}

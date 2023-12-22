package main

// HashSet + add to each set
func containsDuplicate(nums []int) bool {

	numsSet := make(map[int]bool) //O(n)

	for _, n := range nums { // O(n)
		if numsSet[n] {
			return true
		}
		numsSet[n] = true
	}

	// Time - O(n)
	// Space - O(n)
	return false
}

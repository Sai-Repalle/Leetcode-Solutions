package main

import "sort"

// sort + look prev element
func containsDuplicate(nums []int) bool {

	sort.Ints(nums)                  // Time- O(nlogn)
	for i := 1; i < len(nums); i++ { // Time - O(n)
		if nums[i] == nums[i-1] {
			return true
		}
	}

	// Time - O(nlogn) +O(n) --> O(nlogn)
	// Space - O(1)
	return false

}

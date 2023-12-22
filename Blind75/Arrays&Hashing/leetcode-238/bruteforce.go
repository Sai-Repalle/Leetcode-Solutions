package main

// use two loops to add the count except the current idx
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums)) // Space - O(n) -- Ignore as output

	for i := 0; i < len(nums); i++ { // Time - O(n)
		count := 1
		for j := 0; j < len(nums); j++ { // Time - O(n)
			if i != j {
				count *= nums[j]
			}
		}
		res[i] = count
	}

	// Time - O(n) * O(n) --> O(n^2)
	// Space - O(n) --> O(n)
	return res
}

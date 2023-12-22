func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int, 0)

	// Count the frequency of each element
	for _, num := range nums { // Time - O(n) Space - O(n)
		freqMap[num]++
	}

	// Build a heap array
	arr := make([][]int, len(nums)) // Space - O(n)

	// Populate the heap array with elements grouped by frequency
	for key := range freqMap { // space-O(n)
		arr[freqMap[key]-1] = append(arr[freqMap[key]-1], key)
	}

	res := make([]int, 0)

	// Concatenate elements from the highest frequency to the lowest
	for i := len(arr) - 1; i >= 0; i-- { // Time - O(n)
		if len(arr[i]) > 0 {
			res = append(res, arr[i]...)
		}
		if len(res) >= k {
			return res
		}
	}

	// Time -- O(n) + O(n) + O(n) --> O(n)
	// Space - O(n) + O(n) --> O(n)
	return res
}

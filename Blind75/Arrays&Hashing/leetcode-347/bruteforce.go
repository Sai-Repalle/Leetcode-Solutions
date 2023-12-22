
import "sort"

// use hashset + build array + sort the slice and use less function
func topKFrequent(nums []int, k int) []int {

	hashSet := make(map[int]int, 0)

	for _, num := range nums { // - Time - O(n)
		hashSet[num]++
	}
	uniqueNums := make([]int, 0, len(hashSet))

	for num := range hashSet { // - Time -O(n)
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(i, j int) bool { // Time- O(nlogn)
		return hashSet[uniqueNums[i]] > hashSet[uniqueNums[j]]
	})

	// Time complexity  - O(n) + O(n) + O(nlogn) --> O(nlogn)
	// Space - O(n) --> for hashMap and ignoring the output result
	return uniqueNums[:k]
}
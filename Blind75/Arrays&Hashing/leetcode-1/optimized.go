package main

func twoSum(nums []int, target int) []int {
	hashSet := make(map[int]int) // O(n)

	for i, num := range nums { //O(n)
		complement := target - num
		if idx, ok := hashSet[complement]; ok {
			return []int{idx, i}
		}
		hashSet[num] = i
	}
	// Time -> O(n)
	// Space - O(n)
	return nil
}

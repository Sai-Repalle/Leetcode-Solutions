package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 0) //O(1)

	// iterate both arrays and start the second array from the outer idx+1
	for i := 0; i < len(nums); i++ { //O(n)
		for j := i + 1; j < len(nums); j++ { //O(n)
			if nums[i]+nums[j] == target { // compare the target, if found, return
				res = append(res, i, j)
				return res
			}
		}
	}
	//Time - O(n) * O(n) --> O(n^2)
	//Space - O(1)
	return res
}

func productExceptSelf(nums []int) []int {
	// Initialize the result array with the length of the input array
	res := make([]int, len(nums)) // space - O(n) --> Ignore

	// Calculate the left products for each element
	leftVal := 1
	for idx, n := range nums { // Time - O(n)
		// Store the product of all elements to the left of the current element
		res[idx] = leftVal
		// Update leftVal for the next iteration
		leftVal *= n
	}

	// Initialize a variable to keep track of the product to the right of each element
	rightVal := 1

	// Calculate the final result by multiplying the left products with the right products
	for idx := len(res) - 1; idx >= 0; idx-- { // Time -O(n)
		// Multiply the left product with the product to the right of the current element
		res[idx] *= rightVal
		// Update rightVal for the next iteration
		rightVal *= nums[idx]
	}

	// Return the final result array
	// Time -- O(n)
	// Space -O(n)
	return res
}

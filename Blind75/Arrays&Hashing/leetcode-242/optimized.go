func isAnagram(s string, t string) bool {
	// Length of both strings not matched return false
	if len(s) != len(t) {
		return false
	}

	// Build hashset for one string
	hashSet := make(map[rune]int) // Space - O(m)

	for _, char := range s { // Time - O(m)
		if _, ok := hashSet[char]; !ok {
			hashSet[char] = 0
		}
		hashSet[char] += 1
	}

	// reduce the hashset based of another string - return false if not exists
	for _, char := range t { // Time - O(n)
		if _, ok := hashSet[char]; !ok {
			return false
		}
		hashSet[char] -= 1
		if hashSet[char] == 0 {
			delete(hashSet, char)
		}
	}

	// Time - O(m) ---> first string length
	// Space - O(m) --> first string length
	return len(hashSet) == 0

}
func isAnagram(s string, t string) bool {

	// iterate through s + reduce the t string to empty
	for _, char := range s { //O(m)
		// Get the idx for the char in the second string
		idx := stringIndexHelper(char, t)
		if idx == -1 {
			return false
		}
		t = t[:idx] + t[idx+1:]
	}

	// Time --> O(m) * 0(n) --> O(mn)
	// Space --> O(1)
	return len(t) == 0
}

func stringIndexHelper(c rune, t string) int { //O(n)
	for i, char := range t {
		if char == c {
			return i
		}
	}

	return -1
}
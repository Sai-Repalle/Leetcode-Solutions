import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	hashSet := make(map[string]int, 0) // Space - O(n)
	res := make([][]string, 0)         // Space - O(n) --> Ignored as output

	for _, s := range strs { // O(n)
		sorted := sortString(s)
		if idx, ok := hashSet[sorted]; ok {
			res[idx] = append(res[idx], s)
		} else {
			idx := len(res)
			hashSet[sorted] = idx
			res = append(res, []string{s})
		}

	}
	// Time - O(n) * O(klogk) --> O(n*klogk)
	// Space - O(1) -- Ingore output string
	return res
}

func sortString(s string) string { // O(klogk)
	charArray := strings.Split(s, "")
	sort.Strings(charArray)
	return strings.Join(charArray, "")
}

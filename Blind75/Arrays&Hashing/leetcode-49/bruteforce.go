import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	// brute force approach
	res := make([][]string, 0)

	for i := 0; i < len(strs); i++ { // O(n)
		found := false
		for j := 0; j < len(res); j++ { // O(m)
			s1 := sortString(res[j][0])
			s2 := sortString(strs[i])

			if s1 == s2 {
				res[j] = append(res[j], strs[i])
				found = true
				break
			}
		}

		if !found {
			res = append(res, []string{strs[i]})
		}
	}

	// Time - O(n) * O(n), O(klogk) --> O(n^2*klogk)
	// Space - O(1) -- Ingore output string
	return res
}

func sortString(s string) string { // O(klogk)
	charArray := strings.Split(s, "")
	sort.Strings(charArray)
	return strings.Join(charArray, "")
}

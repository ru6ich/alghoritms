package medium

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})
	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if citations[i] >= h {
			return h
		}
	}
	return 0
}

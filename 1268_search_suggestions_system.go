package leetcode

import (
	"slices"
	"strings"
)

func suggestedProductsFirst(products []string, searchWord string) [][]string {
	slices.Sort(products)

	ans := make([][]string, 0)
	prefixStr := ""
	for _, c := range searchWord {
		prefixStr = prefixStr + string(c)

		tmpResults := make([]string, 0)
		for _, product := range products {
			if strings.HasPrefix(product, prefixStr) {
				tmpResults = append(tmpResults, product)
				if len(tmpResults) == 3 {
					break
				}
			} else if len(tmpResults) > 0 {
				break
			}
		}
		ans = append(ans, tmpResults)
	}
	return ans
}

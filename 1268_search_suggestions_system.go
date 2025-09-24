package leetcode

import (
	"slices"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	slices.SortFunc(products, func(a, b string) int {
		aArr := []byte(a)
		bArr := []byte(b)
		aArrLen := len(aArr)
		bArrLen := len(bArr)
		size := aArrLen
		if size > bArrLen {
			size = bArrLen
		}

		for i := 0; i < size; i++ {
			if aArr[i] == bArr[i] {
				continue
			}
			return int(aArr[i]) - int(bArr[i])
		}
		return aArrLen - bArrLen
	})
	//log.Println(products)

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

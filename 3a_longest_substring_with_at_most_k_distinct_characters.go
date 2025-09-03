package leetcode

func longestSubstringWithAtMostKDistinctCharacters(s string, k int) int {
	left, right := 0, 0
	charPresentCount := make([]int, 500)
	charDistinctCount := 0
	ans := 0

	for right < len(s) {
		//
		rightChar := s[right]
		if charPresentCount[rightChar] == 0 {
			charDistinctCount++
		}
		charPresentCount[rightChar]++

		// check charDistinctCount condition
		for charDistinctCount > k {
			leftChar := s[left]
			charPresentCount[leftChar]--
			left++
			if charPresentCount[leftChar] == 0 {
				charDistinctCount--
			}
		}

		// update ans
		if right-left+1 > ans {
			ans = right - left + 1
		}

		right++
	}

	return ans
}

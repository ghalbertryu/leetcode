package slidingwindow

/*Given a string s, find the length of the longest substring without duplicate characters
 */
func lengthOfLongestSubstringFirst(s string) int {
	byteIdxMap := make(map[byte]int)
	strByteArr := []byte(s)

	longestSubStrLength := 0
	subStrHeadIdx := 0
	for idx, bt := range strByteArr {
		if preByteIdx, exists := byteIdxMap[bt]; exists {
			subStrLen := idx - subStrHeadIdx
			//log.Println(string(bt), subStrLen) // debug

			if subStrLen > longestSubStrLength {
				longestSubStrLength = subStrLen
			}

			// move head index to the next of the previous same byte index
			for subStrHeadIdx <= preByteIdx {
				c := strByteArr[subStrHeadIdx]
				delete(byteIdxMap, c)
				subStrHeadIdx++
			}
		}
		byteIdxMap[bt] = idx
	}

	endStrIdx := len(strByteArr) - 1
	subStrLen := endStrIdx - subStrHeadIdx + 1
	if subStrLen > longestSubStrLength {
		longestSubStrLength = subStrLen
	}
	return longestSubStrLength
}

func lengthOfLongestSubstringSecond(s string) int {
	strByteArr := []byte(s)
	charPresentArr := make([]bool, 500)

	ans := 0
	left, right := 0, 0
	for right < len(strByteArr) {
		rightChar := strByteArr[right]
		for charPresentArr[rightChar] {
			leftChar := strByteArr[left]
			charPresentArr[leftChar] = false
			left++
		}

		charPresentArr[rightChar] = true
		subStrLen := right - left + 1
		if subStrLen > ans {
			ans = subStrLen
		}
		right++
	}

	return ans
}

func lengthOfLongestSubstringOptimize(s string) int {
	left, right := 0, 0
	charPresent := make([]bool, 500)
	ans := 0

	for right < len(s) {
		for charPresent[s[right]] {
			charPresent[s[left]] = false
			left++
		}
		charPresent[s[right]] = true
		if right-left+1 > ans {
			ans = right - left + 1
		}
		right++
	}

	return ans
}

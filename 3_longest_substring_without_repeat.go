package leetcode

/*Given a string s, find the length of the longest substring without duplicate characters
 */
func lengthOfLongestSubstring(s string) int {
	byteIdxMap := make(map[byte]int)
	strByteArr := []byte(s)

	longestSubStrLength := 0
	subStrHeadIdx := 0
	for idx, b := range strByteArr {
		if preByteIdx, exists := byteIdxMap[b]; exists {
			subStrLen := idx - subStrHeadIdx
			//log.Println(string(b), subStrLen) // debug

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
		byteIdxMap[b] = idx
	}

	endStrIdx := len(strByteArr) - 1
	subStrLen := endStrIdx - subStrHeadIdx + 1
	if subStrLen > longestSubStrLength {
		longestSubStrLength = subStrLen
	}
	return longestSubStrLength
}

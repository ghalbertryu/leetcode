package leetcode

/*Given a string s, find the length of the longest substring without duplicate characters
 */
func lengthOfLongestSubstring(s string) int {
	startByteSubStrLengthCount := make(map[byte]int)
	strByteArr := []byte(s)

	for _, b := range strByteArr {
		startByteSubStrLengthCount[b] = 1
		//
		//if _, isDuplicatedByte := startByteSubStrLengthCount[b]; isDuplicatedByte {
		//	startByteSubStrLengthCount[b] = 1
		//} else {
		//	startByteSubStrLengthCount[b] = 1
		//}

		for character, subStrLen := range startByteSubStrLengthCount {
			if character != b {
				startByteSubStrLengthCount[character] = subStrLen + 1
			}
		}
	}

	ret := 0
	for _, subStrLen := range startByteSubStrLengthCount {
		if subStrLen > ret {
			ret = subStrLen
		}
	}
	return ret
}

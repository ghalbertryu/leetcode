package leetcode

import "log"

/*Given a string s, find the length of the longest substring without duplicate characters
 */
func lengthOfLongestSubstring(s string) int {
	strByteArr := []byte(s)

	ret := 0
	tmpCounter := 0
	presentedByteMapSet := make(map[byte]struct{})
	for _, b := range strByteArr {
		if _, exist := presentedByteMapSet[b]; !exist {
			log.Println(b)
			tmpCounter++
		} else {
			if tmpCounter > ret {
				ret = tmpCounter
			}

			tmpCounter = 1
			presentedByteMapSet = make(map[byte]struct{}) // reset mapSet
		}

		presentedByteMapSet[b] = struct{}{}
	}

	if tmpCounter > ret {
		ret = tmpCounter
	}
	return ret
}

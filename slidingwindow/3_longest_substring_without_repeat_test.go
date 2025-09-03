package slidingwindow

import (
	"log"
	"testing"
)

func TestLengthOfLongestSubstringFirst(t *testing.T) {
	log.Println(lengthOfLongestSubstringFirst("abcabcbb")) // 3
	log.Println(lengthOfLongestSubstringFirst("a"))        // 1
	log.Println(lengthOfLongestSubstringFirst(" "))        // 1
	log.Println(lengthOfLongestSubstringFirst("au"))       // 2
	log.Println(lengthOfLongestSubstringFirst("dvdf"))     // 3
	log.Println(lengthOfLongestSubstringFirst("pwwkew"))   // 3
}

func TestLengthOfLongestSubstringSecond(t *testing.T) {
	log.Println(lengthOfLongestSubstringSecond("abcabcbb")) // 3
	log.Println(lengthOfLongestSubstringSecond("a"))        // 1
	log.Println(lengthOfLongestSubstringSecond(" "))        // 1
	log.Println(lengthOfLongestSubstringSecond("au"))       // 2
	log.Println(lengthOfLongestSubstringSecond("dvdf"))     // 3
	log.Println(lengthOfLongestSubstringSecond("pwwkew"))   // 3
}

func TestLengthOfLongestSubstringOptimize(t *testing.T) {
	log.Println(lengthOfLongestSubstringOptimize("abcabcbb")) // 3
	log.Println(lengthOfLongestSubstringOptimize("a"))        // 1
	log.Println(lengthOfLongestSubstringOptimize(" "))        // 1
	log.Println(lengthOfLongestSubstringOptimize("au"))       // 2
	log.Println(lengthOfLongestSubstringOptimize("dvdf"))     // 3
	log.Println(lengthOfLongestSubstringOptimize("pwwkew"))   // 3
}

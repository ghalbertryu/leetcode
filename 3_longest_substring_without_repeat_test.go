package leetcode

import (
	"log"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	log.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	log.Println(lengthOfLongestSubstring("a"))        // 1
	log.Println(lengthOfLongestSubstring(" "))        // 1
	log.Println(lengthOfLongestSubstring("au"))       // 2
	log.Println(lengthOfLongestSubstring("dvdf"))     // 3
	log.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}

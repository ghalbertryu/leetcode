package leetcode

import (
	"log"
	"testing"
)

func TestLongestSubstringWithAtMostKDistinctCharacters(t *testing.T) {
	log.Println(longestSubstringWithAtMostKDistinctCharacters("eceba", 2))         // 3 ("ece")
	log.Println(longestSubstringWithAtMostKDistinctCharacters("aa", 1))            // 2 ("aa")
	log.Println(longestSubstringWithAtMostKDistinctCharacters("abcadcacacaca", 3)) // 11
}

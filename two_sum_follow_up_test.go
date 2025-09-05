package leetcode

import (
	"log"
	"testing"
)

func TestTwoSumMultiAns(t *testing.T) {
	log.Println(twoSumMultiAns([]int{-2, 2, 7, 11, 15, -6}, 9))
	log.Println(twoSumMultiAns([]int{0, 1, 2, 3}, 3))
	log.Println(twoSumMultiAns([]int{0, 2, 2, 0}, 2))
	log.Println(twoSumMultiAns([]int{0, 0, 2, 2}, 2))
}

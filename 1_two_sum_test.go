package leetcode

import (
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {
	log.Println(twoSumSecond([]int{2, 7, 11, 15}, 9))
	log.Println(twoSumSecond([]int{3, 2, 4}, 6))
	log.Println(twoSumSecond([]int{3, 3}, 6))
	log.Println(twoSumSecond([]int{-3, 4, 3, 90}, 0))
}

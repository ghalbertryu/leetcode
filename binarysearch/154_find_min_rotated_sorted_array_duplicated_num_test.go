package heetcode

import (
	"log"
	"testing"
)

func TestFindMinDuplicatedNum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 5}
	log.Println(findMinDuplicatedNum(arr))

	arr = []int{2, 2, 2, 3, 4, 5, 1}
	log.Println(findMinDuplicatedNum(arr))

	arr = []int{3, 3, 1}
	log.Println(findMinDuplicatedNum(arr))

	arr = []int{3}
	log.Println(findMinDuplicatedNum(arr))

	arr = []int{1, 2, 0, 0, 0}
	log.Println(findMinDuplicatedNum(arr))

	arr = []int{2, 2, 2, 0, 1}
	log.Println(findMinDuplicatedNum(arr))

	arr = []int{10, 9, 10, 10, 10}
	log.Println(findMinDuplicatedNum(arr))
}

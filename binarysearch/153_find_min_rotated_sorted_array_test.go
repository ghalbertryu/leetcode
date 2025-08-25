package heetcode

import (
	"log"
	"testing"
)

func TestFindMin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	log.Println(findMin(arr))

	arr = []int{3, 4, 5, 1, 2}
	log.Println(findMin(arr))

	arr = []int{3, 4, 2}
	log.Println(findMin(arr))

	arr = []int{3, 4}
	log.Println(findMin(arr))

	arr = []int{3}
	log.Println(findMin(arr))
}

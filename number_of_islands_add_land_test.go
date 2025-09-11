package leetcode

import (
	"log"
	"testing"
)

func TestNumIslandsAddLand(t *testing.T) {
	positions := [][]int{
		{0, 0},
		{0, 1},
		{1, 2},
		{2, 1},
	}
	log.Println(numIslandsAddLand(3, 3, positions))
}

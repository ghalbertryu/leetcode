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
	log.Println(numIslandsAddLand(3, 4, positions))
}

func TestFindParent(t *testing.T) {
	p := []int{
		0, 1, 2, 3, 4, 5, 6,
	}

	log.Println(findParent(p, 1))
	log.Println(findParent(p, 4))
	log.Println(findParent(p, 3))
	log.Println(findParent(p, 0))
}

func TestUnion(t *testing.T) {
	p := []int{
		0, 1, 2, 3, 4, 5, 6,
	}

	union(p, 4, 3)
	log.Println(p)

	union(p, 2, 3)
	log.Println(p)

	union(p, 5, 1)
	log.Println(p)
}

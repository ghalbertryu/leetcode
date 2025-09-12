package graph

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
	log.Println(numIslandsAddLand(3, 4, positions)) //  [1 1 2 3]
}

func TestDisjointSetFindParent(t *testing.T) {
	disjointSet := NewDisjointSet(5)

	disjointSet.Make(1)
	disjointSet.Make(2)
	disjointSet.Make(3)
	log.Println(disjointSet.FindParent(1))
	log.Println(disjointSet.FindParent(2))
	log.Println(disjointSet.FindParent(3))

	disjointSet.Make(4)
	disjointSet.Make(5)

	disjointSet.Union(1, 4)
	log.Println(disjointSet.FindParent(1))
	log.Println(disjointSet.FindParent(4))

	disjointSet.Union(5, 4)
	log.Println(disjointSet.FindParent(5))
	log.Println(disjointSet.FindParent(4))
	log.Println(disjointSet.FindParent(1))
}

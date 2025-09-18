package leetcode

import (
	"log"
	"testing"
)

func TestMinimumObstaclesFirst(t *testing.T) {
	grid1 := [][]int{
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
	}
	log.Println(minimumObstacles(grid1))

	grid2 := [][]int{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
	}
	log.Println(minimumObstacles(grid2))

	grid3 := [][]int{
		{0, 0, 1, 1, 1, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 0, 1, 0, 0, 0, 1, 1, 1, 0},
	}
	log.Println(minimumObstacles(grid3))
}

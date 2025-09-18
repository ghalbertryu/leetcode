package graph

import (
	"log"
	"testing"
)

func TestMinimumObstaclesFirst(t *testing.T) {
	grid0 := [][]int{
		{0, 1},
		{1, 0},
	}
	log.Println(minimumObstaclesOptimize(grid0)) // 1

	grid1 := [][]int{
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
	}
	log.Println(minimumObstaclesOptimize(grid1)) // 2

	grid2 := [][]int{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
	}
	log.Println(minimumObstaclesOptimize(grid2)) // 2

	grid3 := [][]int{
		{0, 0, 1, 1, 1, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 0, 1, 0, 0, 0, 1, 1, 1, 0},
	}
	log.Println(minimumObstaclesOptimize(grid3)) // 5
}

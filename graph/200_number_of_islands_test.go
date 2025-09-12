package graph

import (
	"log"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	log.Println(numIslandsOptimize(grid1)) // 1

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	log.Println(numIslandsOptimize(grid2)) // 3

	grid3 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '1', '0', '1'},
		{'0', '0', '1', '0', '0'},
		{'0', '1', '1', '1', '1'},
	}
	log.Println(numIslandsOptimize(grid3)) // 2

	grid4 := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '0', '1', '0', '0'},
		{'1', '1', '1', '0', '0'},
	}
	log.Println(numIslandsOptimize(grid4)) // 1
}

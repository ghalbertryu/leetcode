package leetcode

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
	log.Println(numIslands(grid1))

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	log.Println(numIslands(grid2))

	grid3 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '1', '0', '1'},
		{'0', '0', '1', '0', '0'},
		{'0', '1', '1', '1', '1'},
	}
	log.Println(numIslands(grid3))
}

package leetcode

const (
	WaterInt = 0
	LandInt  = 999
)

// union
// parent

func numIslandsAddLand(m, n int, positions [][]int) []int {
	// build water grid
	grid := make([][]byte, m)
	for i, _ := range grid {
		grid[i] = make([]byte, n)
	}

	islandNum := 0
	ans := make([]int, len(positions))
	for i, p := range positions {
		// check

		grid[p[0]][p[1]] = LandInt

		islandNum++
	}
	return ans
}

var islandNumMap = map[int][][]int{} // islandNum:landPositions

func markAdjacentMinNumLand(x, y int, grid [][]int) {
	adjacentLandPositions := make([][]int, 0)
	adjacentMinNumLand := 0
	for i := 0; i < 4; i++ {
		tmpX, tmpY := x+dx[i], y+dy[i]
		if tmpX >= 0 && tmpX < m &&
			tmpY >= 0 && tmpY < n &&
			grid[tmpY][tmpX] > 0 {
			adjacentLandPositions = append(adjacentLandPositions, []int{tmpY, tmpX})
			if adjacentMinNumLand == 0 || grid[tmpY][tmpX] < adjacentMinNumLand {
				adjacentMinNumLand = grid[tmpY][tmpX]
			}
		}
	}

	if adjacentMinNumLand > 0 {
		for _, position := range adjacentLandPositions {
			grid[position[0]][position[1]]
		}
	}
}

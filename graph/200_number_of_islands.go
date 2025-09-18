package graph

const (
	Land  = '1'
	Water = '0'
)

func numIslandsOptimize(grid [][]byte) int {
	islandNum := 0
	m = len(grid)
	n = len(grid[0])

	for x, columns := range grid {
		for y, elm := range columns {
			if elm == Land {
				dfs(x, y, grid)
				islandNum++
			}
		}
	}

	return islandNum
}

// Depth-First-Search
func dfs(x, y int, grid [][]byte) {
	grid[x][y] = Water
	for i := 0; i < 4; i++ {
		tmpX, tmpY := x+dx[i], y+dy[i]
		if tmpX >= 0 && tmpX < m &&
			tmpY >= 0 && tmpY < n &&
			grid[tmpX][tmpY] == Land {
			dfs(tmpX, tmpY, grid)
		}
	}
}

var (
	yLen    int
	xLen    int
	xMaxIdx int
	yMaxIdx int
)

func numIslandsFirst(grid [][]byte) int {
	islandNum := 0
	yLen = len(grid)
	xLen = len(grid[0])
	xMaxIdx = xLen - 1
	yMaxIdx = yLen - 1
	islandsNumGrid := make([][]int, yLen)
	for i, _ := range grid {
		islandsNumGrid[i] = make([]int, xLen)
	}

	for y, columns := range grid {
		for x, elm := range columns {
			if elm == Water {
				continue
			}
			if islandsNumGrid[y][x] > 0 {
				continue
			}

			// Land
			// check adjacent
			tmpX := x
			tmpY := y
			// left
			tmpX = x - 1
			tmpY = y
			if islandsNumGrid[y][x] == 0 &&
				isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
				islandsNumGrid[tmpY][tmpX] > 0 {
				islandsNumGrid[y][x] = islandsNumGrid[tmpY][tmpX]
			}
			// up
			tmpX = x
			tmpY = y - 1
			if islandsNumGrid[y][x] == 0 &&
				isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
				islandsNumGrid[tmpY][tmpX] > 0 {
				islandsNumGrid[y][x] = islandsNumGrid[tmpY][tmpX]
			}
			// right
			tmpX = x + 1
			tmpY = y
			if islandsNumGrid[y][x] == 0 &&
				isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
				islandsNumGrid[tmpY][tmpX] > 0 {
				islandsNumGrid[y][x] = islandsNumGrid[tmpY][tmpX]
			}
			// down
			tmpX = x
			tmpY = y + 1
			if islandsNumGrid[y][x] == 0 &&
				isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
				islandsNumGrid[tmpY][tmpX] > 0 {
				islandsNumGrid[y][x] = islandsNumGrid[tmpY][tmpX]
			}

			if islandsNumGrid[y][x] == 0 {
				islandNum++
				islandsNumGrid[y][x] = islandNum
				markAdjacentLandNum(x, y, islandNum, grid, islandsNumGrid)
			}
		}
	}

	//log.Println(islandsNumGrid) // debug
	return islandNum
}

func markAdjacentLandNum(x int, y int, islandNum int, grid [][]byte, islandsNumGrid [][]int) {
	tmpX := x
	tmpY := y
	// left
	tmpX = x - 1
	tmpY = y
	if isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
		islandsNumGrid[tmpY][tmpX] == 0 && grid[tmpY][tmpX] == Land {
		islandsNumGrid[tmpY][tmpX] = islandNum
		markAdjacentLandNum(tmpX, tmpY, islandNum, grid, islandsNumGrid)
	}
	// up
	tmpX = x
	tmpY = y - 1
	if isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
		islandsNumGrid[tmpY][tmpX] == 0 && grid[tmpY][tmpX] == Land {
		islandsNumGrid[tmpY][tmpX] = islandNum
		markAdjacentLandNum(tmpX, tmpY, islandNum, grid, islandsNumGrid)
	}
	// right
	tmpX = x + 1
	tmpY = y
	if isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
		islandsNumGrid[tmpY][tmpX] == 0 && grid[tmpY][tmpX] == Land {
		islandsNumGrid[tmpY][tmpX] = islandNum
		markAdjacentLandNum(tmpX, tmpY, islandNum, grid, islandsNumGrid)
	}
	// down
	tmpX = x
	tmpY = y + 1
	if isValidRange(tmpX, tmpY, xMaxIdx, yMaxIdx) &&
		islandsNumGrid[tmpY][tmpX] == 0 && grid[tmpY][tmpX] == Land {
		islandsNumGrid[tmpY][tmpX] = islandNum
		markAdjacentLandNum(tmpX, tmpY, islandNum, grid, islandsNumGrid)
	}
}

func isValidRange(x int, y int, xMaxIdx int, yMaxIdx int) bool {
	return x >= 0 && y >= 0 && x <= xMaxIdx && y <= yMaxIdx
}

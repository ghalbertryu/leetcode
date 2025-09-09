package leetcode

const (
	Land  = '1'
	Water = '0'
)

func numIslands(grid [][]byte) int {
	islandNum := 0
	islandsNumGrid := make([][]int, len(grid))
	for i, row := range grid {
		islandsNumGrid[i] = make([]int, len(row))
	}

	for y, columns := range grid {
		for x, elm := range columns {
			if elm == Water {
				continue
			}

			// land
			tmpX := x
			tmpY := y

			// check up
			tmpX = x
			tmpY = y - 1
			if tmpX >= 0 && tmpY >= 0 &&
				islandsNumGrid[tmpY][tmpX] > 0 {
				islandsNumGrid[y][x] = islandsNumGrid[tmpY][tmpX]
				continue
			}

			// check left
			tmpX = x - 1
			tmpY = y
			if tmpX >= 0 && tmpY >= 0 &&
				islandsNumGrid[tmpY][tmpX] > 0 {
				islandsNumGrid[y][x] = islandsNumGrid[tmpY][tmpX]
				continue
			}

			islandNum++
			islandsNumGrid[y][x] = islandNum
		}
	}

	//log.Println(islandsNumGrid) // debug
	return islandNum
}

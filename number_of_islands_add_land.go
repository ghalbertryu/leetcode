package leetcode

// union
func union(p []int, a, b int) {
	parentA := findParent(p, a)
	parentB := findParent(p, b)
	if parentA != parentB {
		p[b] = parentA
	}
}

// parent
func findParent(p []int, idx int) int {
	if p[idx] == idx {
		return idx
	}
	return findParent(p, p[idx])
}

func numIslandsAddLand(m, n int, positions [][]int) []int {
	// build water grid
	grid := make([][]int, m)
	for i, _ := range grid {
		grid[i] = make([]int, n)
	}
	landParent := make([]int, len(positions)+1)

	ans := make([]int, 0)
	islandCount := 0
	for i, p := range positions {
		// set island
		landNum := i + 1
		x := p[0]
		y := p[1]
		grid[x][y] = landNum
		landParent[landNum] = landNum

		// union island
		islandCount++
		for i := 0; i < 4; i++ {
			tmpX, tmpY := x+dx[i], y+dy[i]
			if tmpX >= 0 && tmpX < m &&
				tmpY >= 0 && tmpY < n &&
				grid[tmpX][tmpY] > 0 {

				tmpLandNum := grid[tmpX][tmpY]
				if findParent(landParent, landNum) != findParent(landParent, tmpLandNum) {
					union(landParent, tmpLandNum, landNum)
					islandCount--
				}
			}
		}

		ans = append(ans, islandCount)
	}
	return ans
}

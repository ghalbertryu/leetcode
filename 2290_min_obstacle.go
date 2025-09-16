package leetcode

func minimumObstaclesFirst(grid [][]int) int {
	m = len(grid)
	n = len(grid[0])
	endX, endY = m-1, n-1
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
	}

	// path
	minNumOfRemoveObstacles, _ := dfs(grid, 0, 0, path)
	return minNumOfRemoveObstacles
}

var (
	m          int
	n          int
	endX, endY int
	dx         = []int{0, 1, -1, 0}
	dy         = []int{1, 0, 0, -1}
)

// dfs
// return number of remove obstacle
// return true if reach to end X, Y
func dfs(grid [][]int, x, y int, path [][]int) (int, bool) {
	path[x][y] = 1
	removeCount := 0
	if grid[x][y] == 1 {
		removeCount++
	}
	if x == endX && y == endY {
		path[x][y] = 0
		return removeCount, true
	}

	minDfsCount := -1
	for i := 0; i < 4; i++ {
		tmpX, tmpY := x+dx[i], y+dy[i]
		if tmpX >= 0 && tmpX < m &&
			tmpY >= 0 && tmpY < n &&
			path[tmpX][tmpY] == 0 { // 確認沒走過

			dfsCount, success := dfs(grid, tmpX, tmpY, path)
			if !success {
				continue
			}
			if minDfsCount < 0 {
				minDfsCount = dfsCount
			} else if dfsCount < minDfsCount {
				minDfsCount = dfsCount
			}
		}
	}
	//log.Println(x, y, minDfsCount) // debug
	path[x][y] = 0
	return removeCount + minDfsCount, minDfsCount >= 0
}

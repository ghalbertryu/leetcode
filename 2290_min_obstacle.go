package leetcode

var (
	m          int
	n          int
	endX, endY int
	dx         = []int{0, 1, -1, 0}
	dy         = []int{1, 0, 0, -1}
)

func minimumObstaclesSecond(grid [][]int) int {
	m = len(grid)
	n = len(grid[0])
	endX, endY = m-1, n-1

	idGrid := make([][]int, m)
	idXArr := make([]int, m*n+1)
	idYArr := make([]int, m*n+1)

	for i, _ := range idGrid {
		idGrid[i] = make([]int, n)
	}
	id := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idGrid[i][j] = id
			idXArr[id] = i
			idYArr[id] = j
			id++
		}
	}

	// bfs
	visitedIdArr := make([]int, m*n+1)
	disIdArr := make([]int, m*n+1)
	adjacentQueue := make([][]int, 0)
	adjacentQueue = append(adjacentQueue, []int{0, 0})
	for len(adjacentQueue) > 0 {
		pop := adjacentQueue[0]
		adjacentQueue = adjacentQueue[1:]
		popId := pop[0]
		fromId := pop[1]
		if visitedIdArr[popId] == 1 {
			continue
		}

		x := idXArr[popId]
		y := idYArr[popId]
		disIdArr[popId] = disIdArr[fromId] + grid[x][y]
		visitedIdArr[idGrid[x][y]] = 1
		if x == endX && y == endY {
			break
		}

		for i := 0; i < 4; i++ {
			tmpX, tmpY := x+dx[i], y+dy[i]
			if tmpX >= 0 && tmpX < m &&
				tmpY >= 0 && tmpY < n {
				tmpId := idGrid[tmpX][tmpY]
				if grid[tmpX][tmpY] > 0 {
					// append
					adjacentQueue = append(adjacentQueue, []int{tmpId, popId})
				} else {
					// prepend
					adjacentQueue = append([][]int{{tmpId, popId}}, adjacentQueue...)
				}
			}
		}
	}

	//log.Println(disIdArr) // debug
	endId := idGrid[endX][endY]
	return disIdArr[endId]
}

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

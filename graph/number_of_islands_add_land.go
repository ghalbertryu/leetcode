package graph

func numIslandsAddLand(m, n int, positions [][]int) []int {
	// build water grid
	grid := make([][]int, m)
	for i, _ := range grid {
		grid[i] = make([]int, n)
	}
	disjointSet := NewDisjointSet(len(positions))

	ans := make([]int, 0)
	islandCount := 0
	for i, p := range positions {
		// set island
		x := p[0]
		y := p[1]
		landNum := i + 1
		grid[x][y] = landNum
		disjointSet.Make(landNum)

		// union island
		islandCount++
		for i := 0; i < 4; i++ {
			tmpX, tmpY := x+dx[i], y+dy[i]
			if tmpX >= 0 && tmpX < m &&
				tmpY >= 0 && tmpY < n &&
				grid[tmpX][tmpY] > 0 {

				tmpLandNum := grid[tmpX][tmpY]
				if disjointSet.Union(tmpLandNum, landNum) {
					islandCount--
				}
			}
		}

		ans = append(ans, islandCount)
	}
	return ans
}

// disjoint sets
type DisjointSet struct {
	parentArr []int
}

func NewDisjointSet(size int) *DisjointSet {
	ret := &DisjointSet{
		parentArr: make([]int, size+1),
	}
	return ret
}

func (s *DisjointSet) Make(x int) {
	s.parentArr[x] = x
}

func (s *DisjointSet) Union(x, y int) bool {
	parentX := s.FindParent(x)
	parentY := s.FindParent(y)
	if parentX != parentY {
		s.parentArr[parentX] = parentY
		return true
	}
	return false
}

func (s *DisjointSet) FindParent(x int) int {
	if s.parentArr[x] == x {
		return x
	}

	s.parentArr[x] = s.FindParent(s.parentArr[x])
	return s.parentArr[x]
}

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
		landNum := i + 1
		x := p[0]
		y := p[1]
		grid[x][y] = landNum

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

// disjoint set
type DisjointSet struct {
	parentArr []int
}

func NewDisjointSet(maxIdx int) *DisjointSet {
	ret := &DisjointSet{
		parentArr: make([]int, maxIdx+1),
	}
	for i := 0; i <= maxIdx; i++ {
		ret.parentArr[i] = i
	}
	return ret
}

func (s *DisjointSet) Union(a, b int) bool {
	parentA := s.FindParent(a)
	parentB := s.FindParent(b)
	if parentA != parentB {
		s.parentArr[parentA] = parentB
		return true
	}
	return false
}

func (s *DisjointSet) FindParent(idx int) int {
	if s.parentArr[idx] == idx {
		return idx
	}

	s.parentArr[idx] = s.FindParent(s.parentArr[idx])
	return s.parentArr[idx]
}

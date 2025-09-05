package leetcode

// multi ans
// index can be used only once
func twoSumMultiAns(nums []int, target int) [][]int {
	ans := make([][]int, 0)

	numBIdxAMap := make(map[int][]int, 0)
	for i, num := range nums {
		idxASlice, ok := numBIdxAMap[num]
		if ok {
			popIdxA := idxASlice[len(idxASlice)-1]
			ans = append(ans, []int{popIdxA, i})

			idxASlice = idxASlice[:len(idxASlice)-1]
			numBIdxAMap[num] = idxASlice
			if len(idxASlice) == 0 {
				delete(numBIdxAMap, num)
			}
			continue
		}

		numB := target - num
		numBIdxASlice, ok := numBIdxAMap[numB]
		if ok {
			numBIdxASlice = append(numBIdxASlice, i)
			numBIdxAMap[numB] = numBIdxASlice
		} else {
			numBIdxAMap[numB] = []int{i}
		}
	}
	return ans
}

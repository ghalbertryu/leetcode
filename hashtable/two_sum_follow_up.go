package hashtable

// find as many as possible disjoint pair
// whose sum is target
func twoSumMultiAns(nums []int, target int) [][]int {
	ans := make([][]int, 0)

	numIdxSliceMap := make(map[int][]int, 0)
	for i, num := range nums {
		numB := target - num
		idxASlice, ok := numIdxSliceMap[numB]
		if ok {
			popIdxA := idxASlice[len(idxASlice)-1]
			ans = append(ans, []int{popIdxA, i})

			// update numIdxSliceMap
			idxASlice = idxASlice[:len(idxASlice)-1]
			numIdxSliceMap[numB] = idxASlice
			if len(idxASlice) == 0 {
				delete(numIdxSliceMap, numB)
			}
			continue
		}

		// update numIdxSliceMap
		idxAs, ok := numIdxSliceMap[num]
		if ok {
			idxAs = append(idxAs, i)
			numIdxSliceMap[num] = idxAs
		} else {
			numIdxSliceMap[num] = []int{i}
		}
	}
	return ans
}

package hashtable

func twoSumFirst(nums []int, target int) []int {
	sameNumAns := []int{}
	numIdxMap := make(map[int]int, 0) // num:idx
	for i, num := range nums {
		if num*2 == target {
			// sameNumAns condition
			sameNumAns = append(sameNumAns, i)
			if len(sameNumAns) == 2 {
				return sameNumAns
			}
		} else {
			numIdxMap[num] = i
		}
	}

	for num, idxA := range numIdxMap {
		if idxB, ok := numIdxMap[target-num]; ok && idxA != idxB {
			return []int{idxA, idxB}
		}
	}
	return []int{}
}

func twoSumSecond(nums []int, target int) []int {
	numBIdxAMap := make(map[int]int, 0)
	for i, num := range nums {
		idxA, ok := numBIdxAMap[num]
		if ok && idxA != i {
			return []int{idxA, i}
		}

		numB := target - num
		numBIdxAMap[numB] = i
	}

	return nil
}

func twoSumThird(nums []int, target int) []int {
	numIdxMap := make(map[int]int, 0)
	for i, num := range nums {
		numB := target - num
		idxA, ok := numIdxMap[numB]
		if ok && idxA != i {
			return []int{idxA, i}
		}

		numIdxMap[num] = i
	}

	return nil
}

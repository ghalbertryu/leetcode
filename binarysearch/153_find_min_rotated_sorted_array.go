package heetcode

func findMin(nums []int) int {
	minimum := -1
	for {
		arrLen := len(nums)
		if arrLen == 0 {
			panic("array is empty")
		}
		if arrLen == 1 {
			minimum = nums[0]
			return minimum
		}

		halfIdx := arrLen / 2
		leftArr := nums[0:halfIdx]
		rightArr := nums[halfIdx:]
		if leftLen := len(leftArr); leftLen > 1 && leftArr[0] > leftArr[leftLen-1] {
			// minimum in leftArr
			nums = leftArr
		} else if rightLen := len(rightArr); rightLen > 1 && rightArr[0] > rightArr[rightLen-1] {
			// minimum in rightArr
			nums = rightArr
		} else if leftArr[0] < rightArr[0] {
			nums = leftArr
		} else {
			nums = rightArr
		}
	}

	return minimum
}

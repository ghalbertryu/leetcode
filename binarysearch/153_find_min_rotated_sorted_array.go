package heetcode

func findMinFirst(nums []int) int {
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

func findMin(nums []int) int {
	arrLen := len(nums)
	if arrLen == 0 {
		panic("array is empty")
	}

	numsIdx0 := nums[0]
	if numsIdx0 < nums[arrLen-1] {
		// sorted
		return numsIdx0
	}

	left := 0
	right := arrLen - 1
	for right-left > 1 {
		mid := (left + right) / 2
		midNum := nums[mid]
		if numsIdx0 < midNum {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[right]
}

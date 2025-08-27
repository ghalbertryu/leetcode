package heetcode

func findMinDuplicatedNum(nums []int) int {
	arrLen := len(nums)
	if arrLen == 0 {
		panic("array is empty")
	}

	if nums[0] < nums[arrLen-1] {
		// sorted
		return nums[0]
	}

	left := 0
	right := arrLen - 1
	for right-left > 1 {
		mid := (left + right) / 2
		if nums[left] == nums[mid] {
			// find left and right
			minLeft := findMinDuplicatedNum(nums[left:mid])
			minRight := findMinDuplicatedNum(nums[mid : right+1])
			if minLeft <= minRight {
				return minLeft
			} else {
				return minRight
			}
		} else if nums[left] < nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[right]
}

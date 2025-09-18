package leetcode

import "math"

func maxSubArraySecond(nums []int) int {
	ans := math.MinInt64

	left := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 ||
			(left < 0 && left < nums[i]) {
			left = nums[i]
		} else {
			left = left + nums[i]
		}
		if left > ans {
			ans = left
		}
	}
	return ans
}

func maxSubArrayFirst(nums []int) int {
	ans := math.MinInt64

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}

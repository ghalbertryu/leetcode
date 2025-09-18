package leetcode

import (
	"log"
	"math"
)

func maxSubArrayPrefixSum(nums []int) int {
	ans := math.MinInt64
	prefixSum := 0
	minPrefixSum := 0

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		maxSum := prefixSum - minPrefixSum
		if maxSum > ans {
			ans = maxSum
		}
		if prefixSum < minPrefixSum {
			minPrefixSum = prefixSum
		}
	}

	return ans
}

func prefixSumExample(nums []int) {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixSum[i] = nums[i]
			continue
		}
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	log.Println(prefixSum)

	// sum of idx range: [2~5]
	log.Println(prefixSum[5] - prefixSum[1])

	// sum of idx range: [i~j]
	//log.Println(prefixSum[j] - prefixSum[i-1])
}

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

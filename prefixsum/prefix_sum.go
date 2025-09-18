package prefixsum

import "log"

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

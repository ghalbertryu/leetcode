package prefixsum

import (
	"log"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	log.Println(maxSubArrayPrefixSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	log.Println(maxSubArraySecond([]int{1}))                                // 1
	log.Println(maxSubArraySecond([]int{1, 2}))                             // 3
	//log.Println(maxSubArray([]int{-2, 1, -3, 400, -1, 2, 1, -5, 400}))
}

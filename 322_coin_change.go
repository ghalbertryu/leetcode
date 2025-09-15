package leetcode

import (
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 { // solution number: coin + dp[number-coin]
				if dp[i] == -1 {
					dp[i] = dp[i-coin] + 1
				} else if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	return dp[amount]
}

func coinChangeTry(coins []int, amount int) int {
	sort.Ints(coins)
	largeCoin := coins[len(coins)-1]
	coinCount := amount / largeCoin
	mod := amount % largeCoin
	if 0 == mod {
		return coinCount
	}
	if len(coins[:len(coins)-1]) == 0 {
		return -1
	}

	for i := 0; i <= coinCount; i++ {
		left := mod + i*largeCoin
		count := coinCount - i
		reduceCount := coinChangeTry(coins[:len(coins)-1], left)
		if reduceCount > 0 {
			return count + reduceCount
		}
	}
	return -1
}

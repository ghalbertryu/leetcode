package dp

func coinChangeHowManyWays(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 0; i <= amount; i++ {
			if i-coin >= 0 {
				dp[i] = dp[i] + dp[i-coin]
				// dp[i]: 原本湊 i 數量可以的方式數量
				// dp[i-coin] 種: 湊 i 數量, 使用 1個coin 時的方式數量
			}
		}
	}
	return dp[amount]
}

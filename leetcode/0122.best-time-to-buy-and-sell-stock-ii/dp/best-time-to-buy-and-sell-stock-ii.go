package problem0122

func maxProfit_1(prices []int) int {
    n := len(prices)
    dp := make([][2]int, n)
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    return dp[n-1][0]
}

func maxProfit_2(prices []int) int {
    n := len(prices)
    dp0, dp1 := 0, -prices[0]
    for i := 1; i < n; i++ {
        dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
    }
    return dp0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
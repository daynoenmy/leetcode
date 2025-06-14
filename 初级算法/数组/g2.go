package main

// 买卖股票
// 动态规划
func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}
func main() {

}

package main

import "math"

//state machine
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	s0 := make([]int, n)
	s1 := make([]int, n)
	s2 := make([]int, n)

	s1[0] = -prices[0]
	s0[0] = 0
	s2[0] = math.MinInt64

	for i := 1; i < n; i++ {
		s0[i] = maxInt(s0[i-1], s2[i-1])
		s1[i] = maxInt(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	return maxInt(s0[n-1], s2[n-1])
}

//dp
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	buy := make([]int, n+1)
	buy[1] = 0 - prices[0]
	sell := make([]int, n+1)

	for i := 2; i <= n; i++ {
		curPrice := prices[i-1]
		// buy[i] = buy[i-1]  means rest for 1 day
		// buy[i] = sell[i-2] - curPrice means the day before yesterday sell (so - 2)
		buy[i] = maxInt(buy[i-1], sell[i-2]-curPrice)
		// sell[i] = sell[i-1] means rest for 1 day
		// sell[i] = buy[i-1]+curPrice means yesterday buy
		sell[i] = maxInt(sell[i-1], buy[i-1]+curPrice)
	}

	return sell[n]
}

//dp
func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	buy := math.MinInt64
	sell := 0
	//before sell
	lastSell := 0

	for _, price := range prices {
		//cannot buy all sell at same day,
		//so when buy use last sell
		//when sell use last buy

		lastBuy := buy
		//rest or buy
		buy = maxInt(lastBuy, lastSell-price)
		lastSell = sell
		//rest or sell
		sell = maxInt(lastSell, lastBuy+price)
	}
	return sell
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func tripleOperation() {

}

func main() {

}

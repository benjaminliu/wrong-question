package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		j := 1
		square := 1
		for square <= i {
			if dp[i-square] < min {
				min = dp[i-square]
			}
			j++
			square = j * j
		}
		dp[i] = min + 1
	}

	return dp[n]
}

func minInt(left int, right int) int {
	if left <= right {
		return left
	}
	return right
}

func main() {

}

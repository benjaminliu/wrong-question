package main

func rob(nums []int) int {
	len1 := len(nums)
	if len1 == 0 {
		return 0
	}
	if len1 == 1 {
		return nums[0]
	}
	if len1 == 2 {
		if nums[0] >= nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	dp := make([]int, len1)
	dp[0] = nums[0]
	if nums[0] >= nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < len1; i++ {
		temp := dp[i-2] + nums[i]
		if temp >= dp[i-1] {
			dp[i] = temp
		} else {
			dp[i] = dp[i-1]
		}
	}

	if dp[len1-1] >= dp[len1-2] {
		return dp[len1-1]
	}
	return dp[len1-2]
}

func main() {

}

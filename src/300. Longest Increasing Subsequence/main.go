package main

import "fmt"

func lengthOfLIS(nums []int) int {
	len1 := len(nums)
	if len1 == 0 {
		return 0
	}
	dp := make([]int, len1)

	dp[0] = 1
	max := 1
	for i := 1; i < len1; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[j]+1, dp[i])
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func lengthOfLIS1(nums []int) int {
	len1 := len(nums)
	if len1 == 0 {
		return 0
	}

	dp := make([]int, len1+1)
	dp[0] = nums[0]
	length := 1
	for i := 1; i < len1; i++ {
		if nums[i] > dp[length-1] {
			dp[length] = nums[i]
			length++
		} else if nums[i] <= dp[0] {
			dp[0] = nums[i]
		} else {
			idx := binarySearch(dp, 0, length-1, nums[i])
			dp[idx] = nums[i]
		}
		fmt.Println(dp)
	}
	return length
}
func binarySearch(dp []int, left int, right int, target int) int {
	for left <= right {
		mid := left + (right-left)>>1
		if dp[mid] < target {
			left = mid + 1
		} else if dp[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	nums := []int{10,9,2,5,3,7,101,18}

	fmt.Println(lengthOfLIS1(nums))
}

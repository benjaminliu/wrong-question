package main

import "fmt"

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res := nums[0]
	curSum := nums[0]
	for i := 1; i < length; i++ {
		temp := curSum + nums[i]
		if temp > nums[i] {
			curSum = temp
		} else {
			curSum = nums[i]
		}
		//fmt.Println(curSum)
		if res < curSum {
			res = curSum
		}
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	//nums[i]          1, -3,  4, -1,  2,  1, -5,  4
	//curSum+num[i]   -1, -2,  2,  3,  5,  6,  1,  5
	//new curSum = max( curSum + num[i] , nums[i] )
	//new curSum   -2  1, -2   4,  3,  5,  6,  1,  5
	fmt.Println(maxSubArray(nums))
}

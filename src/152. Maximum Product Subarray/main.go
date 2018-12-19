package main

func maxProduct(nums []int) int {
	len1 := len(nums)

	if len1 == 1 {
		return nums[0]
	}

	tempMax, tempMin := nums[0], nums[0]

	max := nums[0]
	for i := 1; i < len1; i++ {
		a := tempMax * nums[i]
		b := tempMin * nums[i]

		tempMax = maxInt(nums[i], maxInt(a, b))
		tempMin = minInt(nums[i], minInt(a, b))

		max = maxInt(max, tempMax)
	}

	return max
}

func maxInt(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minInt(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

func main() {
	
}

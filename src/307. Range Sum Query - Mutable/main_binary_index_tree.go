package main

import "fmt"

type NumArray struct {
	bits    []int
	nums    []int
	bitSize int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	obj := NumArray{bitSize: n + 1}

	obj.nums = nums
	obj.bits = make([]int, n+1)

	for i := 0; i < n; i++ {
		obj.change(i+1, nums[i])
	}
	return obj
}

func (this *NumArray) Update(i int, val int) {
	delta := val - this.nums[i]
	this.nums[i] = val
	this.change(i+1, delta)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.getSum(j+1) - this.getSum(i)
}

func (this *NumArray) change(i int, val int) {
	for i < this.bitSize {
		this.bits[i] += val
		i += lowBit(i)
	}
}

func (this *NumArray) getSum(i int) int {
	sum := 0
	for i > 0 {
		sum += this.bits[i]
		i -= lowBit(i)
	}
	return sum
}

// only keep the lowest (rightest) bit 1 and the 0s after (right to) it, then convert it to int
// if x = 4  (100), then return 4 (100)
// if x = 6  (110), then return 2 (10)
// if x = 8  (1000), then return 8 (1000)
func lowBit(x int) int {
	return x & (-x)
}

func main() {
	fmt.Println(lowBit(1)) //1
	fmt.Println(lowBit(2)) //2
	fmt.Println(lowBit(3)) //1
	fmt.Println(lowBit(4)) //4
	fmt.Println(lowBit(5)) //1
	fmt.Println(lowBit(6)) //2
	fmt.Println(lowBit(7)) //1
	fmt.Println(lowBit(8)) //8
	fmt.Println(lowBit(9)) //1
}

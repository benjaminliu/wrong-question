package main

import "fmt"

type SegmentNode struct {
	start int
	end   int
	sum   int
	left  *SegmentNode
	right *SegmentNode
}

func BuildTree(nums []int, lo, hi int) *SegmentNode {
	if lo > hi {
		return nil
	}

	root := &SegmentNode{start: lo, end: hi}
	if lo == hi {
		root.sum = nums[lo]
		return root
	}

	mid := lo + (hi-lo)/2
	root.left = BuildTree(nums, lo, mid)
	root.right = BuildTree(nums, mid+1, hi)
	root.sum = root.left.sum + root.right.sum
	return root
}

func update(this *SegmentNode, i int, val int) {
	//only 1 children
	if this.start == this.end && this.start == i {
		this.sum = val
		return
	}

	//no in here
	if this.start > i || this.end < i {
		return
	}

	mid := this.start + (this.end-this.start)/2
	if i <= mid {
		update(this.left, i, val)
	} else {
		update(this.right, i, val)
	}
	this.sum = this.left.sum + this.right.sum
}

func sum(this *SegmentNode, lo int, hi int) int {
	if this.start == this.end && this.start == hi {
		return this.sum
	}

	mid := this.start + (this.end-this.start)/2
	if hi <= mid {
		return sum(this.left, lo, hi)
	}
	if lo > mid {
		return sum(this.right, lo, hi)
	}

	return sum(this.left, lo, mid) + sum(this.right, mid+1, hi)
}

type NumArray1 struct {
	root *SegmentNode
}

func Constructor1(nums []int) NumArray1 {
	obj := NumArray1{}
	obj.root = BuildTree(nums, 0, len(nums)-1)
	return obj
}

func (this *NumArray1) Update(i int, val int) {
	update(this.root, i, val)
}

func (this *NumArray1) SumRange(i int, j int) int {
	return sum(this.root, i, j)
}

func main() {
	nums := []int{1, 3, 5}
	obj := Constructor1(nums)
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
}

// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func nextPermutation(nums []int) {
	var i int
	var j int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j = len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					break
				}
			}
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

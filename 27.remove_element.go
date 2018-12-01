// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func removeElement(nums []int, val int) int {
	var pre, cur int

	for cur < len(nums) {
		nums[pre] = nums[cur]
		if nums[pre] != val {
			pre++
			cur++
		} else {
			cur++
		}
	}
	return pre
}

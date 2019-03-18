/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 239.sliding_window_maximum.go
#   Created       : 2019-02-13 11:14:44
#   Last Modified : 2019-02-13 11:14:44
#   Describe      :
#
# ====================================================*/
package main

func maxSlidingWindow(nums []int, k int) []int {
	if k < 1 || k > len(nums) {
		return []int{}
	}
	res := make([]int, len(nums)-k+1)
	begin := 0
	end := begin + k - 1
	for end < len(nums) {
		var max = nums[begin]
		for i := begin + 1; i < end+1; i++ {
			if max < nums[i] {
				max = nums[i]
			}
		}
		res[begin] = max
		end++
		begin++
	}
	return res
}

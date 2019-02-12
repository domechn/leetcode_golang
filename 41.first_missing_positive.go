/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 41.first_missing_positive.go
#   Created       : 2019-02-12 16:00:40
#   Last Modified : 2019-02-12 16:00:40
#   Describe      :
#
# ====================================================*/
package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

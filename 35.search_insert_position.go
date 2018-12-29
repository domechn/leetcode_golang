/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 35.search_insert_position.go
#   Created       : 2018-12-29 10:31:11
#   Last Modified : 2018-12-29 10:31:11
#   Describe      :
#
# ====================================================*/
package main

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	if max == -1 {
		return 0
	}
	if nums[min] > target {
		return 0
	}
	if nums[max] < target {
		return max + 1
	}
	for max >= min {
		mid := (max + min) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return max + 1
}

/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 33.search_in_rotated_sorted_array.go
#   Created       : 2018-12-29 10:13:42
#   Last Modified : 2018-12-29 10:13:42
#   Describe      :
#
# ====================================================*/
package main

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for max >= min {
		mid := (max + min) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[max] {
			// 右边有序
			if nums[mid] < target && nums[max] >= target {
				min = mid + 1
			} else {
				max = mid - 1
			}
		} else {
			// 左边有序
			if nums[mid] > target && nums[min] <= target {
				max = mid - 1
			} else {
				min = mid + 1
			}
		}
	}
	return -1
}

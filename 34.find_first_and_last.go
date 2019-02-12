/* ====================================================
#   Copyright (C)2018 All rights reserved.

#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 34.find_first_and_last.go
#   Created       : 2018-12-28 19:40:16
#   Last Modified : 2018-12-28 19:40:16
#   Describe      :
#
# ====================================================*/
package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	if nums[0] > target || nums[len(nums)-1] < target {
		return res
	}
	min := 0
	max := len(nums)
	mid := (max + min) / 2
	for max >= min {

		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			for i := mid; i >= 0 && i < len(nums) && nums[i] == target; i-- {
				res[0] = i
			}
			for i := mid; i >= 0 && i < len(nums) && nums[i] == target; i++ {
				res[1] = i
			}

			break
		}
		mid = (max + min) / 2
	}
	return res
}

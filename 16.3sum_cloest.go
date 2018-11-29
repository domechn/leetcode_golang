// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	closest := nums[0] + nums[1] + nums[2]
	diff := abs(closest - target)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			newDiff := abs(sum - target)
			if diff > newDiff {
				diff = newDiff
				closest = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

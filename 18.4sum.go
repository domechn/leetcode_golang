// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			fix := nums[j]
			beg := j + 1
			end := len(nums) - 1

			for beg < end {
				sum := nums[i] + fix + nums[beg] + nums[end]

				if sum < target {
					beg++
				} else if sum > target {
					end--
				} else {
					res = apend(res, []int{nums[i], fix, nums[beg], nums[end]})
					beg++
					end--
				}
			}
		}
	}
	return res
}

func apend(a [][]int, b []int) [][]int {
	flag := true
	for i := len(a) - 1; i >= 0; i-- {
		v := a[i]
		if v[0] == b[0] && v[1] == b[1] && v[2] == b[2] {
			flag = false
		}
	}
	if flag {
		a = append(a, b)
	}
	return a
}

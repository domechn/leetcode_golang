// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		j := i + 1
		k := len(nums) - 1
		fix := nums[i]

		for j < k {
			r := nums[j] + fix + nums[k]
			if r > 0 {
				k--
			} else if r < 0 {
				j++
			} else {
				res = apend(res, []int{fix, nums[j], nums[k]})
				j++
				k--
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

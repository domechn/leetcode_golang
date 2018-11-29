// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func maxArea(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1
	for i < j {
		res = max(res, area(height[i], height[j], j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func area(a, b, c int) int {
	if a < b {
		return a * c
	} else {
		return b * c
	}
}

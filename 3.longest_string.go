// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func lengthOfLongestSubstring(s string) int {
	var res int
	var temp = make([]int, 256)
	for k := range temp {
		temp[k] = -1
	}
	left := -1
	for i := 0; i < len(s); i++ {
		left = max(left, temp[s[i]])
		temp[s[i]] = i
		res = max(res, i-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

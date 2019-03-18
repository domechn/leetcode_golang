/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 1004.max_consecutive_ones_iii.go
#   Created       : 2019-03-04 16:05:58
#   Last Modified : 2019-03-04 16:05:58
#   Describe      :
#
# ====================================================*/
package main

func longestOnes(A []int, K int) int {
	var left, count, res int
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			count++
		}

		for count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}

		res = max(i-left+1, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

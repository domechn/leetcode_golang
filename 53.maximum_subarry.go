/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 53.maximum_subarry.go
#   Created       : 2019-02-15 19:36:20
#   Last Modified : 2019-02-15 19:36:20
#   Describe      :
#
# ====================================================*/
package leet

import "math"

func maxSubArray(nums []int) int {
	var curSum int
	var res = math.MinInt32

	for _, v := range nums {
		curSum = max(v, v+curSum)
		res = max(res, curSum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

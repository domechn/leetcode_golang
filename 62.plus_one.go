/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 62.plus_one.go
#   Created       : 2019-02-15 19:31:36
#   Last Modified : 2019-02-15 19:31:36
#   Describe      :
#
# ====================================================*/
package main

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}

	return digits
}

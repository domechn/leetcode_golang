/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 69.sqrtx.go
#   Created       : 2019-02-13 10:51:34
#   Last Modified : 2019-02-13 10:51:34
#   Describe      :
#
# ====================================================*/
package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	begin := 0
	end := x
	mid := (end - begin) / 2
	for {
		if mid*mid == x || (mid*mid < x && (mid+1)*(mid+1) > x) {
			return mid
		}

		if mid*mid > x {
			end = mid - 1
		} else {
			begin = mid + 1
		}
		mid = (end-begin)/2 + begin

	}
	return -1
}

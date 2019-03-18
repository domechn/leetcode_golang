/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 70.climbing_stairs.go
#   Created       : 2019-02-12 17:34:11
#   Last Modified : 2019-02-12 17:34:11
#   Describe      :
#
# ====================================================*/
package main

func climbStairs(n int) int {
	f, g := 1, 2
	for i := 1; i < n; i++ {
		g += f
		f = g - f
	}
	return f
}

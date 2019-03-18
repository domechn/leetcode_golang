/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 50.pow.go
#   Created       : 2019-03-11 17:54:08
#   Last Modified : 2019-03-11 17:54:08
#   Describe      :
#
# ====================================================*/
package main

func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	if n == 0 {
		return 1.0
	}
	half := myPow(x, n/2)
	return half * half * myPow(x, n%2)
}

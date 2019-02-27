/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 204.count_primes.go
#   Created       : 2019-02-27 16:07:46
#   Last Modified : 2019-02-27 16:07:46
#   Describe      :
#
# ====================================================*/
package leet

import "math"

func countPrimes(n int) int {
	var res int
	status := make([]bool, n)
	sq := int(math.Sqrt(float64(n)))
	for i := 2; i < sq+1; i++ {
		if !status[i] {
			for j := 2; j*i < n; j++ {
				status[i*j] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !status[i] {
			res++
		}
	}
	return res
}

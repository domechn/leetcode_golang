/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 169.majority_elment.go
#   Created       : 2019-02-12 16:02:54
#   Last Modified : 2019-02-12 16:02:54
#   Describe      :
#
# ====================================================*/
package leet

func majorityElement(nums []int) int {
	var temp = make(map[int]int)
	var res int
	var t int
	for _, v := range nums {
		temp[v]++
		if temp[v] > t {
			t = temp[v]
			res = v
		}
	}
	return res
}

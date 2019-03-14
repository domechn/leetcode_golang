/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 128.longestConsecutive.go
#   Created       : 2019-03-13 15:24:36
#   Last Modified : 2019-03-13 15:24:36
#   Describe      :
#
# ====================================================*/
package leet

func longestConsecutive(nums []int) int {
	var temp = make(map[int]int)
	var ret int
	for _, v := range nums {
		if _, ok := temp[v]; !ok {
			lg := 1 + temp[v-1] + temp[v+1]

			if lg > ret {
				ret = lg
			}
			temp[v-temp[v-1]] = lg
			temp[v+temp[v+1]] = lg
			temp[v] = lg
		}
	}
	return ret
}

/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 118.pascals_triangle.go
#   Created       : 2019-03-01 20:03:14
#   Last Modified : 2019-03-01 20:03:14
#   Describe      :
#
# ====================================================*/
package leet

func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})

	for i := 2; i <= numRows; i++ {
		new1 := append(append([]int{0}, res[i-2]...), 0)
		var r []int
		for j := 1; j < len(new1); j++ {
			r = append(r, new1[j-1]+new1[j])
		}
		res = append(res, r)
	}
	return res
}

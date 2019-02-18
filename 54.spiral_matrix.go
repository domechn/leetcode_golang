/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 54.spiral_matrix.go
#   Created       : 2019-02-15 19:53:20
#   Last Modified : 2019-02-15 19:53:20
#   Describe      :
#
# ====================================================*/
package leet

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	height := len(matrix)
	width := len(matrix[0])

	up, down, left, right := 0, height-1, 0, width-1
	var res = []int{}
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res

}

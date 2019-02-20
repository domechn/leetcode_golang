/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 48.rotate_image.go
#   Created       : 2019-01-02 11:53:07
#   Describe      :
#
# ====================================================*/
package main

import "fmt"



func rotate(matrix [][]int) {
	for i:= 0 i < n ; i ++ {
		for j := i +1 ; j < n ; j ++ {
			a[i][j],a[j][i] = a[j][i] ,a[i][j]
		}
	}
	sort.Reverse()
}

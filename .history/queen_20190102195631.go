/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : queen.go
#   Created       : 2019-01-02 19:50:37
#   Describe      :
#
# ====================================================*/
package main

import "fmt"

func main() {
	res := make([]int, 8)
	queen(&res, 0)
	fmt.Println(res)
}

func queen(res *[]int, row int) {
	if row == 8 {
		fmt.Println(res)
		return
	}
	for i := 0; i < 8; i++ {
		if isOk(row, i) {
			(*res)[row] = i
			queen(res, row+1)
		}
	}
}

func isOk(row, column int) bool {
	for i := row-1; i >= 0 ; i -- {
		
	}
	return true
}

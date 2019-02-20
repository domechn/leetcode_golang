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

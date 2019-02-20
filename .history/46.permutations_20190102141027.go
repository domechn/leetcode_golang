/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 46.permutations.go
#   Created       : 2019-01-02 11:53:38
#   Describe      :
#
# ====================================================*/
package main

func main() {

}

func permute(nums []int) [][]int {
	res := make([][]int{},500)
	for k,num := range nums {
		permuteDFS(res,0,)
	}
}

func permuteDFS(res [][]int,num int,nums[] int) {
	
}

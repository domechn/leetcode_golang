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

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int
	permuteDFS(res, 0, nums, []int{})
	return res
}

func permuteDFS(res [][]int, level int, nums, out []int) {
	if level == len(nums) {
		res = append(res, out)
		fmt.Println(res)
		return
	}
	for i := 0; i < len(nums); i++ {

		out = append(out, nums[i])
		permuteDFS(res, level+1, append(nums[:i], nums[i+1:]...), out)
		out = out[:len(out)-1]

	}

}
